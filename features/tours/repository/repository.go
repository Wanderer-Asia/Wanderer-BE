package repository

import (
	"context"
	fr "wanderer/features/facilities/repository"
	"wanderer/features/tours"
	"wanderer/helpers/filters"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"gorm.io/gorm"
)

func NewTourRepository(mysqlDB *gorm.DB, cld *cloudinary.Cloudinary) tours.Repository {
	return &tourRepository{
		mysqlDB: mysqlDB,
		cld:     cld,
	}
}

type tourRepository struct {
	mysqlDB *gorm.DB
	cld     *cloudinary.Cloudinary
}

func (repo *tourRepository) GetAll(ctx context.Context, flt filters.Filter) ([]tours.Tour, int, error) {
	var mod []Tour
	var totalData int64

	qry := repo.mysqlDB.WithContext(ctx).Model(&Tour{})

	qry = qry.Select(
		"tours.id",
		"tours.title",
		"tours.quota",
		"tours.available",
		"tours.discount",
		"tours.rating",
		"tours.price",
		"tours.thumbnail",
		"tours.start",
	)

	if flt.Search.Keyword != "" {
		qry = qry.Where("title like ?", "%"+flt.Search.Keyword+"%")
	}

	qry.Count(&totalData)

	if flt.Sort.Column != "" {
		dir := "asc"
		if flt.Sort.Direction {
			dir = "desc"
		}

		switch flt.Sort.Column {
		case "rating", "price", "discount":
		case "location":
			qry = qry.Order("Location.name " + dir)
		case "sold":
			qry = qry.Order("(tours.quota-tours.available) " + dir)
		default:
			qry = qry.Order("id desc")
		}
	}

	qry = qry.Joins("Location").Limit(flt.Pagination.Limit).Offset(flt.Pagination.Start)
	if err := qry.Find(&mod).Error; err != nil {
		return nil, 0, err
	}

	var result []tours.Tour
	for _, tour := range mod {
		result = append(result, *tour.ToEntity(nil))
	}

	return result, int(totalData), nil
}

func (repo *tourRepository) GetDetail(ctx context.Context, id uint) (*tours.Tour, error) {
	var modTour = new(Tour)
	if err := repo.mysqlDB.WithContext(ctx).Joins("Airline").Joins("Location").Where(&Tour{Id: id}).First(modTour).Error; err != nil {
		return nil, err
	}

	var modFile []File
	if err := repo.mysqlDB.WithContext(ctx).Joins("JOIN tour_attachment ON tour_attachment.file_id = files.id AND tour_attachment.tour_id = ?", id).Find(&modFile).Error; err != nil {
		return nil, err
	}
	modTour.Picture = modFile

	var modFacilityInclude []fr.Facility
	if err := repo.mysqlDB.WithContext(ctx).Joins("JOIN tour_facility ON tour_facility.facility_id = facilities.id AND tour_facility.tour_id = ?", id).Find(&modFacilityInclude).Error; err != nil {
		return nil, err
	}
	modTour.Facility = modFacilityInclude

	var facilityIncludes []uint
	for _, facility := range modFacilityInclude {
		facilityIncludes = append(facilityIncludes, facility.Id)
	}

	var modFacilityExclude []fr.Facility
	if err := repo.mysqlDB.WithContext(ctx).Where("id not in (?)", facilityIncludes).Find(&modFacilityExclude).Error; err != nil {
		return nil, err
	}

	var modItinerary []Itinerary
	if err := repo.mysqlDB.WithContext(ctx).Where("tour_id = ?", id).Find(&modItinerary).Error; err != nil {
		return nil, err
	}
	modTour.Itinerary = modItinerary

	return modTour.ToEntity(modFacilityExclude), nil
}

func (repo *tourRepository) Create(ctx context.Context, data tours.Tour) error {
	var UniqueFilename = true

	var mod = new(Tour)
	mod.FromEntity(data)
	mod.Available = mod.Quota

	tx := repo.mysqlDB.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Transaction(func(txPict *gorm.DB) error {
		for i := 0; i < len(data.Picture); i++ {
			res, err := repo.cld.Upload.Upload(ctx, data.Picture[i].Raw, uploader.UploadParams{
				UniqueFilename: &UniqueFilename,
				Folder:         "tours",
			})

			if err != nil {
				return err
			}

			mod.Picture[i].Url = res.URL
		}

		if mod.Picture != nil {
			return txPict.Create(mod.Picture).Error
		}

		return nil
	})
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Transaction(func(txTour *gorm.DB) error {
		res, err := repo.cld.Upload.Upload(ctx, data.Thumbnail.Raw, uploader.UploadParams{
			UniqueFilename: &UniqueFilename,
			Folder:         "tours",
		})

		if err != nil {
			return err
		}

		mod.ThumbnailUrl = res.URL

		return txTour.Create(mod).Error
	})
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (repo *tourRepository) Update(ctx context.Context, id uint, data tours.Tour) error {
	var UniqueFilename = true

	var mod = new(Tour)
	mod.FromEntity(data)

	tx := repo.mysqlDB.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Transaction(func(txPict *gorm.DB) error {
		for i := 0; i < len(data.Picture); i++ {
			res, err := repo.cld.Upload.Upload(ctx, data.Picture[i].Raw, uploader.UploadParams{
				UniqueFilename: &UniqueFilename,
				Folder:         "tours",
			})

			if err != nil {
				return err
			}

			mod.Picture[i].Url = res.URL
		}

		if mod.Picture != nil {
			if err := txPict.Model(&Tour{Id: id}).Association("Picture").Unscoped().Clear(); err != nil {
				return err
			}

			return txPict.Create(mod.Picture).Error
		}

		return nil
	})
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Transaction(func(txFacility *gorm.DB) error {
		return txFacility.Model(&Tour{Id: id}).Association("Facility").Unscoped().Clear()
	})
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Transaction(func(txItinerary *gorm.DB) error {
		return txItinerary.Model(&Tour{Id: id}).Association("Itinerary").Unscoped().Clear()
	})
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Transaction(func(txTour *gorm.DB) error {
		res, err := repo.cld.Upload.Upload(ctx, data.Thumbnail.Raw, uploader.UploadParams{
			UniqueFilename: &UniqueFilename,
			Folder:         "tours",
		})

		if err != nil {
			return err
		}

		mod.ThumbnailUrl = res.URL

		return txTour.Where(&Tour{Id: id}).Updates(mod).Error
	})
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

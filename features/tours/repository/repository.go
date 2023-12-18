package repository

import (
	"context"
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
	panic("unimplemented")
}

func (repo *tourRepository) GetDetail(ctx context.Context) (*tours.Tour, error) {
	panic("unimplemented")
}

func (repo *tourRepository) Create(ctx context.Context, data tours.Tour) error {
	var UniqueFilename = true

	var mod = new(Tour)
	mod.FromEntity(data)

	tx := repo.mysqlDB.Begin()
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
	panic("unimplemented")
}

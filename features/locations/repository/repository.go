package repository

import (
	"context"
	"errors"
	"strings"
	"wanderer/features/locations"
	"wanderer/helpers/filters"
	"wanderer/utils/files"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewLocationRepository(mysqlDB *gorm.DB, cloud files.Cloud) locations.Repository {
	return &locationRepository{
		mysqlDB: mysqlDB,
		cloud:   cloud,
	}
}

type locationRepository struct {
	mysqlDB *gorm.DB
	cloud   files.Cloud
}

func (repo *locationRepository) GetAll(ctx context.Context, flt filters.Filter) ([]locations.Location, error) {
	var data []Location

	qry := repo.mysqlDB

	if flt.Search.Keyword != "" {
		qry = qry.Where("name like ?", "%"+flt.Search.Keyword+"%")
	}

	if flt.Pagination.Limit != 0 {
		qry = qry.Limit(flt.Pagination.Limit)
	}

	qry = qry.Find(&data)
	if qry.Error != nil {
		return nil, qry.Error
	}

	var result []locations.Location
	for _, location := range data {
		result = append(result, *location.ToEntity())
	}

	return result, nil
}

func (repo *locationRepository) Create(ctx context.Context, data locations.Location) error {
	var mod = new(Location)
	mod.FromEntity(data)

	if mod.ImageRaw != nil {
		url, err := repo.cloud.Upload(ctx, "locations", mod.ImageRaw)
		if err != nil {
			return err
		}

		mod.ImageUrl = *url
	}

	qry := repo.mysqlDB.Create(mod)
	if qry.Error != nil {
		if strings.Contains(qry.Error.Error(), "1062") {
			return errors.New("used: location name already exist")
		}

		return qry.Error
	}

	if qry.RowsAffected == 0 {
		return errors.New("failed to create location")
	}

	return nil
}

func (repo *locationRepository) Update(ctx context.Context, id uint, data locations.Location) error {
	var mod = new(Location)
	mod.FromEntity(data)

	if mod.ImageRaw != nil {
		url, err := repo.cloud.Upload(ctx, "locations", mod.ImageRaw)
		if err != nil {
			return err
		}

		mod.ImageUrl = *url
	}

	qry := repo.mysqlDB.Where(&Location{Id: id}).Updates(mod)
	if qry.Error != nil {
		if strings.Contains(qry.Error.Error(), "1062") {
			return errors.New("used: location name already exist")
		}

		return qry.Error
	}

	if qry.RowsAffected == 0 {
		return errors.New("not found: location not found")
	}

	return nil
}

func (repo *locationRepository) Delete(ctx context.Context, id uint) error {
	qry := repo.mysqlDB.Where(&Location{Id: id}).Delete(&Location{})
	if qry.Error != nil {
		if strings.Contains(qry.Error.Error(), "1451") {
			return errors.New("used: location used by other resources")
		}

		return qry.Error
	}

	if qry.RowsAffected == 0 {
		return errors.New("not found: location not found")
	}

	return nil
}

func (repo *locationRepository) GetDetail(ctx context.Context, id uint) (*locations.Location, error) {
	var mod = new(Location)
	if err := repo.mysqlDB.WithContext(ctx).Where(&Location{Id: id}).First(&mod).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("not found: location not found")
		}
		return nil, err
	}

	var modTour []Tour
	if err := repo.mysqlDB.WithContext(ctx).Where(&Tour{LocationId: id}).Find(&modTour).Error; err != nil {
		return nil, err
	}
	mod.Tours = modTour

	return mod.ToEntity(), nil
}

func (repo *locationRepository) Import(ctx context.Context, data []locations.Location) error {
	var model []Location
	for _, air := range data {
		var tmpAir = new(Location)
		tmpAir.FromEntity(air)

		model = append(model, *tmpAir)
	}

	qry := repo.mysqlDB.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).CreateInBatches(model, 1000)
	if qry.Error != nil {
		if strings.Contains(qry.Error.Error(), "1062") {
			return nil
		}

		return qry.Error
	}

	return nil
}

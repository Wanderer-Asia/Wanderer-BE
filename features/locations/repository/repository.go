package repository

import (
	"context"
	"wanderer/features/locations"

	"gorm.io/gorm"
)

func NewLocationRepository(mysqlDB *gorm.DB) locations.Repository {
	return &locationRepository{
		mysqlDB: mysqlDB,
	}
}

type locationRepository struct {
	mysqlDB *gorm.DB
}

func (repo *locationRepository) GetAll(ctx context.Context) ([]locations.Location, error) {
	var data []Location
	if err := repo.mysqlDB.Find(&data).Error; err != nil {
		return nil, err
	}

	var result []locations.Location
	for _, location := range data {
		result = append(result, *location.ToEntity())
	}

	return result, nil
}

func (repo *locationRepository) Create(ctx context.Context, data locations.Location) error {
	panic("unimplemented")
}

func (repo *locationRepository) Update(ctx context.Context, id uint, data locations.Location) error {
	panic("unimplemented")
}

func (repo *locationRepository) Delete(ctx context.Context, id uint) error {
	panic("unimplemented")
}

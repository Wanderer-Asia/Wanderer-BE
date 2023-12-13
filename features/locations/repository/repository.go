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
	panic("unimplemented")
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

package repository

import (
	"context"
	"wanderer/features/tours"
	"wanderer/helpers/filters"

	"github.com/cloudinary/cloudinary-go/v2"
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
	panic("unimplemented")
}

func (repo *tourRepository) Update(ctx context.Context, id uint, data tours.Tour) error {
	panic("unimplemented")
}

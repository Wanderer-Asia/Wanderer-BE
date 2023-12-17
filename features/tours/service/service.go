package service

import (
	"context"
	"wanderer/features/tours"
	"wanderer/helpers/filters"
)

func NewTourService(repo tours.Repository) tours.Service {
	return &tourService{
		repo: repo,
	}
}

type tourService struct {
	repo tours.Repository
}

func (srv *tourService) GetAll(ctx context.Context, flt filters.Filter) ([]tours.Tour, int, error) {
	panic("unimplemented")
}

func (srv *tourService) GetDetail(ctx context.Context) (*tours.Tour, error) {
	panic("unimplemented")
}

func (srv *tourService) Create(ctx context.Context, data tours.Tour) error {
	panic("unimplemented")
}

func (srv *tourService) Update(ctx context.Context, id uint, data tours.Tour) error {
	panic("unimplemented")
}

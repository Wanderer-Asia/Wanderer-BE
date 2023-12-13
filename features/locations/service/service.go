package service

import (
	"context"
	"wanderer/features/locations"
)

func NewLocationService(repo locations.Repository) locations.Service {
	return &locationService{
		repo: repo,
	}
}

type locationService struct {
	repo locations.Repository
}

func (srv *locationService) GetAll(ctx context.Context) ([]locations.Location, error) {
	panic("unimplemented")
}

func (srv *locationService) Create(ctx context.Context, data locations.Location) error {
	panic("unimplemented")
}

func (srv *locationService) Update(ctx context.Context, id uint, data locations.Location) error {
	panic("unimplemented")
}

func (srv *locationService) Delete(ctx context.Context, id uint) error {
	panic("unimplemented")
}

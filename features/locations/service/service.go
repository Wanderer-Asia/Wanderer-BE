package service

import (
	"context"
	"errors"
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
	result, err := srv.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (srv *locationService) Create(ctx context.Context, data locations.Location) error {
	if data.Name == "" {
		return errors.New("validate: please input name")
	}

	if err := srv.repo.Create(ctx, data); err != nil {
		return err
	}

	return nil
}

func (srv *locationService) Update(ctx context.Context, id uint, data locations.Location) error {
	panic("unimplemented")
}

func (srv *locationService) Delete(ctx context.Context, id uint) error {
	panic("unimplemented")
}

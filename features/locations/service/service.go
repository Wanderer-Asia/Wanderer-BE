package service

import (
	"context"
	"errors"
	"wanderer/features/locations"
	"wanderer/helpers/filters"
)

func NewLocationService(repo locations.Repository) locations.Service {
	return &locationService{
		repo: repo,
	}
}

type locationService struct {
	repo locations.Repository
}

func (srv *locationService) GetAll(ctx context.Context, flt filters.Filter) ([]locations.Location, error) {
	result, err := srv.repo.GetAll(ctx, flt)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (srv *locationService) Create(ctx context.Context, data locations.Location) error {
	if data.Name == "" {
		return errors.New("validate: name can't be empty")
	}

	if data.ImageRaw == nil {
		return errors.New("validate: image can't be empty")
	}

	if err := srv.repo.Create(ctx, data); err != nil {
		return err
	}

	return nil
}

func (srv *locationService) Update(ctx context.Context, id uint, data locations.Location) error {
	if id == 0 {
		return errors.New("validate: invalid location id")
	}

	if data.Name == "" {
		return errors.New("validate: name can't be empty")
	}

	if err := srv.repo.Update(ctx, id, data); err != nil {
		return err
	}

	return nil
}

func (srv *locationService) Delete(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("validate: invalid location id")
	}

	if err := srv.repo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (srv *locationService) GetDetail(ctx context.Context, id uint) (*locations.Location, error) {
	if id == 0 {
		return nil, errors.New("validate: invalid location id")
	}

	result, err := srv.repo.GetDetail(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (srv *locationService) Import(ctx context.Context, data []locations.Location) error {
	if err := srv.repo.Import(ctx, data); err != nil {
		return err
	}

	return nil
}

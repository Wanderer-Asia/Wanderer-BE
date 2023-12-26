package service

import (
	"context"
	"errors"
	"wanderer/features/facilities"
	"wanderer/helpers/filters"
)

func NewFacilityService(repo facilities.Repository) facilities.Service {
	return &facilityService{
		repo: repo,
	}
}

type facilityService struct {
	repo facilities.Repository
}

func (srv *facilityService) Create(newfacility facilities.Facility) error {
	if newfacility.Name == "" {
		return errors.New("validate: name can't be empty")
	}

	if err := srv.repo.Create(newfacility); err != nil {
		return err
	}

	return nil
}

func (srv *facilityService) GetAll(flt filters.Filter) ([]facilities.Facility, error) {
	result, err := srv.repo.GetAll(flt)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (srv *facilityService) Update(id uint, updateFacility facilities.Facility) error {
	if id == 0 {
		return errors.New("validate: ivalid facility id")
	}

	if updateFacility.Name == "" {
		return errors.New("validate: name can't be empty")
	}

	if err := srv.repo.Update(id, updateFacility); err != nil {
		return err
	}

	return nil
}

func (srv *facilityService) Delete(id uint) error {
	if id == 0 {
		return errors.New("validate: ivalid facility id")
	}

	if err := srv.repo.Delete(id); err != nil {
		return err
	}

	return nil
}

func (srv *facilityService) Import(ctx context.Context, data []facilities.Facility) error {
	if err := srv.repo.Import(ctx, data); err != nil {
		return err
	}

	return nil
}

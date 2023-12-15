package service

import (
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
	panic("unimplemented")
}

func (srv *facilityService) Update(id uint, updateFacility facilities.Facility) error {
	panic("unimplemented")
}

func (srv *facilityService) Delete(id uint) error {
	panic("unimplemented")
}

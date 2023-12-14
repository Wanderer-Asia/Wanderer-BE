package service

import (
	"errors"
	"wanderer/features/airlines"
)

func NewAirlineService(repo airlines.Repository) airlines.Service {
	return &airlineService{
		repo: repo,
	}
}

type airlineService struct {
	repo airlines.Repository
}

func (srv *airlineService) Create(newAirline airlines.Airline) error {
	if newAirline.Name == "" {
		return errors.New("validate: name can't be empty")
	}

	if err := srv.repo.Create(newAirline); err != nil {
		return err
	}

	return nil
}

func (srv *airlineService) GetAll() ([]airlines.Airline, error) {
	panic("unimplemented")
}

func (srv *airlineService) Update(id uint, updateAirline airlines.Airline) error {
	panic("unimplemented")
}

func (srv *airlineService) Delete(id uint) error {
	panic("unimplemented")
}

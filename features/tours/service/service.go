package service

import (
	"context"
	"errors"
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
	result, totalData, err := srv.repo.GetAll(ctx, flt)
	if err != nil {
		return nil, 0, err
	}

	return result, totalData, nil
}

func (srv *tourService) GetDetail(ctx context.Context, id uint) (*tours.Tour, error) {
	if id == 0 {
		return nil, errors.New("validate: invalid tour id")
	}

	result, err := srv.repo.GetDetail(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (srv *tourService) GetByLocation(ctx context.Context, id uint) ([]tours.Tour, error) {
	result, err := srv.repo.GetByLocation(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (srv *tourService) Create(ctx context.Context, data tours.Tour) error {
	if data.Title == "" {
		return errors.New("validate: please fill title correctly")
	}

	if data.Description == "" {
		return errors.New("validate: please fill description correctly")
	}

	if data.Price == 0 {
		return errors.New("validate: please fill price correctly")
	}

	if data.Start.IsZero() {
		return errors.New("validate: please fill start date correctly")
	}

	if data.Finish.IsZero() {
		return errors.New("validate: please fill finish date correctly")
	}

	if data.Quota == 0 {
		return errors.New("validate: please fill quota correctly")
	}

	if data.Thumbnail.Raw == nil {
		return errors.New("validate: please fill thumbnail correctly")
	}

	if len(data.Itinerary) == 0 {
		return errors.New("validate: please fill itinerary correctly")
	}

	if data.Location.Id == 0 {
		return errors.New("validate: please fill location correctly")
	}

	if data.Airline.Id == 0 {
		return errors.New("validate: please fill airline correctly")
	}

	if err := srv.repo.Create(ctx, data); err != nil {
		return err
	}

	return nil
}

func (srv *tourService) Update(ctx context.Context, id uint, data tours.Tour) error {
	if id == 0 {
		return errors.New("validate: invalid tour id")
	}

	if data.Title == "" {
		return errors.New("validate: please fill title correctly")
	}

	if data.Description == "" {
		return errors.New("validate: please fill description correctly")
	}

	if data.Price == 0 {
		return errors.New("validate: please fill price correctly")
	}

	if data.Start.IsZero() {
		return errors.New("validate: please fill start date correctly")
	}

	if data.Finish.IsZero() {
		return errors.New("validate: please fill finish date correctly")
	}

	if data.Quota == 0 {
		return errors.New("validate: please fill quota correctly")
	}

	if data.Thumbnail.Raw == nil {
		return errors.New("validate: please fill thumbnail correctly")
	}

	if len(data.Itinerary) == 0 {
		return errors.New("validate: please fill itinerary correctly")
	}

	if data.Location.Id == 0 {
		return errors.New("validate: please fill location correctly")
	}

	if data.Airline.Id == 0 {
		return errors.New("validate: please fill airline correctly")
	}

	if err := srv.repo.Update(ctx, id, data); err != nil {
		return err
	}

	return nil
}

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

func (srv *tourService) Create(ctx context.Context, data tours.Tour) error {
	if data.Title == "" {
		return errors.New("validate: title can't be empty")
	}

	if data.Description == "" {
		return errors.New("validate: description can't be empty")
	}

	if data.Price == 0 {
		return errors.New("validate: price can't be empty")
	}

	if data.Start.IsZero() {
		return errors.New("validate: start date can't be empty")
	}

	if data.Finish.IsZero() {
		return errors.New("validate: finish date can't be empty")
	}

	if data.Quota == 0 {
		return errors.New("validate: quota can't be empty")
	}

	if data.Thumbnail.Raw == nil {
		return errors.New("validate: thumbnail can't be empty")
	}

	if len(data.Itinerary) == 0 {
		return errors.New("validate: itinerary can't be empty")
	}

	if data.Location.Id == 0 {
		return errors.New("validate: location can't be empty")
	}

	if data.Airline.Id == 0 {
		return errors.New("validate: airline can't be empty")
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
		return errors.New("validate: title can't be empty")
	}

	if data.Description == "" {
		return errors.New("validate: description can't be empty")
	}

	if data.Price == 0 {
		return errors.New("validate: price can't be empty")
	}

	if data.Start.IsZero() {
		return errors.New("validate: start date can't be empty")
	}

	if data.Finish.IsZero() {
		return errors.New("validate: finish date can't be empty")
	}

	if data.Quota == 0 {
		return errors.New("validate: quota can't be empty")
	}

	if len(data.Itinerary) == 0 {
		return errors.New("validate: itinerary can't be empty")
	}

	if data.Location.Id == 0 {
		return errors.New("validate: location can't be empty")
	}

	if data.Airline.Id == 0 {
		return errors.New("validate: airline can't be empty")
	}

	if err := srv.repo.Update(ctx, id, data); err != nil {
		return err
	}

	return nil
}

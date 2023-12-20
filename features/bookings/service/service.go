package service

import (
	"context"
	"errors"
	"wanderer/features/bookings"
	"wanderer/helpers/filters"
)

func NewBookingService(repo bookings.Repository) bookings.Service {
	return &bookingService{
		repo: repo,
	}
}

type bookingService struct {
	repo bookings.Repository
}

func (srv *bookingService) GetAll(ctx context.Context, flt filters.Filter) ([]bookings.Booking, int, error) {
	panic("unimplemented")
}

func (srv *bookingService) GetDetail(ctx context.Context, code int) (*bookings.Booking, error) {
	if code == 0 {
		return nil, errors.New("validate: please fill booking code correctly")
	}

	result, err := srv.repo.GetDetail(ctx, code)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (srv *bookingService) Create(ctx context.Context, data bookings.Booking) (*bookings.Booking, error) {
	if data.Tour.Id == 0 {
		return nil, errors.New("validate: please fill tour id correctly")
	}

	if data.User.Id == 0 {
		return nil, errors.New("validate: please fill user id correctly")
	}

	if len(data.Detail) == 0 {
		return nil, errors.New("validate: please fill passenger data correctly")
	}

	for _, detail := range data.Detail {
		if detail.DocumentNumber == "" {
			return nil, errors.New("validate: please fill document number correctly")
		}

		if detail.Greeting == "" {
			return nil, errors.New("validate: please fill greeting correctly")
		}

		if detail.Name == "" {
			return nil, errors.New("validate: please fill name correctly")
		}

		if detail.Nationality == "" {
			return nil, errors.New("validate: please fill nationality correctly")
		}

		if detail.DOB.IsZero() {
			return nil, errors.New("validate: please fill date of birth correctly")
		}
	}

	result, err := srv.repo.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (srv *bookingService) Update(ctx context.Context, code int, data bookings.Booking) (*bookings.Booking, error) {
	if code == 0 {
		return nil, errors.New("validate: please fill booking code correctly")
	}

	result, err := srv.repo.Update(ctx, code, data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

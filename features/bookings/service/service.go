package service

import (
	"context"
	"errors"
	"wanderer/features/bookings"
	"wanderer/helpers/filters"

	"github.com/labstack/echo/v4"
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
	result, totalData, err := srv.repo.GetAll(ctx, flt)
	if err != nil {
		return nil, 0, err
	}

	return result, totalData, nil
}

func (srv *bookingService) GetDetail(ctx context.Context, code int) (*bookings.Booking, error) {
	if code == 0 {
		return nil, errors.New("validate: invalid booking code")
	}

	result, err := srv.repo.GetDetail(ctx, code)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (srv *bookingService) Create(ctx context.Context, data bookings.Booking) (*bookings.Booking, error) {
	if data.Tour.Id == 0 {
		return nil, errors.New("validate: tour id can't be empty")
	}

	if data.User.Id == 0 {
		return nil, errors.New("validate: user id can't be empty")
	}

	if data.Payment.Bank == "" {
		return nil, errors.New("validate: payment method can't be empty")
	}

	if len(data.Detail) == 0 {
		return nil, errors.New("validate: passenger data can't be empty")
	}

	for _, detail := range data.Detail {
		if detail.DocumentNumber == "" {
			return nil, errors.New("validate: document number can't be empty")
		}

		if detail.Greeting == "" {
			return nil, errors.New("validate: greeting can't be empty")
		}

		if detail.Name == "" {
			return nil, errors.New("validate: name can't be empty")
		}

		if detail.Nationality == "" {
			return nil, errors.New("validate: nationality can't be empty")
		}

		if detail.DOB.IsZero() {
			return nil, errors.New("validate: date of birth can't be empty")
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
		return nil, errors.New("validate: invalid booking code")
	}

	result, err := srv.repo.Update(ctx, code, data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (srv *bookingService) Export(c echo.Context, typeFile string) error {
	result, err := srv.repo.Export()
	if err != nil {
		return err
	}

	var data []bookings.Booking
	for _, export := range result {
		data = append(data, export)
	}

	switch typeFile {
	case "pdf":
		err = srv.repo.ExportFilePDF(c, data)
	case "csv":
		err = srv.repo.ExportFileCsv(c, data)
	case "xlsx":
		err = srv.repo.ExportFileExcel(c, data)
	default:
		err = errors.New("unsupported file type")
	}

	return err
}

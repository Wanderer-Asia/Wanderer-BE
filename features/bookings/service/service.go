package service

import (
	"context"
	"errors"
	"time"
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

	tour, err := srv.repo.GetTourById(ctx, data.Tour.Id)
	if err != nil {
		return nil, err
	}

	if tour.Start.Before(time.Now()) {
		return nil, errors.New("unprocessable: tour has been started")
	}

	result, err := srv.repo.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (srv *bookingService) UpdateBookingStatus(ctx context.Context, code int, status string) error {
	if code == 0 {
		return errors.New("validate: invalid booking code")
	}

	oldData, err := srv.repo.GetDetail(ctx, code)
	if err != nil {
		return err
	}

	switch status {
	case "cancel":
		if oldData.Status != "pending" {
			return errors.New("unprocessable: booking can't canceled")
		}
	case "refund":
		if oldData.Status != "approved" {
			return errors.New("unprocessable: refund request denied")
		}
	case "refunded":
		if oldData.Status != "refund" {
			return errors.New("unprocessable: can't approve refund without refund request")
		}
	default:
		return errors.New("validate: invalid booking status")
	}

	if err := srv.repo.UpdateBookingStatus(ctx, code, status); err != nil {
		return err
	}

	return nil
}

func (srv *bookingService) UpdatePaymentStatus(ctx context.Context, code int, paymentStatus string) error {
	if code == 0 {
		return errors.New("validate: invalid booking code")
	}

	var bookingStatus = "pending"
	switch paymentStatus {
	case "settlement":
		bookingStatus = "approved"
	case "cancel", "expire":
		bookingStatus = "cancel"
	case "capture", "deny", "pending":
		bookingStatus = "pending"
	default:
		return errors.New("validate: invalid payment status")
	}

	if err := srv.repo.UpdatePaymentStatus(ctx, code, bookingStatus, paymentStatus); err != nil {
		return err
	}

	return nil
}

func (srv *bookingService) ChangePaymentMethod(ctx context.Context, code int, data bookings.Payment) (*bookings.Payment, error) {
	if code == 0 {
		return nil, errors.New("validate: invalid booking code")
	}

	if data.Bank == "" {
		return nil, errors.New("validate: payment method can't be empty")
	}

	oldData, err := srv.repo.GetDetail(ctx, code)
	if err != nil {
		return nil, err
	}

	if oldData.Status != "pending" || oldData.Payment.Status != "pending" {
		return nil, errors.New("unprocessable: can't change payment method")
	}

	if data.Bank == oldData.Payment.Bank && oldData.Payment.ExpiredAt.After(time.Now()) {
		return &oldData.Payment, nil
	}

	oldData.Payment = data
	result, err := srv.repo.ChangePaymentMethod(ctx, code, *oldData)
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

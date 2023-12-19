package repository

import (
	"context"
	"wanderer/features/bookings"
	tr "wanderer/features/tours/repository"
	ur "wanderer/features/users/repository"
	"wanderer/utils/payments"

	"gorm.io/gorm"
)

func NewBookingRepository(mysqlDB *gorm.DB, payment payments.Midtrans) bookings.Repository {
	return &bookingRepository{
		mysqlDB: mysqlDB,
		payment: payment,
	}
}

type bookingRepository struct {
	mysqlDB *gorm.DB
	payment payments.Midtrans
}

func (repo *bookingRepository) GetAll(ctx context.Context) ([]bookings.Booking, int, error) {
	panic("unimplemented")
}

func (repo *bookingRepository) GetDetail(ctx context.Context, code int) (*bookings.Booking, error) {
	panic("unimplemented")
}

func (repo *bookingRepository) Create(ctx context.Context, data bookings.Booking) (*bookings.Booking, error) {
	var modBooking = new(Booking)
	modBooking.FromEntity(data)

	var modTour = new(tr.Tour)
	if err := repo.mysqlDB.Where(&tr.Tour{Id: modBooking.TourId}).First(modTour).Error; err != nil {
		return nil, err
	}
	modBooking.Tour = *modTour

	var modUser = new(ur.User)
	if err := repo.mysqlDB.Where(&ur.User{Id: modBooking.UserId}).First(modUser).Error; err != nil {
		return nil, err
	}
	modBooking.User = *modUser

	modBooking.CalcTotal(*modTour)

	if err := modBooking.GenerateCode(); err != nil {
		return nil, err
	}

	res, err := repo.payment.NewBookingPayment(*modBooking.ToEntity())
	if err != nil {
		return nil, err
	}
	var modPayment = new(Payment)
	modPayment.FromEntity(*res)
	modBooking.Payment = *modPayment

	if err := repo.mysqlDB.Create(modBooking).Error; err != nil {
		return nil, err
	}

	return modBooking.ToEntity(), nil
}

func (repo *bookingRepository) Update(ctx context.Context, code int, data bookings.Booking) error {
	panic("unimplemented")
}

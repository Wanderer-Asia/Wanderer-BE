package repository

import (
	"context"
	"time"
	"wanderer/features/bookings"
	"wanderer/features/tours"
	tr "wanderer/features/tours/repository"
	ur "wanderer/features/users/repository"
	"wanderer/utils/payments"

	"gorm.io/gorm"
)

func NewBookingRepository(mysqlDB *gorm.DB, tourRepository tours.Repository, payment payments.Midtrans) bookings.Repository {
	return &bookingRepository{
		mysqlDB:  mysqlDB,
		payment:  payment,
		tourRepo: tourRepository,
	}
}

type bookingRepository struct {
	mysqlDB  *gorm.DB
	payment  payments.Midtrans
	tourRepo tours.Repository
}

func (repo *bookingRepository) GetAll(ctx context.Context) ([]bookings.Booking, int, error) {
	panic("unimplemented")
}

func (repo *bookingRepository) GetDetail(ctx context.Context, code int) (*bookings.Booking, error) {
	var mod = new(Booking)
	if err := repo.mysqlDB.WithContext(ctx).Where(&Booking{Code: code}).Joins("User").First(mod).Error; err != nil {
		return nil, err
	}
	data := mod.ToEntity()

	var modBookinDetail []BookingDetail
	if err := repo.mysqlDB.WithContext(ctx).Where(&BookingDetail{BookingCode: code}).Find(&modBookinDetail).Error; err != nil {
		return nil, err
	}
	mod.Detail = modBookinDetail

	modTour, err := repo.tourRepo.GetDetail(ctx, mod.TourId)
	if err != nil {
		return nil, err
	}
	data.Tour = *modTour

	return data, nil
}

func (repo *bookingRepository) Create(ctx context.Context, data bookings.Booking) (*bookings.Booking, error) {
	var modBooking = new(Booking)
	modBooking.FromEntity(data)

	var modTour = new(tr.Tour)
	if err := repo.mysqlDB.WithContext(ctx).Where(&tr.Tour{Id: modBooking.TourId}).First(modTour).Error; err != nil {
		return nil, err
	}
	modBooking.Tour = *modTour

	var modUser = new(ur.User)
	if err := repo.mysqlDB.WithContext(ctx).Where(&ur.User{Id: modBooking.UserId}).First(modUser).Error; err != nil {
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

	if err := repo.mysqlDB.WithContext(ctx).Create(modBooking).Error; err != nil {
		return nil, err
	}

	return modBooking.ToEntity(), nil
}

func (repo *bookingRepository) Update(ctx context.Context, code int, data bookings.Booking) (*bookings.Booking, error) {
	var modNewBooking = new(Booking)
	modNewBooking.FromEntity(data)

	var modOldBooking = new(Booking)
	if err := repo.mysqlDB.WithContext(ctx).Where(&Booking{Code: code}).Joins("User").First(modOldBooking).Error; err != nil {
		return nil, err
	}

	var modOldBookingDetail []BookingDetail
	if err := repo.mysqlDB.WithContext(ctx).Where(&BookingDetail{BookingCode: code}).Find(&modOldBookingDetail).Error; err != nil {
		return nil, err
	}
	modOldBooking.Detail = modOldBookingDetail

	tx := repo.mysqlDB.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if modNewBooking.Payment.Bank != "" && modNewBooking.Payment.Bank != modOldBooking.Payment.Bank {
		err := tx.WithContext(ctx).Transaction(func(txTour *gorm.DB) error {
			if err := repo.payment.CancelBookingPayment(code); err != nil {
				return err
			}

			var retries = 2
			var complete = false

			for retries <= 2 || complete {
				modOldBooking.Payment = modNewBooking.Payment
				res, err := repo.payment.NewBookingPayment(*modOldBooking.ToEntity())
				if err != nil && retries < 2 {
					time.Sleep(50)
					retries++

					continue
				} else if err != nil {
					return err
				} else {
					var modPayment = new(Payment)
					modPayment.FromEntity(*res)
					data.Payment = *res
					data.Total = modOldBooking.Total
					modNewBooking.Payment = *modPayment
					return nil
				}
			}

			return nil
		})
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	err := tx.WithContext(ctx).Transaction(func(txTour *gorm.DB) error {
		return txTour.WithContext(ctx).Omit("Detail").Where(&Booking{Code: code}).Updates(modNewBooking).Error
	})
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.WithContext(ctx).Commit().Error; err != nil {
		return nil, err
	}

	return &data, nil
}

package repository

import (
	"context"
	"errors"
	"time"
	"wanderer/features/bookings"
	"wanderer/helpers/filters"
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

func (repo *bookingRepository) GetAll(ctx context.Context, flt filters.Filter) ([]bookings.Booking, int, error) {
	var mod []Booking
	var totalData int64
	var data []bookings.Booking

	qry := repo.mysqlDB.WithContext(ctx).Model(&Booking{})

	qry.Count(&totalData)

	if flt.Pagination.Limit != 0 {
		qry = qry.Limit(flt.Pagination.Limit)
	}

	if flt.Pagination.Start != 0 {
		qry = qry.Offset(flt.Pagination.Start)
	}

	if err := qry.Joins("User").Joins("Tour", repo.mysqlDB.Select("title", "start", "finish").Model(&Tour{})).Find(&mod).Error; err != nil {
		return nil, int(totalData), err
	}

	for _, booking := range mod {
		booking.Payment = Payment{}
		data = append(data, *booking.ToEntity())
	}

	return data, int(totalData), nil
}

func (repo *bookingRepository) GetDetail(ctx context.Context, code int) (*bookings.Booking, error) {
	var mod = new(Booking)
	if err := repo.mysqlDB.WithContext(ctx).Where(&Booking{Code: code}).First(mod).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("not found: booking not found")
		}
		return nil, err
	}

	var modBookinDetail []BookingDetail
	if err := repo.mysqlDB.WithContext(ctx).Where(&BookingDetail{BookingCode: code}).Find(&modBookinDetail).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("not found: booking not found")
		}
		return nil, err
	}
	mod.Detail = modBookinDetail
	data := mod.ToEntity()

	var modTour = new(Tour)
	if err := repo.mysqlDB.WithContext(ctx).Joins("Airline").Where(&Tour{Id: mod.TourId}).First(modTour).Error; err != nil {
		return nil, err
	}

	var modFile []File
	if err := repo.mysqlDB.WithContext(ctx).Joins("JOIN tour_attachment ON tour_attachment.file_id = files.id AND tour_attachment.tour_id = ?", mod.TourId).Find(&modFile).Error; err != nil {
		return nil, err
	}
	modTour.Picture = modFile

	var modFacilityInclude []Facility
	if err := repo.mysqlDB.WithContext(ctx).Joins("JOIN tour_facility ON tour_facility.facility_id = facilities.id AND tour_facility.tour_id = ?", mod.TourId).Find(&modFacilityInclude).Error; err != nil {
		return nil, err
	}
	modTour.Facility = modFacilityInclude

	var facilityIncludes []uint
	for _, facility := range modFacilityInclude {
		facilityIncludes = append(facilityIncludes, facility.Id)
	}

	var modFacilityExclude []Facility
	if err := repo.mysqlDB.WithContext(ctx).Where("id not in (?)", facilityIncludes).Find(&modFacilityExclude).Error; err != nil {
		return nil, err
	}

	var modItinerary []Itinerary
	if err := repo.mysqlDB.WithContext(ctx).Where("tour_id = ?", mod.TourId).Find(&modItinerary).Error; err != nil {
		return nil, err
	}
	modTour.Itinerary = modItinerary

	var modReviews []Review
	if err := repo.mysqlDB.WithContext(ctx).Where("tour_id = ?", mod.TourId).Joins("User").Find(&modReviews).Error; err != nil {
		return nil, err
	}
	modTour.Reviews = modReviews

	data.Tour = *modTour.ToEntity(modFacilityExclude)

	return data, nil
}

func (repo *bookingRepository) GetTourById(ctx context.Context, tourId uint) (*bookings.Tour, error) {
	var mod = new(Tour)

	if err := repo.mysqlDB.WithContext(ctx).Where(&Tour{Id: tourId}).First(mod).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("not found: tour not found")
		}
		return nil, err
	}

	return mod.ToEntity(nil), nil
}

func (repo *bookingRepository) Create(ctx context.Context, data bookings.Booking) (*bookings.Booking, error) {
	var modBooking = new(Booking)
	modBooking.FromEntity(data)

	var modTour = new(Tour)
	if err := repo.mysqlDB.WithContext(ctx).Where(&Tour{Id: modBooking.TourId}).First(modTour).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("not found: tour not found")
		}
		return nil, err
	}
	modBooking.Tour = *modTour

	var modUser = new(User)
	if err := repo.mysqlDB.WithContext(ctx).Where(&User{Id: modBooking.UserId}).First(modUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("not found: user not found")
		}
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

	if err := repo.mysqlDB.WithContext(ctx).Omit("User", "Tour").Create(modBooking).Error; err != nil {
		return nil, err
	}

	modBooking.User = User{}
	modBooking.Tour = Tour{}
	modBooking.Status = ""
	modBooking.Detail = nil

	return modBooking.ToEntity(), nil
}

func (repo *bookingRepository) UpdateBookingStatus(ctx context.Context, code int, status string) error {
	tx := repo.mysqlDB.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if status == "refunded" {
		err := tx.WithContext(ctx).Transaction(func(txTour *gorm.DB) error {
			return txTour.WithContext(ctx).
				Model(&Tour{}).
				Where("id = (SELECT tour_id FROM bookings where code = ? AND status = 'refund')", code).
				Update("available", gorm.Expr("available + (SELECT COUNT(id) FROM booking_details where booking_code = ?)", code)).Error
		})
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if status == "cancel" {
		repo.payment.CancelBookingPayment(code)
	}

	err := tx.WithContext(ctx).Transaction(func(txTour *gorm.DB) error {
		return txTour.WithContext(ctx).Where("code = ?", code).Updates(&Booking{Status: status}).Error
	})
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.WithContext(ctx).Commit().Error; err != nil {
		return err
	}

	return nil
}

func (repo *bookingRepository) UpdatePaymentStatus(ctx context.Context, code int, bookingStatus string, paymentStatus string) error {
	tx := repo.mysqlDB.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if bookingStatus == "approved" {
		err := tx.WithContext(ctx).Transaction(func(txTour *gorm.DB) error {
			return txTour.WithContext(ctx).
				Model(&Tour{}).
				Where("id = (SELECT tour_id FROM bookings where code = ? AND status = 'pending')", code).
				Update("available", gorm.Expr("available - (SELECT COUNT(id) FROM booking_details where booking_code = ?)", code)).Error
		})
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err := tx.WithContext(ctx).Transaction(func(txTour *gorm.DB) error {
		return txTour.WithContext(ctx).Where("code = ?", code).Updates(&Booking{Status: bookingStatus, Payment: Payment{Status: paymentStatus}}).Error
	})
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.WithContext(ctx).Commit().Error; err != nil {
		return err
	}

	return nil
}

func (repo *bookingRepository) ChangePaymentMethod(ctx context.Context, code int, data bookings.Booking) (*bookings.Payment, error) {
	var newPayment = new(bookings.Payment)
	var retries = 2
	var complete = false

	repo.payment.CancelBookingPayment(code)
	for retries <= 2 || complete {
		res, err := repo.payment.NewBookingPayment(data)
		if err == nil {
			newPayment = res

			break
		} else if retries < 2 {
			time.Sleep(50)
			retries++

			continue
		}

		return nil, err
	}

	var modPayment = new(Payment)
	modPayment.FromEntity(*newPayment)
	if err := repo.mysqlDB.WithContext(ctx).Where(Booking{Code: code}).Updates(&Booking{Payment: *modPayment}).Error; err != nil {
		return nil, err
	}

	return newPayment, nil
}

func (repo *bookingRepository) Export(ctx context.Context) ([]bookings.Booking, error) {
	var mod []Booking
	var data []bookings.Booking

	qry := repo.mysqlDB.WithContext(ctx).Model(&Booking{})

	if err := qry.Joins("User").Joins("Tour", repo.mysqlDB.Select("title", "start", "finish").Model(&Tour{})).Find(&mod).Error; err != nil {
		return nil, err
	}

	for _, booking := range mod {
		booking.Payment = Payment{}
		data = append(data, *booking.ToEntity())
	}

	return data, nil
}

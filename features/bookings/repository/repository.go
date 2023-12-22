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

	if flt.Pagination.Limit != 0 {
		qry = qry.Limit(flt.Pagination.Limit)
	}

	if flt.Pagination.Start != 0 {
		qry = qry.Offset(flt.Pagination.Start)
	}

	qry.Count(&totalData)

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

	modBooking.Code = 0
	modBooking.User = User{}
	modBooking.Tour = Tour{}
	modBooking.Status = ""
	modBooking.Detail = nil

	return modBooking.ToEntity(), nil
}

func (repo *bookingRepository) Update(ctx context.Context, code int, data bookings.Booking) (*bookings.Booking, error) {
	var modNewBooking = new(Booking)
	modNewBooking.FromEntity(data)

	var modOldBooking = new(Booking)
	if err := repo.mysqlDB.WithContext(ctx).Where(&Booking{Code: code}).Joins("User").First(modOldBooking).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("not found: booking not found")
		}
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

	if modNewBooking.Payment.Bank != "" {
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

	if modNewBooking.Status != "" {
		switch modNewBooking.Status {
		case "refunded":
			err := tx.WithContext(ctx).Transaction(func(txTour *gorm.DB) error {
				return txTour.WithContext(ctx).Model(&Tour{}).Where(&Tour{Id: modOldBooking.TourId}).Update("available", gorm.Expr("available + ?", len(modOldBooking.Detail))).Error
			})
			if err != nil {
				tx.Rollback()
				return nil, err
			}
		case "approved":
			err := tx.WithContext(ctx).Transaction(func(txTour *gorm.DB) error {
				return txTour.WithContext(ctx).Model(&Tour{}).Where(&Tour{Id: modOldBooking.TourId}).Update("available", gorm.Expr("available - ?", len(modOldBooking.Detail))).Error
			})
			if err != nil {
				tx.Rollback()
				return nil, err
			}
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

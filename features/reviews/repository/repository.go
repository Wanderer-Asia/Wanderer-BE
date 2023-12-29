package repository

import (
	"errors"
	"strings"
	"wanderer/features/reviews"

	"gorm.io/gorm"
)

func NewReviewRepository(mysqlDB *gorm.DB) reviews.Repository {
	return &reviewRepository{
		mysqlDB: mysqlDB,
	}
}

type reviewRepository struct {
	mysqlDB *gorm.DB
}

func (repo *reviewRepository) Create(userId uint, newReview reviews.Review) error {
	var model = new(Review)
	model.FromEntity(newReview)
	model.UserId = userId

	var exist int64
	if err := repo.mysqlDB.Model(&Review{}).Where(&Review{TourId: model.TourId, UserId: userId}).Count(&exist).Error; err != nil {
		return err
	}

	if exist != 0 {
		return errors.New("used: review already exist")
	}

	if err := repo.mysqlDB.Create(model).Error; err != nil {
		if strings.Contains(err.Error(), "1452") && strings.Contains(err.Error(), "tour") {
			return errors.New("not found: tour not found")
		}

		return err
	}

	var averageRating float32
	if err := repo.mysqlDB.Model(&Review{}).Where("tour_id = ?", model.TourId).Select("AVG(rating)").Scan(&averageRating).Error; err != nil {
		return err
	}

	if err := repo.mysqlDB.Model(&Tour{}).Where(&Tour{Id: model.TourId}).Update("rating", averageRating).Error; err != nil {
		return err
	}

	return nil
}

func (repo *reviewRepository) GetTourById(tourId uint) (*reviews.Tour, error) {
	var model = new(Review)
	model.TourId = tourId

	var tour = new(reviews.Tour)
	if err := repo.mysqlDB.Model(&Tour{}).Where(&Tour{Id: model.TourId}).First(&tour).Error; err != nil {
		return nil, err
	}

	return tour, nil
}

func (repo *reviewRepository) IsBooking(tourId uint, userId uint) bool {
	var model = new(Review)
	model.TourId = tourId
	model.UserId = userId

	var booking = new(reviews.Booking)
	if err := repo.mysqlDB.Model(&Booking{}).Where(&Booking{TourId: model.TourId}, &Booking{UserId: model.UserId}).First(&booking).Error; err != nil {
		return false
	}

	return true
}

func (repo *reviewRepository) IsApproved(tourId uint, userId uint) bool {
	var model = new(Review)
	model.TourId = tourId
	model.UserId = userId

	var booking = new(reviews.Booking)
	if err := repo.mysqlDB.Model(&Booking{}).Where(&Booking{TourId: model.TourId}, &Booking{UserId: model.UserId}, &Booking{Status: "approved"}).First(&booking).Error; err != nil {
		return false
	}

	return true
}

package repository

import (
	"context"
	"wanderer/features/reviews"
	"wanderer/features/tours"

	"gorm.io/gorm"
)

func NewReviewRepository(mysqlDB *gorm.DB, tourService tours.Service) reviews.Repository {
	return &reviewRepository{
		mysqlDB:     mysqlDB,
		tourService: tourService,
	}
}

type reviewRepository struct {
	mysqlDB     *gorm.DB
	tourService tours.Service
}

func (repo *reviewRepository) Create(userId uint, newReview reviews.Review) error {
	var model = new(Review)
	model.FromEntity(newReview)
	model.UserId = userId

	if err := repo.mysqlDB.Create(model).Error; err != nil {
		return err
	}

	var averageRating float32
	if err := repo.mysqlDB.Model(&Review{}).Where("tour_id = ?", model.TourId).Select("AVG(rating)").Scan(&averageRating).Error; err != nil {
		return err
	}

	if err := repo.tourService.UpdateRating(context.Background(), model.TourId, tours.Tour{Rating: averageRating}); err != nil {
		return err
	}

	return nil
}

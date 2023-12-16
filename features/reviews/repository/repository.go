package repository

import (
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

	if err := repo.mysqlDB.Create(model).Error; err != nil {
		return err
	}

	return nil
}

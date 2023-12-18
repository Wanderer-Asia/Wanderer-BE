package service

import (
	"errors"
	"wanderer/features/reviews"
)

func NewReviewService(repo reviews.Repository) reviews.Service {
	return &reviewService{
		repo: repo,
	}
}

type reviewService struct {
	repo reviews.Repository
}

func (srv *reviewService) Create(userId uint, newReview reviews.Review) error {
	if newReview.Text == "" {
		return errors.New("validate: review field can't be empty")
	}
	if newReview.Rating == 0 {
		return errors.New("validate: rating must be filled")
	}
	if newReview.Rating > 5.0 {
		return errors.New("validate: rating is between 1 to 5")
	}

	if err := srv.repo.Create(userId, newReview); err != nil {
		return err
	}

	return nil
}

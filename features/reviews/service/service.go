package service

import (
	"errors"
	"time"
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

	tour, err := srv.repo.GetTourById(newReview.TourId)
	if err != nil {
		return err
	}

	if time.Now().Before(tour.Start) {
		return errors.New("cannot create review: tour has not started yet")
	}

	if time.Now().Before(tour.Finish) {
		return errors.New("cannot create review: tour has not finished yet")
	}

	if !srv.repo.IsBooking(newReview.TourId, userId) {
		return errors.New("cannot create review: you have not booked the tour yet")
	}

	if !srv.repo.IsApproved(newReview.TourId, userId) {
		return errors.New("cannot create review: your transaction has not finished or has been canceled")
	}

	if err := srv.repo.Create(userId, newReview); err != nil {
		return err
	}

	return nil
}

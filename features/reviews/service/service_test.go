package service_test

import (
	"errors"
	"testing"
	"time"
	"wanderer/features/reviews"
	"wanderer/features/reviews/mocks"
	"wanderer/features/reviews/service"

	"github.com/stretchr/testify/assert"
)

func TestReviewServiceCreate(t *testing.T) {
	var repo = mocks.NewRepository(t)
	var srv = service.NewReviewService(repo)

	t.Run("invalid review", func(t *testing.T) {
		var caseData = reviews.Review{
			Text:   "",
			Rating: 4.8,
		}

		err := srv.Create(uint(1), caseData)

		assert.ErrorContains(t, err, "review")
	})

	t.Run("empty rating", func(t *testing.T) {
		var caseData = reviews.Review{
			Text:   "Good",
			Rating: 0,
		}

		err := srv.Create(uint(1), caseData)

		assert.ErrorContains(t, err, "rating")
	})

	t.Run("invalid rating", func(t *testing.T) {
		var caseData = reviews.Review{
			Text:   "Good",
			Rating: 5.1,
		}

		err := srv.Create(uint(1), caseData)

		assert.ErrorContains(t, err, "rating")
	})

	t.Run("Error get tour by id", func(t *testing.T) {
		var caseData = reviews.Review{
			Text:   "Good",
			Rating: 4.9,
		}

		repo.On("GetTourById", uint(0)).Return(nil, errors.New("some error from repository")).Once()

		err := srv.Create(uint(1), caseData)

		assert.ErrorContains(t, err, "some error from repository")

		repo.AssertExpectations(t)
	})

	t.Run("tour has not started yet", func(t *testing.T) {
		var caseData = reviews.Review{
			Text:   "Good",
			Rating: 4.9,
			TourId: 1,
		}

		startTime := time.Now().Add(time.Hour)
		var tour = &reviews.Tour{
			Id:    uint(1),
			Start: startTime,
		}

		repo.On("GetTourById", uint(1)).Return(tour, nil).Once()

		err := srv.Create(uint(1), caseData)

		assert.ErrorContains(t, err, "cannot create review: tour has not started yet")

		repo.AssertExpectations(t)
	})

	t.Run("tour has not finished yet", func(t *testing.T) {
		var caseData = reviews.Review{
			Text:   "Good",
			Rating: 4.9,
			TourId: 1,
		}

		finishTime := time.Now().Add(time.Hour)
		var tour = &reviews.Tour{
			Id:     uint(1),
			Finish: finishTime,
		}

		repo.On("GetTourById", uint(1)).Return(tour, nil).Once()

		err := srv.Create(uint(1), caseData)

		assert.ErrorContains(t, err, "cannot create review: tour has not finished yet")

		repo.AssertExpectations(t)
	})

	t.Run("tour has not booked yet", func(t *testing.T) {
		var caseData = reviews.Review{
			Text:   "Good",
			Rating: 4.9,
			TourId: 1,
		}

		finishTime := time.Now()
		var tour = &reviews.Tour{
			Id:     uint(1),
			Finish: finishTime,
		}

		repo.On("GetTourById", uint(1)).Return(tour, nil).Once()

		repo.On("IsBooking", caseData.TourId, uint(1)).Return(false).Once()

		err := srv.Create(uint(1), caseData)

		assert.ErrorContains(t, err, "cannot create review: you have not booked the tour yet")

		repo.AssertExpectations(t)
	})

	t.Run("transaction not approved", func(t *testing.T) {
		var caseData = reviews.Review{
			Text:   "Good",
			Rating: 4.9,
			TourId: 1,
		}

		finishTime := time.Now()
		var tour = &reviews.Tour{
			Id:     uint(1),
			Finish: finishTime,
		}

		repo.On("GetTourById", uint(1)).Return(tour, nil).Once()

		repo.On("IsBooking", caseData.TourId, uint(1)).Return(true).Once()

		repo.On("IsApproved", caseData.TourId, uint(1)).Return(false).Once()

		err := srv.Create(uint(1), caseData)

		assert.ErrorContains(t, err, "cannot create review: your transaction has not finished or has been canceled")

		repo.AssertExpectations(t)
	})

	t.Run("error from repository", func(t *testing.T) {
		var caseData = reviews.Review{
			Text:   "Good",
			Rating: 4.9,
			TourId: 1,
		}

		finishTime := time.Now()
		var tour = &reviews.Tour{
			Id:     uint(1),
			Finish: finishTime,
		}

		repo.On("GetTourById", uint(1)).Return(tour, nil).Once()

		repo.On("IsBooking", caseData.TourId, uint(1)).Return(true).Once()

		repo.On("IsApproved", caseData.TourId, uint(1)).Return(true).Once()

		repo.On("Create", uint(1), caseData).Return(errors.New("some error from repository")).Once()

		err := srv.Create(uint(1), caseData)

		assert.ErrorContains(t, err, "some error from repository")

		repo.AssertExpectations(t)
	})

	t.Run("success case", func(t *testing.T) {
		var caseData = reviews.Review{
			Text:   "Good",
			Rating: 4.9,
			TourId: 1,
		}

		finishTime := time.Now()
		var tour = &reviews.Tour{
			Id:     uint(1),
			Finish: finishTime,
		}

		repo.On("GetTourById", uint(1)).Return(tour, nil).Once()

		repo.On("IsBooking", caseData.TourId, uint(1)).Return(true).Once()

		repo.On("IsApproved", caseData.TourId, uint(1)).Return(true).Once()

		repo.On("Create", uint(1), caseData).Return(nil).Once()

		err := srv.Create(uint(1), caseData)

		assert.NoError(t, err)

		repo.AssertExpectations(t)
	})
}

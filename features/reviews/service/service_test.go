package service_test

import (
	"errors"
	"testing"
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

	t.Run("error from repository", func(t *testing.T) {
		var caseData = reviews.Review{
			Text:   "Good",
			Rating: 4.9,
		}

		repo.On("Create", uint(1), caseData).Return(errors.New("some error from repository")).Once()

		err := srv.Create(uint(1), caseData)

		assert.ErrorContains(t, err, "some error from repository")

		repo.AssertExpectations(t)
	})

	t.Run("success case", func(t *testing.T) {
		var caseData = reviews.Review{
			Text:   "Good",
			Rating: 4.9,
		}

		repo.On("Create", uint(1), caseData).Return(nil).Once()

		err := srv.Create(uint(1), caseData)

		assert.NoError(t, err)

		repo.AssertExpectations(t)
	})
}

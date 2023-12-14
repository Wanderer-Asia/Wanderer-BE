package service_test

import (
	"errors"
	"testing"
	"wanderer/features/airlines"
	"wanderer/features/airlines/mocks"
	"wanderer/features/airlines/service"

	"github.com/stretchr/testify/assert"
)

func TestUserServiceRegister(t *testing.T) {
	var repo = mocks.NewRepository(t)
	var srv = service.NewAirlineService(repo)

	t.Run("invalid name", func(t *testing.T) {
		var caseData = airlines.Airline{
			Name:     "",
			ImageUrl: "test",
		}

		err := srv.Create(caseData)

		assert.ErrorContains(t, err, "name")
	})

	t.Run("error from repository", func(t *testing.T) {
		var caseData = airlines.Airline{
			Name:     "Test Air",
			ImageUrl: "test",
		}

		repo.On("Create", caseData).Return(errors.New("some error from repository")).Once()

		err := srv.Create(caseData)

		assert.ErrorContains(t, err, "some error from repository")

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		var caseData = airlines.Airline{
			Name:     "Test Air",
			ImageUrl: "test",
		}

		repo.On("Create", caseData).Return(nil).Once()

		err := srv.Create(caseData)

		assert.NoError(t, err)

		repo.AssertExpectations(t)
	})
}

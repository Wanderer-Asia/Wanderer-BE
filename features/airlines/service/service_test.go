package service_test

import (
	"errors"
	"testing"
	"wanderer/features/airlines"
	"wanderer/features/airlines/mocks"
	"wanderer/features/airlines/service"

	"github.com/stretchr/testify/assert"
)

func TestAirlineServiceCreate(t *testing.T) {
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

func TestAirlineServiceGetAll(t *testing.T) {
	var repo = mocks.NewRepository(t)
	var srv = service.NewAirlineService(repo)

	t.Run("error from repository", func(t *testing.T) {
		repo.On("GetAll").Return(nil, errors.New("some error from repository")).Once()

		result, err := srv.GetAll()

		assert.ErrorContains(t, err, "some error from repository")
		assert.Nil(t, result)

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		var caseData = []airlines.Airline{
			{
				Id:       1,
				Name:     "Test Air",
				ImageUrl: "test",
			},
			{
				Id:       2,
				Name:     "Cek Air",
				ImageUrl: "test",
			},
		}

		repo.On("GetAll").Return(caseData, nil).Once()

		result, err := srv.GetAll()

		assert.NoError(t, err)
		assert.Equal(t, len(caseData), len(result))

		repo.AssertExpectations(t)
	})
}

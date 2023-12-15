package service_test

import (
	"errors"
	"testing"
	"wanderer/features/facilities"
	"wanderer/features/facilities/mocks"
	"wanderer/features/facilities/service"
	"wanderer/helpers/filters"

	"github.com/stretchr/testify/assert"
)

func TestFacilityServiceCreate(t *testing.T) {
	var repo = mocks.NewRepository(t)
	var srv = service.NewFacilityService(repo)

	t.Run("invalid name", func(t *testing.T) {
		var caseData = facilities.Facility{
			Name: "",
		}

		err := srv.Create(caseData)

		assert.ErrorContains(t, err, "name")
	})

	t.Run("error from repository", func(t *testing.T) {
		var caseData = facilities.Facility{
			Name: "Test Facility",
		}

		repo.On("Create", caseData).Return(errors.New("some error from repository")).Once()

		err := srv.Create(caseData)

		assert.ErrorContains(t, err, "some error from repository")

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		var caseData = facilities.Facility{
			Name: "Test Facility",
		}

		repo.On("Create", caseData).Return(nil).Once()

		err := srv.Create(caseData)

		assert.NoError(t, err)

		repo.AssertExpectations(t)
	})
}

func TestFacilityServiceGetAll(t *testing.T) {
	var repo = mocks.NewRepository(t)
	var srv = service.NewFacilityService(repo)

	t.Run("error from repository", func(t *testing.T) {
		repo.On("GetAll", filters.Filter{}).Return(nil, errors.New("some error from repository")).Once()

		result, err := srv.GetAll(filters.Filter{})

		assert.ErrorContains(t, err, "some error from repository")
		assert.Nil(t, result)

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		var caseData = []facilities.Facility{
			{
				Id:   1,
				Name: "Test Facility",
			},
			{
				Id:   2,
				Name: "Test Facility 2",
			},
			{
				Id:   3,
				Name: "Test Facility 3",
			},
		}

		repo.On("GetAll", filters.Filter{}).Return(caseData, nil).Once()

		result, err := srv.GetAll(filters.Filter{})

		assert.NoError(t, err)
		assert.Equal(t, len(caseData), len(result))

		repo.AssertExpectations(t)
	})
}

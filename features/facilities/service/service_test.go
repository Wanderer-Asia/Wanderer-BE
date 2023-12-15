package service_test

import (
	"errors"
	"testing"
	"wanderer/features/facilities"
	"wanderer/features/facilities/mocks"
	"wanderer/features/facilities/service"

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

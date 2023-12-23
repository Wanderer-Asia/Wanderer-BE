package service_test

import (
	"context"
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

func TestFacilityServiceUpdate(t *testing.T) {
	var repo = mocks.NewRepository(t)
	var srv = service.NewFacilityService(repo)

	t.Run("invalid name", func(t *testing.T) {
		var caseData = facilities.Facility{
			Name: "",
		}

		err := srv.Update(uint(1), caseData)

		assert.ErrorContains(t, err, "name")
	})

	t.Run("invalid id", func(t *testing.T) {
		var caseData = facilities.Facility{
			Name: "test",
		}

		err := srv.Update(uint(0), caseData)

		assert.ErrorContains(t, err, "id")
	})

	t.Run("error from repository", func(t *testing.T) {
		var caseData = facilities.Facility{
			Name: "Test Facility",
		}

		repo.On("Update", uint(1), caseData).Return(errors.New("some error from repository")).Once()

		err := srv.Update(uint(1), caseData)

		assert.ErrorContains(t, err, "some error from repository")

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		var caseData = facilities.Facility{
			Name: "Test Facility",
		}

		repo.On("Update", uint(1), caseData).Return(nil).Once()

		err := srv.Update(uint(1), caseData)

		assert.NoError(t, err)

		repo.AssertExpectations(t)
	})
}

func TestFacilityServiceDelete(t *testing.T) {
	var repo = mocks.NewRepository(t)
	var srv = service.NewFacilityService(repo)

	t.Run("invalid id", func(t *testing.T) {

		err := srv.Delete(uint(0))

		assert.ErrorContains(t, err, "id")
	})

	t.Run("error from repository", func(t *testing.T) {

		repo.On("Delete", uint(1)).Return(errors.New("some error from repository")).Once()

		err := srv.Delete(uint(1))

		assert.ErrorContains(t, err, "some error from repository")

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {

		repo.On("Delete", uint(1)).Return(nil).Once()

		err := srv.Delete(uint(1))

		assert.NoError(t, err)

		repo.AssertExpectations(t)
	})
}

func TestFacilityServiceImport(t *testing.T) {
	var repo = mocks.NewRepository(t)
	var srv = service.NewFacilityService(repo)
	var ctx = context.Background()
	var caseData = []facilities.Facility{
		{
			Id:   1,
			Name: "Test 1",
		},
		{
			Id:   2,
			Name: "Test 2",
		},
	}

	t.Run("error from repository", func(t *testing.T) {
		repo.On("Import", ctx, caseData).Return(errors.New("some error from repository")).Once()

		err := srv.Import(ctx, caseData)

		assert.ErrorContains(t, err, "some error from repository")

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		repo.On("Import", ctx, caseData).Return(nil).Once()

		err := srv.Import(ctx, caseData)

		assert.NoError(t, err)

		repo.AssertExpectations(t)
	})
}

package service_test

import (
	"context"
	"errors"
	"testing"
	"wanderer/features/airlines"
	"wanderer/features/airlines/mocks"
	"wanderer/features/airlines/service"
	"wanderer/helpers/filters"

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
		repo.On("GetAll", filters.Filter{}).Return(nil, errors.New("some error from repository")).Once()

		result, err := srv.GetAll(filters.Filter{})

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

		repo.On("GetAll", filters.Filter{}).Return(caseData, nil).Once()

		result, err := srv.GetAll(filters.Filter{})

		assert.NoError(t, err)
		assert.Equal(t, len(caseData), len(result))

		repo.AssertExpectations(t)
	})
}

func TestAirlineServiceUpdate(t *testing.T) {
	var repo = mocks.NewRepository(t)
	var srv = service.NewAirlineService(repo)

	t.Run("invalid airline id", func(t *testing.T) {
		var caseData = airlines.Airline{
			Name:     "",
			ImageUrl: "test",
		}

		err := srv.Update(uint(0), caseData)

		assert.ErrorContains(t, err, "id")
	})

	t.Run("invalid name", func(t *testing.T) {
		var caseData = airlines.Airline{
			Name:     "",
			ImageUrl: "test",
		}

		err := srv.Update(uint(1), caseData)

		assert.ErrorContains(t, err, "name")
	})

	t.Run("error from repository", func(t *testing.T) {
		var caseData = airlines.Airline{
			Name:     "Test Air",
			ImageUrl: "test",
		}

		repo.On("Update", uint(1), caseData).Return(errors.New("some error from repository")).Once()

		err := srv.Update(uint(1), caseData)

		assert.ErrorContains(t, err, "some error from repository")

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		var caseData = airlines.Airline{
			Name:     "Test Air",
			ImageUrl: "test",
		}

		repo.On("Update", uint(1), caseData).Return(nil).Once()

		err := srv.Update(uint(1), caseData)

		assert.NoError(t, err)

		repo.AssertExpectations(t)
	})
}

func TestAirlineServiceDelete(t *testing.T) {
	var repo = mocks.NewRepository(t)
	var srv = service.NewAirlineService(repo)

	t.Run("invalid airline id", func(t *testing.T) {
		var id = uint(0)

		err := srv.Delete(id)

		assert.ErrorContains(t, err, "airline id")
	})

	t.Run("error from repository", func(t *testing.T) {
		var id = uint(1)

		repo.On("Delete", id).Return(errors.New("some error from repository")).Once()

		err := srv.Delete(id)

		assert.ErrorContains(t, err, "some error from repository")

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		var id = uint(1)

		repo.On("Delete", id).Return(nil).Once()

		err := srv.Delete(1)
		assert.Nil(t, err)

		repo.AssertExpectations(t)
	})
}

func TestAirlineServiceImport(t *testing.T) {
	var repo = mocks.NewRepository(t)
	var srv = service.NewAirlineService(repo)
	var ctx = context.Background()
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

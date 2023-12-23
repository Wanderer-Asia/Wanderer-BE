package service

import (
	"context"
	"errors"
	"strings"
	"testing"
	"wanderer/features/locations"
	"wanderer/features/locations/mocks"
	"wanderer/helpers/filters"

	"github.com/stretchr/testify/assert"
)

func TestLocationServiceGetAll(t *testing.T) {
	repo := mocks.NewRepository(t)
	srv := NewLocationService(repo)
	ctx := context.Background()

	t.Run("error from repository", func(t *testing.T) {
		repo.On("GetAll", ctx, filters.Filter{}).Return(nil, errors.New("some error from repository")).Once()

		result, err := srv.GetAll(ctx, filters.Filter{})

		assert.ErrorContains(t, err, "some error from repository")
		assert.Nil(t, result)

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		caseResult := []locations.Location{
			{
				Id:   1,
				Name: "Japan",
			},
			{
				Id:   2,
				Name: "Korea",
			},
		}
		repo.On("GetAll", ctx, filters.Filter{}).Return(caseResult, nil).Once()

		result, err := srv.GetAll(ctx, filters.Filter{})

		assert.NoError(t, err)
		assert.Equal(t, len(caseResult), len(result))

		repo.AssertExpectations(t)
	})
}

func TestLocationServiceCreate(t *testing.T) {
	repo := mocks.NewRepository(t)
	srv := NewLocationService(repo)
	ctx := context.Background()

	t.Run("invalid name", func(t *testing.T) {
		caseData := locations.Location{}

		err := srv.Create(ctx, caseData)

		assert.ErrorContains(t, err, "validate")
		assert.ErrorContains(t, err, "name")
	})

	t.Run("invalid image", func(t *testing.T) {
		caseData := locations.Location{
			Name: "example location",
		}

		err := srv.Create(ctx, caseData)

		assert.ErrorContains(t, err, "validate")
		assert.ErrorContains(t, err, "image")
	})

	t.Run("error from repository", func(t *testing.T) {
		caseData := locations.Location{
			Name:     "example location",
			ImageRaw: strings.NewReader("example"),
		}
		repo.On("Create", ctx, caseData).Return(errors.New("some error from repository")).Once()

		err := srv.Create(ctx, caseData)

		assert.ErrorContains(t, err, "some error from repository")
		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		caseData := locations.Location{
			Name:     "example location",
			ImageRaw: strings.NewReader("example"),
		}
		repo.On("Create", ctx, caseData).Return(nil).Once()
		err := srv.Create(ctx, caseData)

		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
}

func TestLocationServiceUpdate(t *testing.T) {
	repo := mocks.NewRepository(t)
	srv := NewLocationService(repo)
	ctx := context.Background()

	t.Run("invalid id", func(t *testing.T) {
		caseData := locations.Location{
			Name: "example location",
		}

		err := srv.Update(ctx, 0, caseData)

		assert.ErrorContains(t, err, "validate")
		assert.ErrorContains(t, err, "id")
	})

	t.Run("invalid name", func(t *testing.T) {
		caseData := locations.Location{}

		err := srv.Update(ctx, 1, caseData)

		assert.ErrorContains(t, err, "validate")
		assert.ErrorContains(t, err, "name")
	})

	t.Run("error from repository", func(t *testing.T) {
		caseData := locations.Location{
			Name:     "example location",
			ImageRaw: strings.NewReader("example"),
		}
		repo.On("Update", ctx, uint(1), caseData).Return(errors.New("some error from repository")).Once()

		err := srv.Update(ctx, 1, caseData)

		assert.ErrorContains(t, err, "some error from repository")
		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		caseData := locations.Location{
			Name:     "example location",
			ImageRaw: strings.NewReader("example"),
		}
		repo.On("Update", ctx, uint(1), caseData).Return(nil).Once()
		err := srv.Update(ctx, 1, caseData)

		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
}

func TestLocationServiceDelete(t *testing.T) {
	repo := mocks.NewRepository(t)
	srv := NewLocationService(repo)
	ctx := context.Background()

	t.Run("invalid id", func(t *testing.T) {

		err := srv.Delete(ctx, 0)

		assert.ErrorContains(t, err, "validate")
		assert.ErrorContains(t, err, "id")
	})

	t.Run("error from repository", func(t *testing.T) {
		repo.On("Delete", ctx, uint(1)).Return(errors.New("some error from repository")).Once()

		err := srv.Delete(ctx, 1)

		assert.ErrorContains(t, err, "some error from repository")
		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		repo.On("Delete", ctx, uint(1)).Return(nil).Once()
		err := srv.Delete(ctx, 1)

		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
}

func TestLocationServiceGetByLocation(t *testing.T) {
	repo := mocks.NewRepository(t)
	srv := NewLocationService(repo)
	ctx := context.Background()

	t.Run("invalid id", func(t *testing.T) {
		res, err := srv.GetDetail(ctx, uint(0))

		assert.ErrorContains(t, err, "id")
		assert.Nil(t, res)
	})

	t.Run("error from repository", func(t *testing.T) {
		repo.On("GetDetail", ctx, uint(1)).Return(nil, errors.New("error from repository")).Once()
		res, err := srv.GetDetail(ctx, uint(1))

		assert.ErrorContains(t, err, "error from repository")
		assert.Nil(t, res)
	})

	t.Run("success case", func(t *testing.T) {
		var caseData = &locations.Location{
			Id:       1,
			Name:     "Tokyo",
			ImageRaw: strings.NewReader("example"),
		}
		repo.On("GetDetail", ctx, uint(1)).Return(caseData, nil).Once()
		res, err := srv.GetDetail(ctx, uint(1))

		assert.NoError(t, err)
		assert.Equal(t, caseData, res)

		repo.AssertExpectations(t)
	})
}

func TestLocationServiceImport(t *testing.T) {
	var repo = mocks.NewRepository(t)
	var srv = NewLocationService(repo)
	var ctx = context.Background()
	var caseData = []locations.Location{
		{
			Id:       1,
			Name:     "Test 1",
			ImageUrl: "test",
		},
		{
			Id:       2,
			Name:     "Tes 2",
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

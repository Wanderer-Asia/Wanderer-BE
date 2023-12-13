package service

import (
	"context"
	"errors"
	"testing"
	"wanderer/features/locations"
	"wanderer/features/locations/mocks"

	"github.com/stretchr/testify/assert"
)

func TestLocationServiceGetAll(t *testing.T) {
	repo := mocks.NewRepository(t)
	srv := NewLocationService(repo)
	ctx := context.Background()

	t.Run("error from repository", func(t *testing.T) {
		repo.On("GetAll", ctx).Return(nil, errors.New("some error from repository")).Once()

		result, err := srv.GetAll(ctx)

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
		repo.On("GetAll", ctx).Return(caseResult, nil).Once()

		result, err := srv.GetAll(ctx)

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

	t.Run("error from repository", func(t *testing.T) {
		caseData := locations.Location{
			Name: "example location",
		}
		repo.On("Create", ctx, caseData).Return(errors.New("some error from repository")).Once()

		err := srv.Create(ctx, caseData)

		assert.ErrorContains(t, err, "some error from repository")
		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		caseData := locations.Location{
			Name: "example location",
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
			Name: "example location",
		}
		repo.On("Update", ctx, uint(1), caseData).Return(errors.New("some error from repository")).Once()

		err := srv.Update(ctx, 1, caseData)

		assert.ErrorContains(t, err, "some error from repository")
		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		caseData := locations.Location{
			Name: "example location",
		}
		repo.On("Update", ctx, uint(1), caseData).Return(nil).Once()
		err := srv.Update(ctx, 1, caseData)

		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
}

func TestLocationServiceDelete(t *testing.T) {

}

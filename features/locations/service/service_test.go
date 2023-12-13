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

}

func TestLocationServiceUpdate(t *testing.T) {

}

func TestLocationServiceDelete(t *testing.T) {

}

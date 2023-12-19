package service

import (
	"context"
	"errors"
	"testing"
	"time"
	"wanderer/features/bookings"
	"wanderer/features/bookings/mocks"
	"wanderer/features/tours"
	"wanderer/features/users"

	"github.com/stretchr/testify/assert"
)

func TestBookingServiceCreate(t *testing.T) {
	repo := mocks.NewRepository(t)
	srv := NewBookingService(repo)
	ctx := context.Background()

	data := bookings.Booking{
		Total:     10000,
		Status:    "pending",
		BookedAt:  time.Now(),
		DeletedAt: time.Now(),
		User: users.User{
			Id: 1,
		},
		Tour: tours.Tour{
			Id: 1,
		},
		Detail: []bookings.Detail{
			{
				DocumentNumber: "123",
				Greeting:       "mr",
				Name:           "maman",
				Nationality:    "indonesia",
				DOB:            time.Now(),
			},
		},
		Payment: bookings.Payment{
			Bank: "bri",
		},
	}

	t.Run("invalid tour id", func(t *testing.T) {
		caseData := data
		tourId := caseData.Tour.Id
		caseData.Tour.Id = 0
		result, err := srv.Create(ctx, caseData)

		assert.ErrorContains(t, err, "validate")
		assert.ErrorContains(t, err, "tour id")
		assert.Nil(t, result)

		caseData.Tour.Id = tourId
	})

	t.Run("invalid user id", func(t *testing.T) {
		caseData := data
		userId := caseData.User.Id
		caseData.User.Id = 0
		result, err := srv.Create(ctx, caseData)

		assert.ErrorContains(t, err, "validate")
		assert.ErrorContains(t, err, "user id")
		assert.Nil(t, result)

		caseData.User.Id = userId
	})

	t.Run("empty pasenger", func(t *testing.T) {
		caseData := data
		detail := caseData.Detail
		caseData.Detail = nil
		result, err := srv.Create(ctx, caseData)

		assert.ErrorContains(t, err, "validate")
		assert.ErrorContains(t, err, "passenger data")
		assert.Nil(t, result)

		caseData.Detail = detail
	})

	t.Run("invalid pasenger document number", func(t *testing.T) {
		caseData := data
		documentNumber := caseData.Detail[0].DocumentNumber
		caseData.Detail[0].DocumentNumber = ""
		result, err := srv.Create(ctx, caseData)

		assert.ErrorContains(t, err, "validate")
		assert.ErrorContains(t, err, "document number")
		assert.Nil(t, result)

		caseData.Detail[0].DocumentNumber = documentNumber
	})

	t.Run("invalid pasenger greeting", func(t *testing.T) {
		caseData := data
		greeting := caseData.Detail[0].Greeting
		caseData.Detail[0].Greeting = ""
		result, err := srv.Create(ctx, caseData)

		assert.ErrorContains(t, err, "validate")
		assert.ErrorContains(t, err, "greeting")
		assert.Nil(t, result)

		caseData.Detail[0].Greeting = greeting
	})

	t.Run("invalid pasenger name", func(t *testing.T) {
		caseData := data
		name := caseData.Detail[0].Name
		caseData.Detail[0].Name = ""
		result, err := srv.Create(ctx, caseData)

		assert.ErrorContains(t, err, "validate")
		assert.ErrorContains(t, err, "name")
		assert.Nil(t, result)

		caseData.Detail[0].Name = name
	})

	t.Run("invalid pasenger nationality", func(t *testing.T) {
		caseData := data
		nationality := caseData.Detail[0].Nationality
		caseData.Detail[0].Nationality = ""
		result, err := srv.Create(ctx, caseData)

		assert.ErrorContains(t, err, "validate")
		assert.ErrorContains(t, err, "nationality")
		assert.Nil(t, result)

		caseData.Detail[0].Nationality = nationality
	})

	t.Run("invalid pasenger date of birth", func(t *testing.T) {
		caseData := data
		dob := caseData.Detail[0].DOB
		caseData.Detail[0].DOB = time.Time{}
		result, err := srv.Create(ctx, caseData)

		assert.ErrorContains(t, err, "validate")
		assert.ErrorContains(t, err, "date of birth")
		assert.Nil(t, result)

		caseData.Detail[0].DOB = dob
	})

	t.Run("error from repository", func(t *testing.T) {
		caseData := data
		repo.On("Create", ctx, caseData).Return(nil, errors.New("some error from repository")).Once()

		result, err := srv.Create(ctx, caseData)

		assert.ErrorContains(t, err, "some error from repository")
		assert.Nil(t, result)

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		caseData := data
		repo.On("Create", ctx, caseData).Return(&caseData, nil).Once()

		result, err := srv.Create(ctx, caseData)

		assert.NoError(t, err)
		assert.Equal(t, &caseData, result)

		repo.AssertExpectations(t)
	})
}

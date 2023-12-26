package service

import (
	"context"
	"errors"
	"testing"
	"time"
	"wanderer/features/bookings"
	"wanderer/features/bookings/mocks"
	"wanderer/helpers/filters"

	"github.com/stretchr/testify/assert"
)

func TestBookingServiceGetAll(t *testing.T) {
	repo := mocks.NewRepository(t)
	srv := NewBookingService(repo)
	ctx := context.Background()

	data := []bookings.Booking{
		{
			Code:      123,
			Total:     10000,
			Status:    "pending",
			BookedAt:  time.Now(),
			DeletedAt: time.Now(),
			User: bookings.User{
				Id: 1,
			},
			Tour: bookings.Tour{
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
		},
		{
			Code:      234,
			Total:     10000,
			Status:    "pending",
			BookedAt:  time.Now(),
			DeletedAt: time.Now(),
			User: bookings.User{
				Id: 1,
			},
			Tour: bookings.Tour{
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
		},
	}

	t.Run("error from repository", func(t *testing.T) {
		repo.On("GetAll", ctx, filters.Filter{}).Return(nil, 0, errors.New("some error from repository")).Once()

		result, totalData, err := srv.GetAll(ctx, filters.Filter{})

		assert.ErrorContains(t, err, "some error from repository")
		assert.Equal(t, 0, totalData)
		assert.Nil(t, result)

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		caseData := data
		repo.On("GetAll", ctx, filters.Filter{}).Return(caseData, 2, nil).Once()

		result, totalData, err := srv.GetAll(ctx, filters.Filter{})

		assert.NoError(t, err)
		assert.Equal(t, 2, totalData)
		assert.Equal(t, caseData, result)

		repo.AssertExpectations(t)
	})
}

func TestBookingServiceGetDetail(t *testing.T) {
	repo := mocks.NewRepository(t)
	srv := NewBookingService(repo)
	ctx := context.Background()

	data := bookings.Booking{
		Total:     10000,
		Status:    "pending",
		BookedAt:  time.Now(),
		DeletedAt: time.Now(),
		User: bookings.User{
			Id: 1,
		},
		Tour: bookings.Tour{
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

	t.Run("invalid booking code", func(t *testing.T) {
		result, err := srv.GetDetail(ctx, 0)

		assert.ErrorContains(t, err, "validate")
		assert.ErrorContains(t, err, "booking code")
		assert.Nil(t, result)
	})

	t.Run("error from repository", func(t *testing.T) {
		repo.On("GetDetail", ctx, 123).Return(nil, errors.New("some error from repository")).Once()

		result, err := srv.GetDetail(ctx, 123)

		assert.ErrorContains(t, err, "some error from repository")
		assert.Nil(t, result)

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		caseData := data
		repo.On("GetDetail", ctx, 123).Return(&caseData, nil).Once()

		result, err := srv.GetDetail(ctx, 123)

		assert.NoError(t, err)
		assert.Equal(t, &caseData, result)

		repo.AssertExpectations(t)
	})
}

func TestBookingServiceCreate(t *testing.T) {
	repo := mocks.NewRepository(t)
	srv := NewBookingService(repo)
	ctx := context.Background()

	data := bookings.Booking{
		Total:     10000,
		Status:    "pending",
		BookedAt:  time.Now(),
		DeletedAt: time.Now(),
		User: bookings.User{
			Id: 1,
		},
		Tour: bookings.Tour{
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

	t.Run("invalid payment method", func(t *testing.T) {
		caseData := data
		paymentBank := caseData.Payment.Bank
		caseData.Payment.Bank = ""
		result, err := srv.Create(ctx, caseData)

		assert.ErrorContains(t, err, "validate")
		assert.ErrorContains(t, err, "payment method")
		assert.Nil(t, result)

		caseData.Payment.Bank = paymentBank
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

func TestBookingServiceUpdate(t *testing.T) {
	repo := mocks.NewRepository(t)
	srv := NewBookingService(repo)
	ctx := context.Background()

	data := bookings.Booking{
		Total:     10000,
		Status:    "pending",
		BookedAt:  time.Now(),
		DeletedAt: time.Now(),
		User: bookings.User{
			Id: 1,
		},
		Tour: bookings.Tour{
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

	t.Run("invalid booking code", func(t *testing.T) {
		caseData := data

		result, err := srv.Update(ctx, 0, caseData)

		assert.ErrorContains(t, err, "validate")
		assert.ErrorContains(t, err, "booking code")
		assert.Nil(t, result)
	})

	t.Run("error from repository", func(t *testing.T) {
		caseData := data
		repo.On("Update", ctx, 123, caseData).Return(nil, errors.New("some error from repository")).Once()

		result, err := srv.Update(ctx, 123, caseData)

		assert.ErrorContains(t, err, "some error from repository")
		assert.Nil(t, result)

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		caseData := data
		repo.On("Update", ctx, 123, caseData).Return(&caseData, nil).Once()

		result, err := srv.Update(ctx, 123, caseData)

		assert.NoError(t, err)
		assert.Equal(t, &caseData, result)

		repo.AssertExpectations(t)
	})
}

func TestBookingServiceExport(t *testing.T) {
	repo := mocks.NewRepository(t)
	srv := NewBookingService(repo)
	ctx := context.Background()
	t.Run("Error from repository", func(t *testing.T) {
		repo.On("Export", ctx).Return(nil, errors.New("some error from repository")).Once()

		result, err := srv.Export(ctx)

		assert.ErrorContains(t, err, "some error from repository")
		assert.Nil(t, result)

		repo.AssertExpectations(t)
	})
	t.Run("Succes Case", func(t *testing.T) {
		data := []bookings.Booking{
			{
				Code:   123,
				Total:  10000,
				Status: "pending",
				User: bookings.User{
					Id: 1,
				},
				Tour: bookings.Tour{
					Id: 1,
				},
			},
			{
				Code:   234,
				Total:  10000,
				Status: "pending",
				User: bookings.User{
					Id: 1,
				},
				Tour: bookings.Tour{
					Id: 1,
				},
			},
		}

		repo.On("Export", ctx).Return(data, nil).Once()

		result, err := srv.Export(ctx)

		assert.NoError(t, err)
		assert.Equal(t, data, result)

		repo.AssertExpectations(t)
	})
}

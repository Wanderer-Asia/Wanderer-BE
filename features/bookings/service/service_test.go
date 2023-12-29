package service

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"wanderer/features/bookings"
	"wanderer/features/bookings/mocks"
	"wanderer/helpers/filters"

	"github.com/labstack/echo/v4"
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

	t.Run("tour not found", func(t *testing.T) {
		caseData := data
		repo.On("GetTourById", ctx, uint(caseData.Tour.Id)).Return(nil, errors.New("not found: tour not found")).Once()

		result, err := srv.Create(ctx, caseData)

		assert.ErrorContains(t, err, "not found")
		assert.ErrorContains(t, err, "tour")
		assert.Nil(t, result)

		repo.AssertExpectations(t)
	})

	t.Run("tour started", func(t *testing.T) {
		caseData := data
		repo.On("GetTourById", ctx, uint(caseData.Tour.Id)).Return(&bookings.Tour{Start: time.Now().Add(-1 * time.Hour)}, nil).Once()

		result, err := srv.Create(ctx, caseData)

		assert.ErrorContains(t, err, "unprocessable")
		assert.ErrorContains(t, err, "tour")
		assert.ErrorContains(t, err, "started")
		assert.Nil(t, result)

		repo.AssertExpectations(t)
	})

	t.Run("error from repository", func(t *testing.T) {
		caseData := data
		repo.On("GetTourById", ctx, uint(caseData.Tour.Id)).Return(&bookings.Tour{Start: time.Now().Add(time.Hour)}, nil).Once()
		repo.On("Create", ctx, caseData).Return(nil, errors.New("some error from repository")).Once()

		result, err := srv.Create(ctx, caseData)

		assert.ErrorContains(t, err, "some error from repository")
		assert.Nil(t, result)

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		caseData := data
		repo.On("GetTourById", ctx, uint(caseData.Tour.Id)).Return(&bookings.Tour{Start: time.Now().Add(time.Hour)}, nil).Once()
		repo.On("Create", ctx, caseData).Return(&caseData, nil).Once()

		result, err := srv.Create(ctx, caseData)

		assert.NoError(t, err)
		assert.Equal(t, &caseData, result)

		repo.AssertExpectations(t)
	})
}

func TestBookingServiceUpdateBookingStatus(t *testing.T) {
	repo := mocks.NewRepository(t)
	srv := NewBookingService(repo)
	ctx := context.Background()

	t.Run("invalid booking code", func(t *testing.T) {
		err := srv.UpdateBookingStatus(ctx, 0, "cancel")

		assert.ErrorContains(t, err, "validate")
		assert.ErrorContains(t, err, "invalid booking code")
	})

	t.Run("booking not found", func(t *testing.T) {
		repo.On("GetDetail", ctx, 123).Return(nil, errors.New("not found: booking not found")).Once()

		err := srv.UpdateBookingStatus(ctx, 123, "cancel")

		assert.ErrorContains(t, err, "not found")
		assert.ErrorContains(t, err, "booking")

		repo.AssertExpectations(t)
	})

	t.Run("invalid booking status", func(t *testing.T) {
		repoGetDetail := &bookings.Booking{}
		repo.On("GetDetail", ctx, 123).Return(repoGetDetail, nil).Once()

		err := srv.UpdateBookingStatus(ctx, 123, "pending")

		assert.ErrorContains(t, err, "validate")
		assert.ErrorContains(t, err, "invalid booking status")

		repo.AssertExpectations(t)
	})

	t.Run("cancel booking while before status isn't pending", func(t *testing.T) {
		repoGetDetail := &bookings.Booking{Status: "approved"}
		repo.On("GetDetail", ctx, 123).Return(repoGetDetail, nil).Once()

		err := srv.UpdateBookingStatus(ctx, 123, "cancel")

		assert.ErrorContains(t, err, "unprocessable")
		assert.ErrorContains(t, err, "cancel")

		repo.AssertExpectations(t)
	})

	t.Run("request refund while booking status isn't approved", func(t *testing.T) {
		repoGetDetail := &bookings.Booking{Status: "pending"}
		repo.On("GetDetail", ctx, 123).Return(repoGetDetail, nil).Once()

		err := srv.UpdateBookingStatus(ctx, 123, "refund")

		assert.ErrorContains(t, err, "unprocessable")
		assert.ErrorContains(t, err, "refund")

		repo.AssertExpectations(t)
	})

	t.Run("approve refund while refund not requested", func(t *testing.T) {
		repoGetDetail := &bookings.Booking{Status: "approved"}
		repo.On("GetDetail", ctx, 123).Return(repoGetDetail, nil).Once()

		err := srv.UpdateBookingStatus(ctx, 123, "refunded")

		assert.ErrorContains(t, err, "unprocessable")
		assert.ErrorContains(t, err, "approve refund")

		repo.AssertExpectations(t)
	})

	t.Run("error from repository", func(t *testing.T) {
		repoGetDetail := &bookings.Booking{Status: "approved"}
		repo.On("GetDetail", ctx, 123).Return(repoGetDetail, nil).Once()
		repo.On("UpdateBookingStatus", ctx, 123, "refund").Return(errors.New("some error from repository")).Once()

		err := srv.UpdateBookingStatus(ctx, 123, "refund")

		assert.ErrorContains(t, err, "some error from repository")

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		repoGetDetail := &bookings.Booking{Status: "approved"}
		repo.On("GetDetail", ctx, 123).Return(repoGetDetail, nil).Once()
		repo.On("UpdateBookingStatus", ctx, 123, "refund").Return(nil).Once()

		err := srv.UpdateBookingStatus(ctx, 123, "refund")

		assert.NoError(t, err)

		repo.AssertExpectations(t)
	})
}

func TestBookingServiceUpdatePaymentStatus(t *testing.T) {
	repo := mocks.NewRepository(t)
	srv := NewBookingService(repo)
	ctx := context.Background()

	t.Run("invalid booking code", func(t *testing.T) {
		err := srv.UpdatePaymentStatus(ctx, 0, "cancel")

		assert.ErrorContains(t, err, "validate")
		assert.ErrorContains(t, err, "invalid booking code")
	})

	t.Run("error from repository", func(t *testing.T) {
		repo.On("UpdatePaymentStatus", ctx, 123, "approved", "settlement").Return(errors.New("some error from repository")).Once()

		err := srv.UpdatePaymentStatus(ctx, 123, "settlement")

		assert.ErrorContains(t, err, "some error from repository")

		repo.AssertExpectations(t)
	})

	t.Run("payment settlement", func(t *testing.T) {
		repo.On("UpdatePaymentStatus", ctx, 123, "approved", "settlement").Return(nil).Once()

		err := srv.UpdatePaymentStatus(ctx, 123, "settlement")

		assert.NoError(t, err)

		repo.AssertExpectations(t)
	})

	t.Run("payment cancel", func(t *testing.T) {
		repo.On("UpdatePaymentStatus", ctx, 123, "cancel", "cancel").Return(nil).Once()

		err := srv.UpdatePaymentStatus(ctx, 123, "cancel")

		assert.NoError(t, err)

		repo.AssertExpectations(t)
	})

	t.Run("payment expire", func(t *testing.T) {
		repo.On("UpdatePaymentStatus", ctx, 123, "cancel", "expire").Return(nil).Once()

		err := srv.UpdatePaymentStatus(ctx, 123, "expire")

		assert.NoError(t, err)

		repo.AssertExpectations(t)
	})

	t.Run("payment capture", func(t *testing.T) {
		repo.On("UpdatePaymentStatus", ctx, 123, "pending", "capture").Return(nil).Once()

		err := srv.UpdatePaymentStatus(ctx, 123, "capture")

		assert.NoError(t, err)

		repo.AssertExpectations(t)
	})

	t.Run("payment deny", func(t *testing.T) {
		repo.On("UpdatePaymentStatus", ctx, 123, "pending", "deny").Return(nil).Once()

		err := srv.UpdatePaymentStatus(ctx, 123, "deny")

		assert.NoError(t, err)

		repo.AssertExpectations(t)
	})

	t.Run("payment pending", func(t *testing.T) {
		repo.On("UpdatePaymentStatus", ctx, 123, "pending", "pending").Return(nil).Once()

		err := srv.UpdatePaymentStatus(ctx, 123, "pending")

		assert.NoError(t, err)

		repo.AssertExpectations(t)
	})

	t.Run("invalid payment status", func(t *testing.T) {
		err := srv.UpdatePaymentStatus(ctx, 123, "tes")

		assert.ErrorContains(t, err, "validate")
		assert.ErrorContains(t, err, "invalid payment status")
	})
}

func TestBookingServiceChangePaymentMethod(t *testing.T) {
	repo := mocks.NewRepository(t)
	srv := NewBookingService(repo)
	ctx := context.Background()

	t.Run("invalid booking code", func(t *testing.T) {
		caseData := bookings.Payment{Bank: "bri"}
		result, err := srv.ChangePaymentMethod(ctx, 0, caseData)

		assert.ErrorContains(t, err, "validate")
		assert.ErrorContains(t, err, "invalid booking code")
		assert.Nil(t, result)
	})

	t.Run("empty payment method", func(t *testing.T) {
		caseData := bookings.Payment{Bank: ""}
		result, err := srv.ChangePaymentMethod(ctx, 123, caseData)

		assert.ErrorContains(t, err, "validate")
		assert.ErrorContains(t, err, "payment method")
		assert.ErrorContains(t, err, "empty")
		assert.Nil(t, result)
	})

	t.Run("booking not found", func(t *testing.T) {
		caseData := bookings.Payment{Bank: "bri"}

		repo.On("GetDetail", ctx, 123).Return(nil, errors.New("not found: booking not found")).Once()

		result, err := srv.ChangePaymentMethod(ctx, 123, caseData)

		assert.ErrorContains(t, err, "not found")
		assert.ErrorContains(t, err, "booking")
		assert.Nil(t, result)

		repo.AssertExpectations(t)
	})

	t.Run("change payment method while booking wasn't pending", func(t *testing.T) {
		caseData := bookings.Payment{Bank: "bri"}

		repoGetDetail := &bookings.Booking{Status: "approved"}
		repo.On("GetDetail", ctx, 123).Return(repoGetDetail, nil).Once()

		result, err := srv.ChangePaymentMethod(ctx, 123, caseData)

		assert.ErrorContains(t, err, "unprocessable")
		assert.ErrorContains(t, err, "payment method")
		assert.Nil(t, result)

		repo.AssertExpectations(t)
	})

	t.Run("change payment method while same method and payment isn't expired", func(t *testing.T) {
		caseData := bookings.Payment{Bank: "bri"}

		repoGetDetail := &bookings.Booking{
			Status: "pending",
			Payment: bookings.Payment{
				Status:    "pending",
				Bank:      "bri",
				ExpiredAt: time.Now().Add(time.Hour),
			},
		}
		repo.On("GetDetail", ctx, 123).Return(repoGetDetail, nil).Once()

		result, err := srv.ChangePaymentMethod(ctx, 123, caseData)

		assert.NoError(t, err)
		assert.Equal(t, &repoGetDetail.Payment, result)

		repo.AssertExpectations(t)
	})

	t.Run("error from repository", func(t *testing.T) {
		caseData := bookings.Payment{Bank: "bri"}

		repoGetDetail := &bookings.Booking{
			Status: "pending",
			Payment: bookings.Payment{
				Status:    "pending",
				Bank:      "bri",
				ExpiredAt: time.Now(),
			},
		}
		repo.On("GetDetail", ctx, 123).Return(repoGetDetail, nil).Once()

		repoChangePaymentMethod := bookings.Booking{
			Status:  "pending",
			Payment: caseData,
		}
		repo.On("ChangePaymentMethod", ctx, 123, repoChangePaymentMethod).Return(nil, errors.New("some error from repository")).Once()

		result, err := srv.ChangePaymentMethod(ctx, 123, caseData)

		assert.ErrorContains(t, err, "some error from repository")
		assert.Nil(t, result)

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		caseData := bookings.Payment{Bank: "bri"}

		repoGetDetail := &bookings.Booking{
			Status: "pending",
			Payment: bookings.Payment{
				Status:    "pending",
				Bank:      "bri",
				ExpiredAt: time.Now(),
			},
		}
		repo.On("GetDetail", ctx, 123).Return(repoGetDetail, nil).Once()

		repoChangePaymentMethod := bookings.Booking{
			Status:  "pending",
			Payment: caseData,
		}
		repo.On("ChangePaymentMethod", ctx, 123, repoChangePaymentMethod).Return(&caseData, nil).Once()

		result, err := srv.ChangePaymentMethod(ctx, 123, caseData)

		assert.NoError(t, err)
		assert.Equal(t, caseData.Bank, result.Bank)

		repo.AssertExpectations(t)
	})
}

func TestBookingServiceExport(t *testing.T) {
	repo := mocks.NewRepository(t)
	srv := NewBookingService(repo)
	// ctx := context.Background()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	t.Run("Error from repository", func(t *testing.T) {
		repo.On("Export").Return(nil, errors.New("some error from repository")).Once()

		err := srv.Export(c, "pdf")

		assert.ErrorContains(t, err, "some error from repository")

		repo.AssertExpectations(t)
	})
	t.Run("Export PDF Succes Case", func(t *testing.T) {
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

		repo.On("Export").Return(data, nil).Once()

		repo.On("ExportFilePDF", c, data).Return(nil).Once()

		err := srv.Export(c, "pdf")

		assert.NoError(t, err)

		repo.AssertExpectations(t)
	})
	t.Run("Export Excel Succes Case", func(t *testing.T) {
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

		repo.On("Export").Return(data, nil).Once()

		repo.On("ExportFileExcel", c, data).Return(nil).Once()

		err := srv.Export(c, "xlsx")

		assert.NoError(t, err)

		repo.AssertExpectations(t)
	})
	t.Run("Export CSV Succes Case", func(t *testing.T) {
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

		repo.On("Export").Return(data, nil).Once()

		repo.On("ExportFileCsv", c, data).Return(nil).Once()

		err := srv.Export(c, "csv")

		assert.NoError(t, err)

		repo.AssertExpectations(t)
	})
	t.Run("Error export file Case", func(t *testing.T) {
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

		repo.On("Export").Return(data, nil).Once()

		err := srv.Export(c, "doc")

		assert.ErrorContains(t, err, "unsupported file type")
	})
}

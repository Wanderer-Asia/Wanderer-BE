package service

import (
	"context"
	"errors"
	"testing"
	"time"
	"wanderer/features/reports"
	"wanderer/features/reports/mocks"

	"github.com/stretchr/testify/assert"
)

func TestReportServiceDashboard(t *testing.T) {
	var repo = mocks.NewRepository(t)
	var srv = NewReportService(repo)
	var ctx = context.Background()

	var data = reports.Report{
		TotalUser:     1,
		TotalBooking:  1,
		TotalLocation: 1,
		TotalTour:     1,
		GraphBooking: []reports.GraphBooking{
			{
				Label: "test",
				Total: 10,
			},
		},
		RecentBooking: []reports.Booking{
			{
				Code:     1,
				Location: "test",
				Price:    10000,
			},
		},
		TopTours: []reports.Tour{
			{
				Id:        1,
				Title:     "test",
				Price:     10000,
				Discount:  0,
				Start:     time.Now(),
				Quota:     10,
				Rating:    5,
				Thumbnail: "test",
				Location: reports.Location{
					Name: "test",
				},
			},
		},
	}

	t.Run("error get total user from repository", func(t *testing.T) {
		repo.On("GetTotalUser", ctx).Return(0, errors.New("some error from repository")).Once()

		result, err := srv.Dashboard(ctx)

		assert.ErrorContains(t, err, "some error from repository")
		assert.Nil(t, result)

		repo.AssertExpectations(t)
	})

	t.Run("error get total booking from repository", func(t *testing.T) {
		repo.On("GetTotalUser", ctx).Return(data.TotalUser, nil).Once()
		repo.On("GetTotalBooking", ctx).Return(0, errors.New("some error from repository")).Once()

		result, err := srv.Dashboard(ctx)

		assert.ErrorContains(t, err, "some error from repository")
		assert.Nil(t, result)

		repo.AssertExpectations(t)
	})

	t.Run("error get total location from repository", func(t *testing.T) {
		repo.On("GetTotalUser", ctx).Return(data.TotalUser, nil).Once()
		repo.On("GetTotalBooking", ctx).Return(data.TotalBooking, nil).Once()
		repo.On("GetTotalLocation", ctx).Return(0, errors.New("some error from repository")).Once()

		result, err := srv.Dashboard(ctx)

		assert.ErrorContains(t, err, "some error from repository")
		assert.Nil(t, result)

		repo.AssertExpectations(t)
	})

	t.Run("error get total tour from repository", func(t *testing.T) {
		repo.On("GetTotalUser", ctx).Return(data.TotalUser, nil).Once()
		repo.On("GetTotalBooking", ctx).Return(data.TotalBooking, nil).Once()
		repo.On("GetTotalLocation", ctx).Return(data.TotalLocation, nil).Once()
		repo.On("GetTotalTour", ctx).Return(0, errors.New("some error from repository")).Once()

		result, err := srv.Dashboard(ctx)

		assert.ErrorContains(t, err, "some error from repository")
		assert.Nil(t, result)

		repo.AssertExpectations(t)
	})

	t.Run("error get graph booking from repository", func(t *testing.T) {
		repo.On("GetTotalUser", ctx).Return(data.TotalUser, nil).Once()
		repo.On("GetTotalBooking", ctx).Return(data.TotalBooking, nil).Once()
		repo.On("GetTotalLocation", ctx).Return(data.TotalLocation, nil).Once()
		repo.On("GetTotalTour", ctx).Return(data.TotalTour, nil).Once()
		repo.On("GetBookingCurrentYear", ctx).Return(nil, errors.New("some error from repository")).Once()

		result, err := srv.Dashboard(ctx)

		assert.ErrorContains(t, err, "some error from repository")
		assert.Nil(t, result)

		repo.AssertExpectations(t)
	})

	t.Run("error get recent booking from repository", func(t *testing.T) {
		repo.On("GetTotalUser", ctx).Return(data.TotalUser, nil).Once()
		repo.On("GetTotalBooking", ctx).Return(data.TotalBooking, nil).Once()
		repo.On("GetTotalLocation", ctx).Return(data.TotalLocation, nil).Once()
		repo.On("GetTotalTour", ctx).Return(data.TotalTour, nil).Once()
		repo.On("GetBookingCurrentYear", ctx).Return(data.GraphBooking, nil).Once()
		repo.On("GetRecentBooking", ctx).Return(nil, errors.New("some error from repository")).Once()

		result, err := srv.Dashboard(ctx)

		assert.ErrorContains(t, err, "some error from repository")
		assert.Nil(t, result)

		repo.AssertExpectations(t)
	})

	t.Run("error get top tour from repository", func(t *testing.T) {
		repo.On("GetTotalUser", ctx).Return(data.TotalUser, nil).Once()
		repo.On("GetTotalBooking", ctx).Return(data.TotalBooking, nil).Once()
		repo.On("GetTotalLocation", ctx).Return(data.TotalLocation, nil).Once()
		repo.On("GetTotalTour", ctx).Return(data.TotalTour, nil).Once()
		repo.On("GetBookingCurrentYear", ctx).Return(data.GraphBooking, nil).Once()
		repo.On("GetRecentBooking", ctx).Return(data.RecentBooking, nil).Once()
		repo.On("GetTopTour", ctx).Return(nil, errors.New("some error from repository")).Once()

		result, err := srv.Dashboard(ctx)

		assert.ErrorContains(t, err, "some error from repository")
		assert.Nil(t, result)

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		repo.On("GetTotalUser", ctx).Return(data.TotalUser, nil).Once()
		repo.On("GetTotalBooking", ctx).Return(data.TotalBooking, nil).Once()
		repo.On("GetTotalLocation", ctx).Return(data.TotalLocation, nil).Once()
		repo.On("GetTotalTour", ctx).Return(data.TotalTour, nil).Once()
		repo.On("GetBookingCurrentYear", ctx).Return(data.GraphBooking, nil).Once()
		repo.On("GetRecentBooking", ctx).Return(data.RecentBooking, nil).Once()
		repo.On("GetTopTour", ctx).Return(data.TopTours, nil).Once()

		result, err := srv.Dashboard(ctx)

		assert.NoError(t, err)
		assert.Equal(t, &data, result)

		repo.AssertExpectations(t)
	})
}

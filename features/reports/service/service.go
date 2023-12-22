package service

import (
	"context"
	"wanderer/features/reports"
)

func NewReportService(repo reports.Repository) reports.Service {
	return &reportService{
		repo: repo,
	}
}

type reportService struct {
	repo reports.Repository
}

func (srv *reportService) Dashboard(ctx context.Context) (*reports.Report, error) {
	totalUser, err := srv.repo.GetTotalUser(ctx)
	if err != nil {
		return nil, err
	}

	totalBooking, err := srv.repo.GetTotalBooking(ctx)
	if err != nil {
		return nil, err
	}

	totalLocation, err := srv.repo.GetTotalLocation(ctx)
	if err != nil {
		return nil, err
	}

	totalTour, err := srv.repo.GetTotalTour(ctx)
	if err != nil {
		return nil, err
	}

	graphBooking, err := srv.repo.GetBookingCurrentYear(ctx)
	if err != nil {
		return nil, err
	}

	recentBooking, err := srv.repo.GetRecentBooking(ctx)
	if err != nil {
		return nil, err
	}

	topTour, err := srv.repo.GetTopTour(ctx)
	if err != nil {
		return nil, err
	}

	return &reports.Report{
		TotalUser:     totalUser,
		TotalBooking:  totalBooking,
		TotalLocation: totalLocation,
		TotalTour:     totalTour,
		GraphBooking:  graphBooking,
		RecentBooking: recentBooking,
		TopTours:      topTour,
	}, nil
}

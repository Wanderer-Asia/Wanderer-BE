package repository

import (
	"context"
	"time"
	"wanderer/features/reports"

	"gorm.io/gorm"
)

func NewReportRepository(mysqlDB *gorm.DB) reports.Repository {
	return &reportRepository{
		mysqlDB: mysqlDB,
	}
}

type reportRepository struct {
	mysqlDB *gorm.DB
}

func (repo *reportRepository) GetTotalUser(ctx context.Context) (int, error) {
	var totalUser int64
	if err := repo.mysqlDB.WithContext(ctx).Model(&User{}).Where(&User{Role: "user"}).Count(&totalUser).Error; err != nil {
		return 0, err
	}
	return int(totalUser), nil
}

func (repo *reportRepository) GetTotalBooking(ctx context.Context) (int, error) {
	var totalBooking int64
	if err := repo.mysqlDB.WithContext(ctx).Model(&Booking{}).Where(&Booking{Status: "approved"}).Count(&totalBooking).Error; err != nil {
		return 0, err
	}
	return int(totalBooking), nil
}

func (repo *reportRepository) GetTotalLocation(ctx context.Context) (int, error) {
	var totalLocation int64
	if err := repo.mysqlDB.WithContext(ctx).Model(&Location{}).Count(&totalLocation).Error; err != nil {
		return 0, err
	}
	return int(totalLocation), nil
}

func (repo *reportRepository) GetTotalTour(ctx context.Context) (int, error) {
	var totalTour int64
	if err := repo.mysqlDB.WithContext(ctx).Model(&Tour{}).Count(&totalTour).Error; err != nil {
		return 0, err
	}
	return int(totalTour), nil
}

func (repo *reportRepository) GetBookingCurrentYear(ctx context.Context) ([]reports.GraphBooking, error) {
	rows, err := repo.mysqlDB.WithContext(ctx).Raw("SELECT yr.mnt, COALESCE(COUNT(code), 0) as total FROM (SELECT 1 AS mnt UNION SELECT 2 AS mnt UNION SELECT 3 AS mnt UNION SELECT 4 AS mnt UNION SELECT 5 AS mnt UNION SELECT 6 AS mnt UNION SELECT 7 AS mnt UNION SELECT 8 AS mnt UNION SELECT 9 AS mnt UNION SELECT 10 AS mnt UNION SELECT 11 AS mnt UNION SELECT 12 AS mnt) yr LEFT JOIN bookings ON DATE_FORMAT(bookings.booked_at, '%m') = yr.mnt AND DATE_FORMAT(bookings.booked_at, '%Y') = ? AND bookings.status = 'approved' GROUP BY yr.mnt", time.Now().Year()).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []reports.GraphBooking
	for rows.Next() {
		var result = new(GraphBooking)
		rows.Scan(
			&result.Label,
			&result.Total,
		)

		parse, err := result.ToEntity()
		if err != nil {
			return nil, err
		}

		data = append(data, *parse)
	}

	return data, nil
}

func (repo *reportRepository) GetRecentBooking(ctx context.Context) ([]reports.Booking, error) {
	var modBooking []Booking
	if err := repo.mysqlDB.WithContext(ctx).Model(&Booking{}).Joins("Tour.Location").Order("booked_at desc").Limit(5).Find(&modBooking).Error; err != nil {
		return nil, err
	}

	var data []reports.Booking
	for _, booking := range modBooking {
		data = append(data, *booking.ToEntity())
	}

	return data, nil
}

func (repo *reportRepository) GetTopTour(ctx context.Context) ([]reports.Tour, error) {
	var modTour []Tour
	if err := repo.mysqlDB.WithContext(ctx).Model(&Tour{}).Joins("Location").Order("ABS(quota - available) desc").Limit(5).Find(&modTour).Error; err != nil {
		return nil, err
	}

	var data []reports.Tour
	for _, booking := range modTour {
		data = append(data, *booking.ToEntity())
	}

	return data, nil
}

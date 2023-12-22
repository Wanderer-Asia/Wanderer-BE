package reports

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
)

type Report struct {
	TotalUser     int
	TotalBooking  int
	TotalLocation int
	TotalTour     int

	GraphBooking  []GraphBooking
	RecentBooking []Booking

	TopTours []Tour
}

type GraphBooking struct {
	Label string
	Total int
}

type Booking struct {
	Code     int
	Location string
	Price    float64
}

type Tour struct {
	Id       uint
	Title    string
	Price    float64
	Discount int
	Start    time.Time
	Quota    int
	Rating   float32

	Thumbnail string

	Location Location
}

type Location struct {
	Name string
}

type Handler interface {
	Dashboard() echo.HandlerFunc
}

type Service interface {
	Dashboard(ctx context.Context) (*Report, error)
}

type Repository interface {
	GetTotalUser(ctx context.Context) (int, error)
	GetTotalBooking(ctx context.Context) (int, error)
	GetTotalLocation(ctx context.Context) (int, error)
	GetTotalTour(ctx context.Context) (int, error)
	GetBookingCurrentYear(ctx context.Context) ([]GraphBooking, error)
	GetRecentBooking(ctx context.Context) ([]Booking, error)
	GetTopTour(ctx context.Context) ([]Tour, error)
}

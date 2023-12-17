package tours

import (
	"context"
	"io"
	"time"
	"wanderer/features/airlines"
	"wanderer/features/facilities"
	"wanderer/features/locations"
	"wanderer/helpers/filters"

	"github.com/labstack/echo/v4"
)

type Tour struct {
	Id          int
	Title       string
	Description string
	Price       float64
	AdminFee    float64
	Discount    int
	Start       time.Time
	Finish      time.Time
	Quota       int
	Available   int
	Rating      float32

	Thumbnail File
	Picture   []File

	Itinerary []Itinerary

	Facility []facilities.Facility

	Airline airlines.Airline

	Location locations.Location

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type File struct {
	Id int

	Raw io.Reader
	Url string

	CreatedAt time.Time
}

type Itinerary struct {
	Id          int
	Location    string
	Description string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Handler interface {
	GetAll() echo.HandlerFunc
	GetDetail() echo.HandlerFunc
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
}

type Service interface {
	GetAll(ctx context.Context, flt filters.Filter) ([]Tour, int, error)
	GetDetail(ctx context.Context) (*Tour, error)
	Create(ctx context.Context, data Tour) error
	Update(ctx context.Context, id uint, data Tour) error
}

type Repository interface {
	GetAll(ctx context.Context, flt filters.Filter) ([]Tour, int, error)
	GetDetail(ctx context.Context) (*Tour, error)
	Create(ctx context.Context, data Tour) error
	Update(ctx context.Context, id uint, data Tour) error
}

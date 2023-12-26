package tours

import (
	"context"
	"io"
	"time"
	"wanderer/helpers/filters"

	"github.com/labstack/echo/v4"
)

type Tour struct {
	Id          uint
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

	FacilityInclude []Facility
	FacilityExclude []Facility
	Airline         Airline
	Location        Location
	Reviews         []Review

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

type Facility struct {
	Id   uint
	Name string
}

type Airline struct {
	Id   uint
	Name string
}

type Location struct {
	Id   uint
	Name string
}

type Review struct {
	Id        uint
	Text      string
	Rating    float32
	CreatedAt time.Time
	User      User
	TourId    uint
}

type User struct {
	Id    uint
	Name  string
	Image string
}

type Handler interface {
	GetAll() echo.HandlerFunc
	GetDetail() echo.HandlerFunc
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
}

type Service interface {
	GetAll(ctx context.Context, flt filters.Filter) ([]Tour, int, error)
	GetDetail(ctx context.Context, id uint) (*Tour, error)
	Create(ctx context.Context, data Tour) error
	Update(ctx context.Context, id uint, data Tour) error
}

type Repository interface {
	GetAll(ctx context.Context, flt filters.Filter) ([]Tour, int, error)
	GetDetail(ctx context.Context, id uint) (*Tour, error)
	Create(ctx context.Context, data Tour) error
	Update(ctx context.Context, id uint, data Tour) error
}

package bookings

import (
	"context"
	"time"
	"wanderer/features/tours"
	"wanderer/features/users"
	"wanderer/helpers/filters"

	"github.com/labstack/echo/v4"
)

type Booking struct {
	Code      int
	Total     float64
	Status    string
	BookedAt  time.Time
	DeletedAt time.Time

	User users.User
	Tour tours.Tour

	Detail  []Detail
	Payment Payment
}

type Detail struct {
	Id             uint
	DocumentNumber string
	Greeting       string
	Name           string
	Nationality    string
	DOB            time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	BookingCode int
}

type Payment struct {
	Id            uint
	Method        string
	Bank          string
	VirtualNumber string
	BillKey       string
	BillCode      string
	Status        string

	BookingCode int

	CreatedAt time.Time
	ExpiredAt time.Time
	PaidAt    time.Time
}

type Handler interface {
	GetAll() echo.HandlerFunc
	GetDetail() echo.HandlerFunc
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	PaymentNotification() echo.HandlerFunc
}

type Service interface {
	GetAll(ctx context.Context, flt filters.Filter) ([]Booking, int, error)
	GetDetail(ctx context.Context, code int) (*Booking, error)
	Create(ctx context.Context, data Booking) (*Booking, error)
	Update(ctx context.Context, code int, data Booking) (*Booking, error)
}

type Repository interface {
	GetAll(ctx context.Context) ([]Booking, int, error)
	GetDetail(ctx context.Context, code int) (*Booking, error)
	Create(ctx context.Context, data Booking) (*Booking, error)
	Update(ctx context.Context, code int, data Booking) (*Booking, error)
}

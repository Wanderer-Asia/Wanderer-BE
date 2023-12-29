package bookings

import (
	"context"
	"time"
	"wanderer/helpers/filters"

	"github.com/labstack/echo/v4"
)

type Booking struct {
	Code      int
	Total     float64
	Status    string
	BookedAt  time.Time
	DeletedAt time.Time

	User User
	Tour Tour

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

	BookingCode  int
	BookingTotal float64

	CreatedAt time.Time
	ExpiredAt time.Time
	PaidAt    time.Time
}

type User struct {
	Id    uint
	Name  string
	Email string
	Phone string

	ImageUrl string
}

type Tour struct {
	Id          uint
	Title       string
	Description string
	Price       float64
	Start       time.Time
	Finish      time.Time
	Quota       int
	Available   int
	Rating      *float32

	Picture []File

	Itinerary []Itinerary

	FacilityInclude []Facility
	FacilityExclude []Facility

	Airline Airline

	Location Location

	Reviews []Review
}

type File struct {
	Id  int
	Url string
}

type Itinerary struct {
	Id          int
	Location    string
	Description string
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
	Id     uint
	TourId uint
	Text   string
	Rating float32

	User User

	CreatedAt time.Time
}

type Handler interface {
	GetAll() echo.HandlerFunc
	GetDetail() echo.HandlerFunc
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	PaymentNotification() echo.HandlerFunc
	ExportReportTransaction() echo.HandlerFunc
}

type Service interface {
	GetAll(ctx context.Context, flt filters.Filter) ([]Booking, int, error)
	GetDetail(ctx context.Context, code int) (*Booking, error)
	Create(ctx context.Context, data Booking) (*Booking, error)
	UpdateBookingStatus(ctx context.Context, code int, status string) error
	UpdatePaymentStatus(ctx context.Context, code int, paymentStatus string) error
	ChangePaymentMethod(ctx context.Context, code int, data Payment) (*Payment, error)
	Export(c echo.Context, typeFile string) error
}

type Repository interface {
	GetAll(ctx context.Context, flt filters.Filter) ([]Booking, int, error)
	GetDetail(ctx context.Context, code int) (*Booking, error)
	GetTourById(ctx context.Context, tourId uint) (*Tour, error)
	Create(ctx context.Context, data Booking) (*Booking, error)
	UpdateBookingStatus(ctx context.Context, code int, status string) error
	UpdatePaymentStatus(ctx context.Context, code int, bookingStatus string, paymentStatus string) error
	ChangePaymentMethod(ctx context.Context, code int, data Booking) (*Payment, error)
	Export() ([]Booking, error)
	ExportFileCsv(c echo.Context, data []Booking) error
	ExportFileExcel(c echo.Context, data []Booking) error
	ExportFilePDF(c echo.Context, data []Booking) error
}

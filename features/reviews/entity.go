package reviews

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Review struct {
	Id     uint
	TourId uint
	Text   string
	Rating float32

	User User

	CreatedAt time.Time
}

type User struct {
	Id    uint
	Name  string
	Image string
}

type Tour struct {
	Id     uint
	Finish time.Time
	Start  time.Time
}

type Booking struct {
	Code   int
	UserId uint
	TourId uint
	Status string
}

type Handler interface {
	Create() echo.HandlerFunc
}

type Repository interface {
	Create(userId uint, newReview Review) error
	GetTourById(tourId uint) (*Tour, error)
	IsBooking(tourId uint, userId uint) bool
	IsApproved(tourId uint, userId uint) bool
}

type Service interface {
	Create(userId uint, newReview Review) error
}

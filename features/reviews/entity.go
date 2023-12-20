package reviews

import (
	"time"
	"wanderer/features/users"

	"github.com/labstack/echo/v4"
)

type Review struct {
	Id     uint
	TourId uint
	Text   string
	Rating float32

	User users.User

	CreatedAt time.Time
}

type Handler interface {
	Create() echo.HandlerFunc
}

type Repository interface {
	Create(userId uint, newReview Review) error
	// GetByTourId(tourId uint) ([]Review, error)
}

type Service interface {
	Create(userId uint, newReview Review) error
	// GetAverage(tourId uint) (float32, error)
}

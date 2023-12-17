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

	CreatedAt time.Time
}

type Handler interface {
	Create() echo.HandlerFunc
}

type Repository interface {
	Create(userId uint, newReview Review) error
}

type Service interface {
	Create(userId uint, newReview Review) error
}

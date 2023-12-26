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

type Handler interface {
	Create() echo.HandlerFunc
}

type Repository interface {
	Create(userId uint, newReview Review) error
}

type Service interface {
	Create(userId uint, newReview Review) error
}

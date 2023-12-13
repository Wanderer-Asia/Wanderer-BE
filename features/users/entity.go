package users

import (
	"io"
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       uint
	Name     string
	Email    string
	Phone    string
	Password string
	Role     string

	ImageUrl string
	ImageRaw io.Reader

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Handler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type Service interface {
	Register() error
	Login() (*User, error)
	Update() error
	Delete() error
}

type Repository interface {
	Register() error
	Login() (*User, error)
	Update() error
	Delete() error
}

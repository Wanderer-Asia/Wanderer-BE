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

	TourCount   int
	ReviewCount int
	Bookings    []Booking

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Booking struct {
	Code        int
	DetailCount int
	Status      string
	Tour        Tour
}

type Tour struct {
	Id    uint
	Title string
}

type Handler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Detail() echo.HandlerFunc
}

type Service interface {
	Register(newUser User) error
	Login(email string, password string) (*User, error)
	Update(id uint, updateUser User) error
	Delete(id uint) error
	Detail(id uint) (*User, error)
}

type Repository interface {
	Register(newUser User) error
	Login(email string) (*User, error)
	Update(id uint, updateUser User) error
	Delete(id uint) error
	Detail(id uint) (*User, error)
}

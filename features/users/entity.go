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

	Bookings []Booking
	Reviews  []Review

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Booking struct {
	Code   int
	Status string
	UserId uint

	BookingDetails []BookingDetail
	DetailCount    int

	TourID uint
	Tour   Tour
}

type BookingDetail struct {
	Id          uint
	BookingCode int
}

type Tour struct {
	Id    uint
	Title string
}

type Review struct {
	Id     uint
	UserId uint
}

type Handler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetById() echo.HandlerFunc
}

type Service interface {
	Register(newUser User) error
	Login(email string, password string) (*User, error)
	Update(id uint, updateUser User) error
	Delete(id uint) error
	GetById(id uint) (*User, error)
}

type Repository interface {
	Register(newUser User) error
	Login(email string) (*User, error)
	Update(id uint, updateUser User) error
	Delete(id uint) error
	GetById(id uint) (*User, error)
}

package airlines

import (
	"context"
	"io"
	"time"
	"wanderer/helpers/filters"

	"github.com/labstack/echo/v4"
)

type Airline struct {
	Id   uint
	Name string

	ImageUrl string
	ImageRaw io.Reader

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Handler interface {
	Create() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	ImportTemplate() echo.HandlerFunc
	Import() echo.HandlerFunc
}

type Service interface {
	Create(newAirline Airline) error
	GetAll(flt filters.Filter) ([]Airline, error)
	Update(id uint, updateAirline Airline) error
	Delete(id uint) error
	Import(ctx context.Context, data []Airline) error
}

type Repository interface {
	Create(newAirline Airline) error
	GetAll(flt filters.Filter) ([]Airline, error)
	Update(id uint, updateAirline Airline) error
	Delete(id uint) error
	Import(ctx context.Context, data []Airline) error
}

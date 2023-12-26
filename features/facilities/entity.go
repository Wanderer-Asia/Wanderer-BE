package facilities

import (
	"context"
	"time"
	"wanderer/helpers/filters"

	"github.com/labstack/echo/v4"
)

type Facility struct {
	Id   uint
	Name string

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
	Create(newFacility Facility) error
	GetAll(flt filters.Filter) ([]Facility, error)
	Update(id uint, updateFacility Facility) error
	Delete(id uint) error
	Import(ctx context.Context, data []Facility) error
}

type Repository interface {
	Create(newFacility Facility) error
	GetAll(flt filters.Filter) ([]Facility, error)
	Update(id uint, updateFacility Facility) error
	Delete(id uint) error
	Import(ctx context.Context, data []Facility) error
}

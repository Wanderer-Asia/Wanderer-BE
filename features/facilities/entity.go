package facilities

import (
	"time"
	"wanderer/helpers/filters"

	"github.com/labstack/echo/v4"
)

type Facility struct {
	Id   uint
	Name string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Handler interface {
	Create() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type Service interface {
	Create(newFacility Facility) error
	GetAll(flt filters.Filter) ([]Facility, error)
	Update(id uint, updateFacility Facility) error
	Delete(id uint) error
}

type Repository interface {
	Create(newFacility Facility) error
	GetAll(flt filters.Filter) ([]Facility, error)
	Update(id uint, updateFacility Facility) error
	Delete(id uint) error
}

package locations

import (
	"context"
	"io"
	"time"
	"wanderer/helpers/filters"

	"github.com/labstack/echo/v4"
)

type Location struct {
	Id   uint
	Name string

	ImageUrl string
	ImageRaw io.Reader

	Tours []Tour

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Tour struct {
	Id       uint
	Title    string
	Discount int
	Start    time.Time
	Quota    int
	Rating   float32

	Thumbnail string
}

type Handler interface {
	GetAll() echo.HandlerFunc
	GetDetail() echo.HandlerFunc
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	ImportTemplate() echo.HandlerFunc
	Import() echo.HandlerFunc
}

type Service interface {
	GetAll(ctx context.Context, flt filters.Filter) ([]Location, error)
	GetDetail(ctx context.Context, id uint) (*Location, error)
	Create(ctx context.Context, data Location) error
	Update(ctx context.Context, id uint, data Location) error
	Delete(ctx context.Context, id uint) error
	Import(ctx context.Context, data []Location) error
}

type Repository interface {
	GetAll(ctx context.Context, flt filters.Filter) ([]Location, error)
	GetDetail(ctx context.Context, id uint) (*Location, error)
	Create(ctx context.Context, data Location) error
	Update(ctx context.Context, id uint, data Location) error
	Delete(ctx context.Context, id uint) error
	Import(ctx context.Context, data []Location) error
}

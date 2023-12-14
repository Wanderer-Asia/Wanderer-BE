package locations

import (
	"context"
	"io"
	"time"

	"github.com/labstack/echo/v4"
)

type Location struct {
	Id   uint
	Name string

	ImageUrl string
	ImageRaw io.Reader

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Handler interface {
	GetAll() echo.HandlerFunc
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type Service interface {
	GetAll(ctx context.Context) ([]Location, error)
	Create(ctx context.Context, data Location) error
	Update(ctx context.Context, id uint, data Location) error
	Delete(ctx context.Context, id uint) error
}

type Repository interface {
	GetAll(ctx context.Context) ([]Location, error)
	Create(ctx context.Context, data Location) error
	Update(ctx context.Context, id uint, data Location) error
	Delete(ctx context.Context, id uint) error
}

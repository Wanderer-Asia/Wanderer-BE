package locations

import (
	"context"

	"github.com/labstack/echo/v4"
)

type Location struct {
	Id   uint
	Name string
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

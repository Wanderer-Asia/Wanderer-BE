package handler

import (
	"wanderer/features/tours"

	echo "github.com/labstack/echo/v4"
)

func NewTourHandler(tourService tours.Service) tours.Handler {
	return &tourHandler{
		tourService: tourService,
	}
}

type tourHandler struct {
	tourService tours.Service
}

func (hdl *tourHandler) Create() echo.HandlerFunc {
	panic("unimplemented")
}

func (hdl *tourHandler) GetAll() echo.HandlerFunc {
	panic("unimplemented")
}

func (hdl *tourHandler) GetDetail() echo.HandlerFunc {
	panic("unimplemented")
}

func (hdl *tourHandler) Update() echo.HandlerFunc {
	panic("unimplemented")
}

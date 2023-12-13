package handler

import (
	"wanderer/features/locations"

	echo "github.com/labstack/echo/v4"
)

func NewLocationHandler(locationService locations.Service) locations.Handler {
	return &locationHandler{
		locationService: locationService,
	}
}

type locationHandler struct {
	locationService locations.Service
}

func (hdl *locationHandler) GetAll() echo.HandlerFunc {
	panic("unimplemented")
}

func (hdl *locationHandler) Create() echo.HandlerFunc {
	panic("unimplemented")
}

func (hdl *locationHandler) Update() echo.HandlerFunc {
	panic("unimplemented")
}

func (hdl *locationHandler) Delete() echo.HandlerFunc {
	panic("unimplemented")
}

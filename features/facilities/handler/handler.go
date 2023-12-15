package handler

import (
	"net/http"
	"strings"
	"wanderer/features/facilities"

	"github.com/labstack/echo/v4"
)

func NewFacilityHandler(facilityService facilities.Service) facilities.Handler {
	return &facilityHandler{
		facilityService: facilityService,
	}
}

type facilityHandler struct {
	facilityService facilities.Service
}

func (hdl *facilityHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)
		var request = new(CreateRequest)

		if err := c.Bind(request); err != nil {
			c.Logger().Error(err)

			response["message"] = "incorrect input data"
			return c.JSON(http.StatusBadRequest, response)
		}

		var data = request.ToEntity()

		if err := hdl.facilityService.Create(*data); err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "validate") {
				response["message"] = "this facility is already in the system"
				return c.JSON(http.StatusConflict, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		response["message"] = "create facility success"
		return c.JSON(http.StatusCreated, response)
	}
}

func (hdl *facilityHandler) GetAll() echo.HandlerFunc {
	panic("unimplemented")
}

func (hdl *facilityHandler) Update() echo.HandlerFunc {
	panic("unimplemented")
}

func (hdl *facilityHandler) Delete() echo.HandlerFunc {
	panic("unimplemented")
}

package handler

import (
	"net/http"
	"strconv"
	"strings"
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

func (hdl *tourHandler) GetAll() echo.HandlerFunc {
	panic("unimplemented")
}

func (hdl *tourHandler) GetDetail() echo.HandlerFunc {
	panic("unimplemented")
}

func (hdl *tourHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)
		var request = new(TourCreateUpdateRequest)

		if err := request.Bind(c); err != nil {
			c.Logger().Error(err)

			response["message"] = "bad request"
			return c.JSON(http.StatusBadRequest, response)
		}

		if err := hdl.tourService.Create(c.Request().Context(), request.ToEntity()); err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "validate: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "validate: ", "")
				return c.JSON(http.StatusBadRequest, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		response["message"] = "create tour success"
		return c.JSON(http.StatusCreated, response)
	}
}

func (hdl *tourHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)
		var request = new(TourCreateUpdateRequest)

		tourId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "invalid tour id"
			return c.JSON(http.StatusBadRequest, response)
		}

		if err := request.Bind(c); err != nil {
			c.Logger().Error(err)

			response["message"] = "bad request"
			return c.JSON(http.StatusBadRequest, response)
		}

		data := request.ToEntity()
		data.Id = uint(tourId)

		if err := hdl.tourService.Update(c.Request().Context(), uint(tourId), data); err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "validate: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "validate: ", "")
				return c.JSON(http.StatusBadRequest, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		response["message"] = "update tour success"
		return c.JSON(http.StatusOK, response)
	}
}

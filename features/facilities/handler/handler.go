package handler

import (
	"net/http"
	"strconv"
	"strings"
	"wanderer/features/facilities"
	"wanderer/helpers/filters"

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
	return func(c echo.Context) error {
		var response = make(map[string]any)
		var filter = new(filters.Filter)

		var search = new(filters.Search)
		c.Bind(search)
		filter.Search = *search

		result, err := hdl.facilityService.GetAll(*filter)
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		var data []GetAllResponse
		for _, v := range result {
			tmpFacility := new(GetAllResponse)
			tmpFacility.FromEntity(v)

			data = append(data, *tmpFacility)
		}

		response["message"] = "get all facility success"
		response["data"] = data
		return c.JSON(http.StatusOK, response)
	}
}

func (hdl *facilityHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)
		var request = new(CreateRequest)

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "invalid facility id"
		}

		if c.Bind(request); err != nil {
			c.Logger().Error(err)

			response["message"] = "incorrect input"
			return c.JSON(http.StatusBadRequest, response)
		}

		if err := hdl.facilityService.Update(uint(id), *request.ToEntity()); err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "validate: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "validate: ", "")
				return c.JSON(http.StatusBadRequest, response)
			}

			if strings.Contains(err.Error(), "not found: ") {
				response["message"] = "facility not found"
				return c.JSON(http.StatusNotFound, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		response["message"] = "update facility success"
		return c.JSON(http.StatusOK, response)
	}
}

func (hdl *facilityHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "invalid facility id"
		}

		if err := hdl.facilityService.Delete(uint(id)); err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "not found: ") {
				response["message"] = "facility not found"
				return c.JSON(http.StatusNotFound, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		response["message"] = "delete facility success"
		return c.JSON(http.StatusOK, response)

	}

}

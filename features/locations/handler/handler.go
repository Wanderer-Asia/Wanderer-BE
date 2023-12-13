package handler

import (
	"net/http"
	"strconv"
	"strings"
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
	return func(c echo.Context) error {
		var response = make(map[string]any)

		result, err := hdl.locationService.GetAll(c.Request().Context())
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		var data []LocationResponse
		for _, res := range result {
			tmpLoc := new(LocationResponse)
			tmpLoc.FromEntity(res)

			data = append(data, *tmpLoc)
		}

		response["message"] = "get all location success"
		response["data"] = data
		return c.JSON(http.StatusOK, response)
	}
}

func (hdl *locationHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)
		var request = new(LocationCreateUpdateRequest)

		if err := c.Bind(request); err != nil {
			c.Logger().Error(err)

			response["message"] = "please fill input correctly"
			return c.JSON(http.StatusBadRequest, response)
		}

		err := hdl.locationService.Create(c.Request().Context(), request.ToEntity())
		if err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "validate: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "validate: ", "")
				return c.JSON(http.StatusBadRequest, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		response["message"] = "create location success"
		return c.JSON(http.StatusCreated, response)
	}
}

func (hdl *locationHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)
		var request = new(LocationCreateUpdateRequest)

		locationId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "invalid location id"
			return c.JSON(http.StatusBadRequest, response)
		}

		if err := c.Bind(request); err != nil {
			c.Logger().Error(err)

			response["message"] = "please fill input correctly"
			return c.JSON(http.StatusBadRequest, response)
		}

		if err := hdl.locationService.Update(c.Request().Context(), uint(locationId), request.ToEntity()); err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "validate: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "validate: ", "")
				return c.JSON(http.StatusBadRequest, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		response["message"] = "update location success"
		return c.JSON(http.StatusOK, response)
	}
}

func (hdl *locationHandler) Delete() echo.HandlerFunc {
	panic("unimplemented")
}

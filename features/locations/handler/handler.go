package handler

import (
	"net/http"
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
	panic("unimplemented")
}

func (hdl *locationHandler) Update() echo.HandlerFunc {
	panic("unimplemented")
}

func (hdl *locationHandler) Delete() echo.HandlerFunc {
	panic("unimplemented")
}

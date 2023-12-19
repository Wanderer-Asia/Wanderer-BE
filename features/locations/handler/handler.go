package handler

import (
	"net/http"
	"strconv"
	"strings"
	"wanderer/features/locations"
	"wanderer/features/tours"
	"wanderer/helpers/filters"

	echo "github.com/labstack/echo/v4"
)

func NewLocationHandler(locationService locations.Service, tourService tours.Service) locations.Handler {
	return &locationHandler{
		locationService: locationService,
		tourService:     tourService,
	}
}

type locationHandler struct {
	locationService locations.Service
	tourService     tours.Service
}

func (hdl *locationHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)
		var filter = new(filters.Filter)

		var pagination = new(filters.Pagination)
		c.Bind(pagination)
		filter.Pagination = *pagination

		var search = new(filters.Search)
		c.Bind(search)
		filter.Search = *search

		result, err := hdl.locationService.GetAll(c.Request().Context(), *filter)
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

		file, _ := c.FormFile("image")
		if file != nil {
			src, err := file.Open()
			if err != nil {
				return err
			}
			defer src.Close()

			request.ImageRaw = src
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

		file, _ := c.FormFile("image")
		if file != nil {
			src, err := file.Open()
			if err != nil {
				return err
			}
			defer src.Close()

			request.ImageRaw = src
		}

		if err := hdl.locationService.Update(c.Request().Context(), uint(locationId), request.ToEntity()); err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "validate: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "validate: ", "")
				return c.JSON(http.StatusBadRequest, response)
			}

			if strings.Contains(err.Error(), "not found") {
				response["message"] = "location not found"
				return c.JSON(http.StatusNotFound, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		response["message"] = "update location success"
		return c.JSON(http.StatusOK, response)
	}
}

func (hdl *locationHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)

		locationId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "invalid location id"
			return c.JSON(http.StatusBadRequest, response)
		}

		if err := hdl.locationService.Delete(c.Request().Context(), uint(locationId)); err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "validate: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "validate: ", "")
				return c.JSON(http.StatusBadRequest, response)
			}

			if strings.Contains(err.Error(), "not found") {
				response["message"] = "location not found"
				return c.JSON(http.StatusNotFound, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		response["message"] = "delete location success"
		return c.JSON(http.StatusOK, response)
	}
}

func (hdl *locationHandler) GetDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)

		locationId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "invalid location id"
			return c.JSON(http.StatusBadRequest, response)
		}

		location, err := hdl.locationService.GetDetail(c.Request().Context(), uint(locationId))
		if err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "validate: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "validate: ", "")
				return c.JSON(http.StatusBadRequest, response)
			}

			if strings.Contains(err.Error(), "not found") {
				response["message"] = "location not found"
				return c.JSON(http.StatusNotFound, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		tours, err := hdl.tourService.GetByLocation(c.Request().Context(), uint(locationId))
		if err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "validate: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "validate: ", "")
				return c.JSON(http.StatusBadRequest, response)
			}

			if strings.Contains(err.Error(), "not found") {
				response["message"] = "location not found"
				return c.JSON(http.StatusNotFound, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		var data = new(DetailLocationResponse)
		if location != nil || tours != nil {
			data.FromEntity(*location, tours)
		}

		response["message"] = "get detail location success"
		response["data"] = data
		return c.JSON(http.StatusOK, response)
	}
}

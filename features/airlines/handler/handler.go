package handler

import (
	"net/http"
	"strconv"
	"strings"
	"wanderer/features/airlines"
	"wanderer/helpers/filters"

	echo "github.com/labstack/echo/v4"
)

func NewAirlineHandler(airlineService airlines.Service) airlines.Handler {
	return &airlineHandler{
		airlineService: airlineService,
	}
}

type airlineHandler struct {
	airlineService airlines.Service
}

func (hdl *airlineHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)
		var request = new(CreateRequest)

		if err := c.Bind(request); err != nil {
			c.Logger().Error(err)

			response["message"] = "bad request"
			return c.JSON(http.StatusBadRequest, response)
		}

		file, _ := c.FormFile("logo")
		if file != nil {
			src, err := file.Open()
			if err != nil {
				return err
			}
			defer src.Close()

			request.Image = src
		}

		var data = request.ToEntity()

		if err := hdl.airlineService.Create(*data); err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "validate: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "validate: ", "")
				return c.JSON(http.StatusBadRequest, response)
			}

			if strings.Contains(err.Error(), "used: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "used: ", "")
				return c.JSON(http.StatusConflict, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		response["message"] = "create airline success"
		return c.JSON(http.StatusCreated, response)
	}
}

func (hdl *airlineHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)
		var filter = new(filters.Filter)

		var search = new(filters.Search)
		c.Bind(search)
		filter.Search = *search

		result, err := hdl.airlineService.GetAll(*filter)
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		var data []GetAllResponse
		for _, v := range result {
			tmpAir := new(GetAllResponse)
			tmpAir.FromEntity(v)

			data = append(data, *tmpAir)
		}

		response["message"] = "get all airlines success"
		response["data"] = data
		return c.JSON(http.StatusOK, response)
	}
}

func (hdl *airlineHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)
		var request = new(CreateRequest)

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "invalid airline id"
		}

		if c.Bind(request); err != nil {
			c.Logger().Error(err)

			response["message"] = "bad request"
			return c.JSON(http.StatusBadRequest, response)
		}

		file, _ := c.FormFile("logo")
		if file != nil {
			src, err := file.Open()
			if err != nil {
				return err
			}
			defer src.Close()

			request.Image = src
		}

		if err := hdl.airlineService.Update(uint(id), *request.ToEntity()); err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "validate: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "validate: ", "")
				return c.JSON(http.StatusBadRequest, response)
			}

			if strings.Contains(err.Error(), "used: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "used: ", "")
				return c.JSON(http.StatusConflict, response)
			}

			if strings.Contains(err.Error(), "not found: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "not found: ", "")
				return c.JSON(http.StatusNotFound, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		response["message"] = "update airline success"
		return c.JSON(http.StatusOK, response)
	}
}

func (hdl *airlineHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "invalid airline id"
		}

		if err := hdl.airlineService.Delete(uint(id)); err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "not found: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "not found: ", "")
				return c.JSON(http.StatusNotFound, response)
			}

			if strings.Contains(err.Error(), "used: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "used: ", "")
				return c.JSON(http.StatusConflict, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		response["message"] = "delete airline success"
		return c.JSON(http.StatusOK, response)
	}
}

func (hdl *airlineHandler) ImportTemplate() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Attachment("./helpers/imports/templates/airline.csv", "airline_import.csv")
	}
}

func (hdl *airlineHandler) Import() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)
		var request = new(ImportAirlineRequest)

		if err := request.Bind(c); err != nil {
			c.Logger().Error(err)

			response["message"] = "bad request"
			return c.JSON(http.StatusBadRequest, response)
		}

		data, err := request.ToEntity()
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "bad request"
			return c.JSON(http.StatusBadRequest, response)
		}

		if err := hdl.airlineService.Import(c.Request().Context(), data); err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "validate: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "validate: ", "")
				return c.JSON(http.StatusBadRequest, response)
			}

			if strings.Contains(err.Error(), "used: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "used: ", "")
				return c.JSON(http.StatusConflict, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		response["message"] = "import airline success"
		return c.JSON(http.StatusCreated, response)
	}
}

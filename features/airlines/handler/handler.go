package handler

import (
	"net/http"
	"strings"
	"wanderer/features/airlines"

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

			response["message"] = "incorrect input data"
			return c.JSON(http.StatusBadRequest, response)
		}

		file, _ := c.FormFile("image")
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

			if strings.Contains(err.Error(), "validate") {
				response["message"] = strings.ReplaceAll(err.Error(), "validate: ", "")
				return c.JSON(http.StatusBadRequest, response)
			}

			if strings.Contains(err.Error(), "Duplicate") {
				response["message"] = "airline name is already in the system"
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

		result, err := hdl.airlineService.GetAll()
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
	panic("unimplemented")
}

func (hdl *airlineHandler) Delete() echo.HandlerFunc {
	panic("unimplemented")
}

package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"wanderer/config"
	"wanderer/features/tours"
	"wanderer/helpers/filters"

	echo "github.com/labstack/echo/v4"
)

func NewTourHandler(tourService tours.Service, jwtConfig config.JWT) tours.Handler {
	return &tourHandler{
		tourService: tourService,
		jwtConfig:   jwtConfig,
	}
}

type tourHandler struct {
	tourService tours.Service
	jwtConfig   config.JWT
}

func (hdl *tourHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)
		var baseUrl = c.Scheme() + "://" + c.Request().Host

		var pagination = new(filters.Pagination)
		c.Bind(pagination)
		if pagination.Start != 0 && pagination.Limit == 0 {
			pagination.Limit = 5
		}

		var search = new(filters.Search)
		c.Bind(search)

		var sort = new(filters.Sort)
		c.Bind(sort)

		result, totalData, err := hdl.tourService.GetAll(context.Background(), filters.Filter{Search: *search, Pagination: *pagination, Sort: *sort})
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		var data []TourResponse
		for _, tour := range result {
			var tmpTour = new(TourResponse)
			tmpTour.FromEntity(tour, false)

			data = append(data, *tmpTour)
		}
		response["data"] = data

		if pagination.Limit != 0 {
			var paginationResponse = make(map[string]any)
			if pagination.Start >= (pagination.Limit) {
				prev := fmt.Sprintf("%s%s?start=%d&limit=%d", baseUrl, c.Path(), pagination.Start-pagination.Limit, pagination.Limit)

				if search.Keyword != "" {
					prev += "&keyword=" + search.Keyword
				}

				if sort.Column != "" {
					prev += "&sort=" + sort.Column
				}

				if sort.Direction {
					prev += "&dir=true"
				} else {
					prev += "&dir=false"
				}

				paginationResponse["prev"] = prev
			} else {
				paginationResponse["prev"] = nil
			}

			if totalData > pagination.Start+pagination.Limit {
				next := fmt.Sprintf("%s%s?start=%d&limit=%d", baseUrl, c.Path(), pagination.Start+pagination.Limit, pagination.Limit)

				if search.Keyword != "" {
					next += "&keyword=" + search.Keyword
				}

				if sort.Column != "" {
					next += "&sort=" + sort.Column
				}

				if sort.Direction {
					next += "&dir=true"
				} else {
					next += "&dir=false"
				}

				paginationResponse["next"] = next
			} else {
				paginationResponse["next"] = nil
			}
			response["pagination"] = paginationResponse
		}

		response["message"] = "get all tour success"
		return c.JSON(http.StatusOK, response)
	}
}

func (hdl *tourHandler) GetDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)

		tourId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "invalid tour id"
			return c.JSON(http.StatusBadRequest, response)
		}

		result, err := hdl.tourService.GetDetail(c.Request().Context(), uint(tourId))
		if err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "validate: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "validate: ", "")
				return c.JSON(http.StatusBadRequest, response)
			}

			if strings.Contains(err.Error(), "not found: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "not found: ", "")
				return c.JSON(http.StatusNotFound, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		var data = new(TourResponse)
		if result != nil {
			data.FromEntity(*result, true)
		}

		response["message"] = "get detail tour success"
		response["data"] = data
		return c.JSON(http.StatusOK, response)
	}
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

			if strings.Contains(err.Error(), "not found: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "not found: ", "")
				return c.JSON(http.StatusNotFound, response)
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

			if strings.Contains(err.Error(), "not found: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "not found: ", "")
				return c.JSON(http.StatusNotFound, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		response["message"] = "update tour success"
		return c.JSON(http.StatusOK, response)
	}
}

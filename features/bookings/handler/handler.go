package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"wanderer/config"
	"wanderer/features/bookings"
	"wanderer/helpers/filters"
	"wanderer/helpers/tokens"

	"github.com/golang-jwt/jwt/v5"
	echo "github.com/labstack/echo/v4"
)

func NewBookingHandler(bookingService bookings.Service, jwtConfig config.JWT) bookings.Handler {
	return &bookingHandler{
		bookingService: bookingService,
		jwtConfig:      jwtConfig,
	}
}

type bookingHandler struct {
	bookingService bookings.Service
	jwtConfig      config.JWT
}

func (hdl *bookingHandler) GetAll() echo.HandlerFunc {
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

		result, totalData, err := hdl.bookingService.GetAll(context.Background(), filters.Filter{Pagination: *pagination, Sort: *sort})
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		var data []BookingResponse
		for _, booking := range result {
			var tmpBooking = new(BookingResponse)
			tmpBooking.FromEntity(booking)

			tmpBooking.User.Image = ""

			data = append(data, *tmpBooking)
		}
		response["data"] = data

		if pagination.Limit != 0 {
			var paginationResponse = make(map[string]any)
			if pagination.Start >= (pagination.Limit) {
				prev := fmt.Sprintf("%s%s?start=%d&limit=%d", baseUrl, c.Path(), pagination.Start-pagination.Limit, pagination.Limit)
				if search.Keyword != "" {
					prev += "&keyword=" + search.Keyword
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

func (hdl *bookingHandler) GetDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)

		bookingCode, err := strconv.Atoi(c.Param("code"))
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "invalid booking code"
			return c.JSON(http.StatusBadRequest, response)
		}

		result, err := hdl.bookingService.GetDetail(c.Request().Context(), bookingCode)
		if err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "not found: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "not found: ", "")
				return c.JSON(http.StatusNotFound, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		if result != nil {
			var data = new(BookingResponse)
			data.FromEntity(*result)

			response["data"] = data
		}

		response["message"] = "get detail booking success"
		return c.JSON(http.StatusOK, response)
	}
}

func (hdl *bookingHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)
		var request = new(BookingCreateUpdateRequest)

		token := c.Get("user")
		if token == nil {
			response["message"] = "unauthorized access"
			return c.JSON(http.StatusUnauthorized, response)
		}

		userId, err := tokens.ExtractToken(hdl.jwtConfig.Secret, token.(*jwt.Token))
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "unauthorized"
			return c.JSON(http.StatusUnauthorized, response)
		}

		if err := c.Bind(request); err != nil {
			c.Logger().Error(err)

			response["message"] = "bad request"
			return c.JSON(http.StatusBadRequest, response)
		}

		result, err := hdl.bookingService.Create(c.Request().Context(), request.ToEntity(userId))
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

		var data = new(BookingResponse)
		data.FromEntity(*result)

		response["message"] = "create booking success"
		response["data"] = data
		return c.JSON(http.StatusCreated, response)
	}
}

func (hdl *bookingHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)
		var request = new(BookingCreateUpdateRequest)

		bookingCode, err := strconv.Atoi(c.Param("code"))
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "invalid booking code"
			return c.JSON(http.StatusBadRequest, response)
		}

		if err := c.Bind(request); err != nil {
			c.Logger().Error(err)

			response["message"] = "bad request"
			return c.JSON(http.StatusBadRequest, response)
		}

		result, err := hdl.bookingService.Update(c.Request().Context(), bookingCode, request.ToEntity(0))
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

		if request.Status == "refund" {
			response["message"] = "refund success"
		} else if request.Status == "refunded" {
			response["message"] = "approve refund success"
		} else if request.Bank != "" {
			var data = new(BookingResponse)
			data.FromEntity(*result)

			response["message"] = "change payment method success"
			response["data"] = data
		} else {
			response["message"] = "update booking success"
		}

		return c.JSON(http.StatusOK, response)
	}
}

func (hdl *bookingHandler) PaymentNotification() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request = new(PaymentNotificationRequest)

		if err := c.Bind(request); err != nil {
			c.Logger().Error(err)

			return c.JSON(http.StatusBadRequest, "bad request")
		}

		var data = new(bookings.Booking)

		switch request.Status {
		case "settlement":
			data.Status = "approved"
			data.Payment.Status = request.Status
		case "cancel", "expire":
			data.Status = "cancel"
			data.Payment.Status = request.Status
		case "capture", "deny", "pending":
			data.Status = "pending"
			data.Payment.Status = request.Status
		default:
			c.Logger().Error(errors.New("invalid payment status"))

			return c.JSON(http.StatusBadRequest, "invalid payment status")
		}

		code, err := strconv.Atoi(request.Code)
		if err != nil {
			c.Logger().Error(err)

			return c.JSON(http.StatusBadRequest, "bad request")
		}

		_, err = hdl.bookingService.Update(c.Request().Context(), code, *data)
		if err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "validate: ") {
				return c.JSON(http.StatusBadRequest, strings.ReplaceAll(err.Error(), "validate: ", ""))
			}

			return c.JSON(http.StatusInternalServerError, "internal server error")
		}

		return c.JSON(http.StatusOK, "ok")
	}
}

package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"wanderer/config"
	"wanderer/features/bookings"
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
	panic("unimplemented")
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

			if strings.Contains(err.Error(), "not found") {
				response["message"] = "booking not found"
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

			response["message"] = "please fill input correctly"
			return c.JSON(http.StatusBadRequest, response)
		}

		result, err := hdl.bookingService.Create(c.Request().Context(), request.ToEntity(userId))
		if err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "validate: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "validate: ", "")
				return c.JSON(http.StatusBadRequest, response)
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

			response["message"] = "please fill input correctly"
			return c.JSON(http.StatusBadRequest, response)
		}

		result, err := hdl.bookingService.Update(c.Request().Context(), bookingCode, request.ToEntity(0))
		if err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "validate: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "validate: ", "")
				return c.JSON(http.StatusBadRequest, response)
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

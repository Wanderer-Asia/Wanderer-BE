package handler

import (
	"net/http"
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
	panic("unimplemented")
}

func (hdl *bookingHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)
		var request = new(BookingCreateRequest)

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
	panic("unimplemented")
}

package handler

import (
	"net/http"
	"strings"
	"wanderer/config"
	"wanderer/features/reviews"
	"wanderer/helpers/tokens"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func NewReviewHandler(reviewService reviews.Service, jwtConfig config.JWT) reviews.Handler {
	return &reviewHandler{
		reviewService: reviewService,
		jwtConfig:     jwtConfig,
	}
}

type reviewHandler struct {
	reviewService reviews.Service
	jwtConfig     config.JWT
}

func (hdl *reviewHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)
		var request = new(CreateRequest)

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

		var data = request.ToEntity()

		if err := hdl.reviewService.Create(userId, *data); err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "validate: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "validate: ", "")
				return c.JSON(http.StatusBadRequest, response)
			}

			if strings.Contains(err.Error(), "started") {
				response["message"] = strings.ReplaceAll(err.Error(), "cannot create review: ", "")
				return c.JSON(http.StatusBadRequest, response)
			}

			if strings.Contains(err.Error(), "finished") {
				response["message"] = strings.ReplaceAll(err.Error(), "cannot create review: ", "")
				return c.JSON(http.StatusBadRequest, response)
			}

			if strings.Contains(err.Error(), "booked") {
				response["message"] = strings.ReplaceAll(err.Error(), "cannot create review: ", "")
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

		response["message"] = "create review success"
		return c.JSON(http.StatusOK, response)
	}
}

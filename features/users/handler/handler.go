package handler

import (
	"net/http"
	"strings"
	"wanderer/config"
	"wanderer/features/users"
	"wanderer/helpers/tokens"

	"github.com/golang-jwt/jwt/v5"
	echo "github.com/labstack/echo/v4"
)

func NewUserHandler(userService users.Service, jwtConfig config.JWT) users.Handler {
	return &userHandler{
		userService: userService,
		jwtConfig:   jwtConfig,
	}
}

type userHandler struct {
	userService users.Service
	jwtConfig   config.JWT
}

func (hdl *userHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)
		var request = new(RegisterRequest)

		if err := c.Bind(request); err != nil {
			c.Logger().Error(err)

			response["message"] = "incorrect input data"
			return c.JSON(http.StatusBadRequest, response)
		}

		var data = request.ToEntity()

		if err := hdl.userService.Register(*data); err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "validate") {
				response["message"] = strings.ReplaceAll(err.Error(), "validate: ", "")
				return c.JSON(http.StatusBadRequest, response)
			}

			if strings.Contains(err.Error(), "Duplicate") {
				response["message"] = "email is already in use"
				return c.JSON(http.StatusConflict, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		response["message"] = "register success"
		return c.JSON(http.StatusCreated, response)
	}
}

func (hdl *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)
		var request = new(LoginRequest)

		if err := c.Bind(request); err != nil {
			c.Logger().Error(err)

			response["message"] = "please fill input correctly"
			return c.JSON(http.StatusBadRequest, response)
		}

		var input = request.ToEntity()

		result, err := hdl.userService.Login(input.Email, input.Password)
		if err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "validate") {
				response["message"] = strings.ReplaceAll(err.Error(), "validate: ", "")
				return c.JSON(http.StatusBadRequest, response)
			}

			if strings.Contains(err.Error(), "wrong password") {
				response["message"] = err.Error()
				return c.JSON(http.StatusBadRequest, response)
			}

			if strings.Contains(err.Error(), "not found") {
				response["message"] = "user not found"
				return c.JSON(http.StatusNotFound, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		strToken, err := tokens.GenerateJWT(hdl.jwtConfig.Secret, result.Id)
		if err != nil {
			return err
		}

		var data = new(LoginResponse)
		data.FromEntity(*result)
		data.Token = strToken

		response["message"] = "login success"
		response["data"] = data
		return c.JSON(http.StatusOK, response)
	}
}

func (hdl *userHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)
		var request = new(UpdateRequest)

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

		if c.Bind(request); err != nil {
			c.Logger().Error(err)

			response["message"] = "incorect input"
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

		if err := hdl.userService.Update(userId, *request.ToEntity()); err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "validate: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "validate: ", "")
				return c.JSON(http.StatusBadRequest, response)
			}

			if strings.Contains(err.Error(), "not found: ") {
				response["message"] = "user not found"
				return c.JSON(http.StatusNotFound, response)
			}

			if strings.Contains(err.Error(), "Duplicate") {
				response["message"] = "this email has been used, please use another email"
				return c.JSON(http.StatusConflict, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		response["message"] = "update user success"
		return c.JSON(http.StatusOK, response)
	}
}

func (hdl *userHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]interface{})

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

		if err := hdl.userService.Delete(userId); err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "not found: ") {
				response["message"] = "user not found"
				return c.JSON(http.StatusNotFound, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		response["message"] = "delete user success"
		return c.JSON(http.StatusOK, response)
	}
}

func (hdl *userHandler) Detail() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]interface{})

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

		result, err := hdl.userService.Detail(userId)
		if err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "not found: ") {
				response["message"] = "user not found"
				return c.JSON(http.StatusNotFound, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		if result != nil {
			var data = new(UserResponse)
			data.FromEntity(*result)

			response["data"] = data
		}

		response["message"] = "get profile user success"
		return c.JSON(http.StatusOK, response)
	}
}

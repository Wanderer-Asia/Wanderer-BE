package handler

import (
	"wanderer/features/users"

	echo "github.com/labstack/echo/v4"
)

func NewUserHandler(userService users.Service) users.Handler {
	return &userHandler{
		userService: userService,
	}
}

type userHandler struct {
	userService users.Service
}

func (hdl *userHandler) Register() echo.HandlerFunc {
	panic("unimplemented")
}

func (hdl *userHandler) Login() echo.HandlerFunc {
	panic("unimplemented")
}

func (hdl *userHandler) Update() echo.HandlerFunc {
	panic("unimplemented")
}

func (hdl *userHandler) Delete() echo.HandlerFunc {
	panic("unimplemented")
}

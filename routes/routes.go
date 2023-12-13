package routes

import (
	"wanderer/features/users"

	"github.com/labstack/echo/v4"
)

type Routes struct {
	Server      *echo.Echo
	UserHandler users.Handler
}

func (router Routes) InitRouter() {
}

func (router *Routes) UserRouter() {
}

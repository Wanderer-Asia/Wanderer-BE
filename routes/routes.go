package routes

import (
	"github.com/labstack/echo/v4"
)

type Routes struct {
	Server *echo.Echo
}

func (router Routes) InitRouter() {
}

func (router *Routes) UserRouter() {
}

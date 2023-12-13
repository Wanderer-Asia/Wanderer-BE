package routes

import (
	"wanderer/features/users"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type Routes struct {
	Server *echo.Echo

	UserHandler users.Handler
}

func (router Routes) InitRouter() {
	router.UserRouter()
}

func (router *Routes) UserRouter() {
	router.Server.POST("/register", router.UserHandler.Register())
	router.Server.POST("/login", router.UserHandler.Login())
	router.Server.PATCH("/users", router.UserHandler.Update(), echojwt.JWT([]byte("altamantul")))
}

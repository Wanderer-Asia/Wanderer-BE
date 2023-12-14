package routes

import (
	"wanderer/features/airlines"
	"wanderer/features/users"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type Routes struct {
	Server *echo.Echo

	UserHandler    users.Handler
	AirlineHandler airlines.Handler
}

func (router Routes) InitRouter() {
	router.UserRouter()
	router.AirlineRouter()
}

func (router *Routes) UserRouter() {
	router.Server.POST("/register", router.UserHandler.Register())
	router.Server.POST("/login", router.UserHandler.Login())
	router.Server.PATCH("/users", router.UserHandler.Update(), echojwt.JWT([]byte("altamantul")))
	router.Server.DELETE("/users", router.UserHandler.Delete(), echojwt.JWT([]byte("altamantul")))
}

func (router *Routes) AirlineRouter() {
	router.Server.POST("/airlines", router.AirlineHandler.Create(), echojwt.JWT([]byte("altamantul")))
}

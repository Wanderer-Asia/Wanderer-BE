package routes

import (
	"wanderer/features/airlines"
	"wanderer/features/users"
	"wanderer/features/locations"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type Routes struct {
	JWTKey          string
	Server          *echo.Echo
	UserHandler    users.Handler
	AirlineHandler airlines.Handler
	LocationHandler locations.Handler
}

func (router Routes) InitRouter() {
	router.UserRouter()
	router.AirlineRouter()
	router.LocationRouter()
}

func (router *Routes) UserRouter() {
	router.Server.POST("/register", router.UserHandler.Register())
	router.Server.POST("/login", router.UserHandler.Login())
	router.Server.PATCH("/users", router.UserHandler.Update(), echojwt.JWT([]byte(router.JWTKey)))
	router.Server.DELETE("/users", router.UserHandler.Delete(), echojwt.JWT([]byte(router.JWTKey)))
}

func (router *Routes) AirlineRouter() {
	router.Server.POST("/airlines", router.AirlineHandler.Create(), echojwt.JWT([]byte(router.JWTKey)))
	router.Server.GET("/airlines", router.AirlineHandler.GetAll())
	router.Server.PUT("/airlines/:id", router.AirlineHandler.Update(), echojwt.JWT([]byte(router.JWTKey)))
	router.Server.DELETE("airlines/:id", router.AirlineHandler.Delete(), echojwt.JWT([]byte(router.JWTKey)))
}

func (router *Routes) LocationRouter() {
	router.Server.GET("/locations", router.LocationHandler.GetAll())
	router.Server.POST("/locations", router.LocationHandler.Create(), echojwt.JWT([]byte(router.JWTKey)))
	router.Server.PUT("/locations/:id", router.LocationHandler.Update(), echojwt.JWT([]byte(router.JWTKey)))
	router.Server.DELETE("/locations/:id", router.LocationHandler.Delete(), echojwt.JWT([]byte(router.JWTKey)))
}
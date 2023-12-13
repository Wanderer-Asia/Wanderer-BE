package routes

import (
	"wanderer/features/locations"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type Routes struct {
	JWTKey          string
	Server          *echo.Echo
	LocationHandler locations.Handler
}

func (router Routes) InitRouter() {
	router.LocationRouter()
}

func (router *Routes) LocationRouter() {
	router.Server.GET("/locations", router.LocationHandler.GetAll())
	router.Server.POST("/locations", router.LocationHandler.Create(), echojwt.JWT([]byte(router.JWTKey)))
	router.Server.PUT("/locations/:id", router.LocationHandler.Update(), echojwt.JWT([]byte(router.JWTKey)))
	router.Server.DELETE("/locations/:id", router.LocationHandler.Delete(), echojwt.JWT([]byte(router.JWTKey)))
}

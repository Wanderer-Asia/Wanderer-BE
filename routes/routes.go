package routes

import (
	"wanderer/features/locations"

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
}

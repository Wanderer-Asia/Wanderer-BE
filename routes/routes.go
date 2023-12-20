package routes

import (
	"wanderer/features/airlines"
	"wanderer/features/bookings"
	"wanderer/features/facilities"
	"wanderer/features/locations"
	"wanderer/features/reviews"
	"wanderer/features/tours"
	"wanderer/features/users"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type Routes struct {
	JWTKey          string
	Server          *echo.Echo
	UserHandler     users.Handler
	AirlineHandler  airlines.Handler
	LocationHandler locations.Handler
	FacilityHandler facilities.Handler
	TourHandler     tours.Handler
	ReviewHandler   reviews.Handler
	BookingHandler  bookings.Handler
}

func (router Routes) InitRouter() {
	router.UserRouter()
	router.AirlineRouter()
	router.LocationRouter()
	router.FacilityRouter()
	router.TourRouter()
	router.ReviewRouter()
	router.BookingRouter()
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

func (router *Routes) FacilityRouter() {
	router.Server.POST("/facilities", router.FacilityHandler.Create(), echojwt.JWT([]byte(router.JWTKey)))
	router.Server.GET("/facilities", router.FacilityHandler.GetAll())
	router.Server.PUT("/facilities/:id", router.FacilityHandler.Update(), echojwt.JWT([]byte(router.JWTKey)))
	router.Server.DELETE("/facilities/:id", router.FacilityHandler.Delete(), echojwt.JWT([]byte(router.JWTKey)))
}

func (router *Routes) TourRouter() {
	router.Server.GET("/tours", router.TourHandler.GetAll())
	router.Server.POST("/tours", router.TourHandler.Create(), echojwt.JWT([]byte(router.JWTKey)))
	router.Server.PUT("/tours/:id", router.TourHandler.Update(), echojwt.JWT([]byte(router.JWTKey)))
	router.Server.GET("/tours/:id", router.TourHandler.GetDetail(), echojwt.JWT([]byte(router.JWTKey)))
}

func (router *Routes) ReviewRouter() {
	router.Server.POST("/reviews", router.ReviewHandler.Create(), echojwt.JWT([]byte(router.JWTKey)))
}

func (router *Routes) BookingRouter() {
	router.Server.POST("/bookings", router.BookingHandler.Create(), echojwt.JWT([]byte(router.JWTKey)))

	router.Server.POST("/payments", router.BookingHandler.PaymentNotification())
}

package main

import (
	"wanderer/config"
	"wanderer/helpers/encrypt"
	"wanderer/routes"
	"wanderer/utils/database"

	uh "wanderer/features/users/handler"
	ur "wanderer/features/users/repository"
	us "wanderer/features/users/service"

	ah "wanderer/features/airlines/handler"
	ar "wanderer/features/airlines/repository"
	as "wanderer/features/airlines/service"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	var dbConfig = new(config.DatabaseMysql)
	if err := dbConfig.LoadFromEnv(); err != nil {
		panic(err)
	}

	dbConnection, err := database.MysqlInit(*dbConfig)
	if err != nil {
		panic(err)
	}

	if err := database.MysqlMigrate(dbConnection); err != nil {
		panic(err)
	}

	var cldConfig = new(config.Cloudinary)
	if err := cldConfig.LoadFromEnv(); err != nil {
		panic(err)
	}

	cld, err := cloudinary.NewFromParams(cldConfig.CloudName, cldConfig.ApiKey, cldConfig.ApiSecret)
	if err != nil {
		panic(err)
	}

	enc := encrypt.NewBcrypt(10)

	userRepository := ur.NewUserRepository(dbConnection, cld)
	userService := us.NewUserService(userRepository, enc)
	userHandler := uh.NewUserHandler(userService)

	airlineRepository := ar.NewAirlineRepository(dbConnection, cld)
	airlineService := as.NewAirlineService(airlineRepository)
	airlineHandler := ah.NewAirlineHandler(airlineService)

	app := echo.New()
	app.Use(middleware.Recover())
	app.Use(middleware.CORS())

	route := routes.Routes{
		Server:         app,
		UserHandler:    userHandler,
		AirlineHandler: airlineHandler,
	}

	route.InitRouter()

	app.Logger.Fatal(app.Start(":8000"))
}

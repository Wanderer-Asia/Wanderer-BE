package main

import (
	"wanderer/config"
	"wanderer/helpers/encrypt"
	"wanderer/routes"
	"wanderer/utils/database"
	"wanderer/utils/files"

	uh "wanderer/features/users/handler"
	ur "wanderer/features/users/repository"
	us "wanderer/features/users/service"

	ah "wanderer/features/airlines/handler"
	ar "wanderer/features/airlines/repository"
	as "wanderer/features/airlines/service"

	lh "wanderer/features/locations/handler"
	lr "wanderer/features/locations/repository"
	ls "wanderer/features/locations/service"

	fh "wanderer/features/facilities/handler"
	fr "wanderer/features/facilities/repository"
	fs "wanderer/features/facilities/service"

	th "wanderer/features/tours/handler"
	tr "wanderer/features/tours/repository"
	ts "wanderer/features/tours/service"

	rh "wanderer/features/reviews/handler"
	rr "wanderer/features/reviews/repository"
	rs "wanderer/features/reviews/service"

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

	var jwtConfig = new(config.JWT)
	if err := jwtConfig.LoadFromEnv(); err != nil {
		panic(err)
	}

	var cldConfig = new(config.Cloudinary)
	if err := cldConfig.LoadFromEnv(); err != nil {
		panic(err)
	}

	cld, err := files.NewCloudinary(*cldConfig)
	if err != nil {
		panic(err)
	}

	enc := encrypt.NewBcrypt(10)

	userRepository := ur.NewUserRepository(dbConnection, cld)
	userService := us.NewUserService(userRepository, enc)
	userHandler := uh.NewUserHandler(userService, *jwtConfig)

	airlineRepository := ar.NewAirlineRepository(dbConnection, cld)
	airlineService := as.NewAirlineService(airlineRepository)
	airlineHandler := ah.NewAirlineHandler(airlineService)

	locationRepository := lr.NewLocationRepository(dbConnection, cld)
	locationService := ls.NewLocationService(locationRepository)
	locationHandler := lh.NewLocationHandler(locationService)

	facilityRepository := fr.NewFacilityRepository(dbConnection)
	facilityService := fs.NewFacilityService(facilityRepository)
	facilityHandler := fh.NewFacilityHandler(facilityService)

	tourRepository := tr.NewTourRepository(dbConnection, cld)
	tourService := ts.NewTourService(tourRepository)
	tourHandler := th.NewTourHandler(tourService)

	reviewRepository := rr.NewReviewRepository(dbConnection)
	reviewService := rs.NewReviewService(reviewRepository)
	reviewHandler := rh.NewReviewHandler(reviewService, *jwtConfig)

	app := echo.New()
	app.Use(middleware.Recover())
	app.Use(middleware.CORS())

	route := routes.Routes{
		JWTKey:          jwtConfig.Secret,
		Server:          app,
		UserHandler:     userHandler,
		AirlineHandler:  airlineHandler,
		LocationHandler: locationHandler,
		FacilityHandler: facilityHandler,
		TourHandler:     tourHandler,
		ReviewHandler:   reviewHandler,
	}

	route.InitRouter()

	app.Logger.Fatal(app.Start(":8000"))
}

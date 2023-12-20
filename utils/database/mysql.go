package database

import (
	"fmt"
	"wanderer/config"

	ar "wanderer/features/airlines/repository"
	br "wanderer/features/bookings/repository"
	fr "wanderer/features/facilities/repository"
	lr "wanderer/features/locations/repository"
	rr "wanderer/features/reviews/repository"
	tr "wanderer/features/tours/repository"
	ur "wanderer/features/users/repository"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlInit(cfg config.DatabaseMysql) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func MysqlMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&ur.User{},
		&ar.Airline{},
		&lr.Location{},
		&fr.Facility{},
		&tr.File{},
		&tr.Tour{},
		&tr.Itinerary{},
		&rr.Review{},
		&br.Booking{},
		&br.BookingDetail{},
	)

	if err != nil {
		return err
	}

	return nil
}

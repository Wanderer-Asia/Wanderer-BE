package repository

import (
	"io"
	"time"
	ar "wanderer/features/airlines/repository"
	fr "wanderer/features/facilities/repository"
	lr "wanderer/features/locations/repository"

	"gorm.io/gorm"
)

type Tour struct {
	Id          uint      `gorm:"column:id; primaryKey;"`
	Title       string    `gorm:"column:title; type:varchar(200); index;"`
	Description string    `gorm:"column:description; type:text; index;"`
	Price       float64   `gorm:"column:price; type:decimal(16,2); index;"`
	AdminFee    float64   `gorm:"column:admin_fee; type:decimal(16,2); index;"`
	Discount    int       `gorm:"column:discount; index;"`
	Start       time.Time `gorm:"column:start; type:timestamp; index;"`
	Finish      time.Time `gorm:"column:finish; type:timestamp; index;"`
	Quota       int       `gorm:"column:quota; index;"`
	Available   int       `gorm:"column:available; index;"`
	Rating      float32   `gorm:"column:rating; type:float; index;"`

	ThumbnailUrl string    `gorm:"column:thumbnail; type:text;"`
	ThumbnailRaw io.Reader `gorm:"-"`

	Picture []File `gorm:"many2many:tour_attachment"`

	Facility []fr.Facility `gorm:"many2many:tour_facility"`

	Itinerary []Itinerary

	AirlineId uint
	Airline   ar.Airline

	LocationId uint
	Location   lr.Location

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type File struct {
	Id int `gorm:"column:id; primaryKey;"`

	Raw io.Reader `gorm:"-"`
	Url string    `gorm:"column:file; type:text;"`

	CreatedAt time.Time
}

type Itinerary struct {
	Id          int    `gorm:"column:id; primaryKey;"`
	Location    string `gorm:"column:location; type:varchar(200); index;"`
	Description string `gorm:"column:description; type:text; index;"`

	TourId uint

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

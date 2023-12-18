package repository

import (
	"io"
	"time"
	ar "wanderer/features/airlines/repository"
	fr "wanderer/features/facilities/repository"
	lr "wanderer/features/locations/repository"
	"wanderer/features/tours"

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

	Itinerary []Itinerary `gorm:"foreignKey:TourId"`

	AirlineId uint
	Airline   ar.Airline

	LocationId uint
	Location   lr.Location

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (mod *Tour) FromEntity(ent tours.Tour) {
	if ent.Id != 0 {
		mod.Id = ent.Id
	}

	if ent.Title != "" {
		mod.Title = ent.Title
	}

	if ent.Description != "" {
		mod.Description = ent.Description
	}

	if ent.Price != 0 {
		mod.Price = ent.Price
	}

	if ent.AdminFee != 0 {
		mod.AdminFee = ent.AdminFee
	}

	if ent.Discount != 0 {
		mod.Discount = ent.Discount
	}

	if !ent.Start.IsZero() {
		mod.Start = ent.Start
	}

	if !ent.Finish.IsZero() {
		mod.Finish = ent.Finish
	}

	if ent.Quota != 0 {
		mod.Quota = ent.Quota
	}

	if ent.Thumbnail.Raw != nil {
		mod.ThumbnailRaw = ent.Thumbnail.Raw
	}

	for _, picture := range ent.Picture {
		var modPicture = new(File)
		modPicture.FromEntity(picture)
		mod.Picture = append(mod.Picture, *modPicture)
	}

	for _, facility := range ent.Facility {
		if facility.Id != 0 {
			mod.Facility = append(mod.Facility, fr.Facility{Id: facility.Id})
		}
	}

	for _, it := range ent.Itinerary {
		var modItinerary = new(Itinerary)
		modItinerary.FromEntity(it)
		mod.Itinerary = append(mod.Itinerary, *modItinerary)
	}

	if ent.Airline.Id != 0 {
		mod.AirlineId = ent.Airline.Id
	}

	if ent.Location.Id != 0 {
		mod.LocationId = ent.Location.Id
	}
}

type File struct {
	Id int `gorm:"column:id; primaryKey;"`

	Raw io.Reader `gorm:"-"`
	Url string    `gorm:"column:file; type:text;"`

	CreatedAt time.Time
}

func (mod *File) FromEntity(ent tours.File) {
	if ent.Raw != nil {
		mod.Raw = ent.Raw
	}
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

func (mod *Itinerary) FromEntity(ent tours.Itinerary) {
	if ent.Location != "" {
		mod.Location = ent.Location
	}

	if ent.Description != "" {
		mod.Description = ent.Description
	}
}

package repository

import (
	"fmt"
	"reflect"
	"time"
	"wanderer/features/reports"
)

type Booking struct {
	Code   int
	Status string

	TourId uint
	Tour   Tour
}

func (mod *Booking) ToEntity() *reports.Booking {
	var ent = new(reports.Booking)

	if mod.Code != 0 {
		ent.Code = mod.Code
	}

	if !reflect.ValueOf(mod.Tour).IsZero() {
		if mod.Tour.Price != 0 {
			ent.Price = mod.Tour.Price
		}

		if !reflect.ValueOf(mod.Tour.Location).IsZero() && mod.Tour.Location.Name != "" {
			ent.Location = mod.Tour.Location.Name
		}
	}

	return ent
}

type User struct {
	Id   uint
	Role string
}

type Tour struct {
	Id       uint
	Title    string
	Price    float64
	Discount int
	Start    time.Time
	Quota    int
	Rating   float32

	Thumbnail string

	LocationId uint
	Location   Location
}

func (mod *Tour) ToEntity() *reports.Tour {
	var ent = new(reports.Tour)

	if mod.Id != 0 {
		ent.Id = mod.Id
	}

	if mod.Title != "" {
		ent.Title = mod.Title
	}

	if mod.Price != 0 {
		ent.Price = mod.Price
	}

	if mod.Discount != 0 {
		ent.Discount = mod.Discount
	}

	if !mod.Start.IsZero() {
		ent.Start = mod.Start
	}

	if mod.Quota != 0 {
		ent.Quota = mod.Quota
	}

	if mod.Thumbnail != "" {
		ent.Thumbnail = mod.Thumbnail
	}

	if !reflect.ValueOf(mod.Location).IsZero() {
		ent.Location = *mod.Location.ToEntity()
	}

	return ent
}

type Location struct {
	Id   uint
	Name string
}

func (mod *Location) ToEntity() *reports.Location {
	var ent = new(reports.Location)

	if mod.Name != "" {
		ent.Name = mod.Name
	}

	return ent
}

type GraphBooking struct {
	Label int
	Total int
}

func (mod *GraphBooking) ToEntity() (*reports.GraphBooking, error) {
	var ent = new(reports.GraphBooking)

	if mod.Label != 0 {
		label, err := time.Parse("1", fmt.Sprintf("%d", mod.Label))
		if err != nil {
			return nil, err
		}

		ent.Label = label.Format("Jan")
	}

	if mod.Total != 0 {
		ent.Total = mod.Total
	}

	return ent, nil
}

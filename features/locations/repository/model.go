package repository

import (
	"io"
	"time"
	"wanderer/features/locations"
)

type Location struct {
	Id   uint   `gorm:"column:id; primaryKey;"`
	Name string `gorm:"column:name; type:varchar(200); unique;"`

	ImageUrl string    `gorm:"column:image; type:text;"`
	ImageRaw io.Reader `gorm:"-"`

	Tours []Tour

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (mod *Location) ToEntity() *locations.Location {
	var ent = new(locations.Location)

	if mod.Id != 0 {
		ent.Id = mod.Id
	}

	if mod.Name != "" {
		ent.Name = mod.Name
	}

	if mod.ImageUrl != "" {
		ent.ImageUrl = mod.ImageUrl
	}

	for _, tour := range mod.Tours {
		ent.Tours = append(ent.Tours, *tour.ToEntity())
	}

	if !mod.CreatedAt.IsZero() {
		ent.CreatedAt = mod.CreatedAt
	}

	if !mod.UpdatedAt.IsZero() {
		ent.UpdatedAt = mod.UpdatedAt
	}

	return ent
}

func (mod *Location) FromEntity(ent locations.Location) {
	if ent.Id != 0 {
		mod.Id = ent.Id
	}

	if ent.Name != "" {
		mod.Name = ent.Name
	}

	if ent.ImageUrl != "" {
		mod.ImageUrl = ent.ImageUrl
	} else if ent.ImageRaw != nil {
		mod.ImageRaw = ent.ImageRaw
	}

	if !ent.CreatedAt.IsZero() {
		mod.CreatedAt = ent.CreatedAt
	}

	if !ent.UpdatedAt.IsZero() {
		mod.UpdatedAt = ent.UpdatedAt
	}
}

type Tour struct {
	Id       uint
	Title    string
	Discount int
	Start    time.Time
	Quota    int
	Rating   float32

	Thumbnail string

	LocationId uint
}

func (mod *Tour) ToEntity() *locations.Tour {
	var ent = new(locations.Tour)

	if mod.Id != 0 {
		ent.Id = mod.Id
	}

	if mod.Title != "" {
		ent.Title = mod.Title
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

	if mod.Rating != 0 {
		ent.Rating = mod.Rating
	}

	if mod.Thumbnail != "" {
		ent.Thumbnail = mod.Thumbnail
	}

	return ent
}

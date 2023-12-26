package repository

import (
	"time"
	"wanderer/features/airlines"
)

type Airline struct {
	Id    uint   `gorm:"column:id; primaryKey;"`
	Name  string `gorm:"column:name; type:varchar(55); unique;"`
	Image string `gorm:"column:image; type:text; default:null;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (mod *Airline) FromEntity(ent airlines.Airline) {
	if ent.Name != "" {
		mod.Name = ent.Name
	}

	if ent.ImageUrl != "" {
		mod.Image = ent.ImageUrl
	}
}

func (mod *Airline) ToEntity() *airlines.Airline {
	var ent = new(airlines.Airline)

	if mod.Id != 0 {
		ent.Id = mod.Id
	}

	if mod.Name != "" {
		ent.Name = mod.Name
	}

	if mod.Image != "" {
		ent.ImageUrl = mod.Image
	}

	if !mod.CreatedAt.IsZero() {
		ent.CreatedAt = mod.CreatedAt
	}

	if !mod.UpdatedAt.IsZero() {
		ent.UpdatedAt = mod.UpdatedAt
	}

	return ent
}

package repository

import (
	"time"
	"wanderer/features/locations"

	"gorm.io/gorm"
)

type Location struct {
	Id   uint   `gorm:"column:id; primaryKey;"`
	Name string `gorm:"column:name; type:varchar(200);"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (mod *Location) ToEntity() *locations.Location {
	var ent = new(locations.Location)

	if mod.Id != 0 {
		ent.Id = mod.Id
	}

	if mod.Name != "" {
		ent.Name = mod.Name
	}

	if !mod.CreatedAt.IsZero() {
		ent.CreatedAt = mod.CreatedAt
	}

	if !mod.UpdatedAt.IsZero() {
		ent.UpdatedAt = mod.UpdatedAt
	}

	if !mod.DeletedAt.Time.IsZero() {
		ent.DeletedAt = mod.DeletedAt.Time
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

	if !ent.CreatedAt.IsZero() {
		mod.CreatedAt = ent.CreatedAt
	}

	if !ent.UpdatedAt.IsZero() {
		mod.UpdatedAt = ent.UpdatedAt
	}

	if !ent.DeletedAt.IsZero() {
		mod.DeletedAt = gorm.DeletedAt{Time: ent.DeletedAt}
	}
}

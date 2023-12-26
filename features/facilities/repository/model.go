package repository

import (
	"time"
	"wanderer/features/facilities"
)

type Facility struct {
	Id   uint   `gorm:"column:id; primaryKey;"`
	Name string `gorm:"column:name; type:varchar(200); unique;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (mod *Facility) FromEntity(ent facilities.Facility) {
	if ent.Name != "" {
		mod.Name = ent.Name
	}

}

func (mod *Facility) ToEntity() *facilities.Facility {
	var ent = new(facilities.Facility)

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

	return ent
}

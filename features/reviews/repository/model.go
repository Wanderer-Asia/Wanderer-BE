package repository

import (
	"time"
	"wanderer/features/reviews"
)

type Review struct {
	Id     uint    `gorm:"column:id; primaryKey;"`
	UserId uint    `gorm:"column:user_id;"`
	Text   string  `gorm:"column:text; type:varchar(255);"`
	Rating float32 `gorm:"column:rating; type:float;"`

	CreatedAt time.Time
}

func (mod *Review) FromEntity(ent reviews.Review) {
	if ent.Text != "" {
		mod.Text = ent.Text
	}

	if ent.Rating != 0 {
		mod.Rating = ent.Rating
	}
}

func (mod *Review) ToEntity() *reviews.Review {
	var ent = new(reviews.Review)

	if mod.Id != 0 {
		ent.Id = mod.Id
	}

	if mod.Text != "" {
		ent.Text = mod.Text
	}

	if mod.Rating != 0 {
		ent.Rating = mod.Rating
	}

	if !mod.CreatedAt.IsZero() {
		ent.CreatedAt = mod.CreatedAt
	}

	return ent
}

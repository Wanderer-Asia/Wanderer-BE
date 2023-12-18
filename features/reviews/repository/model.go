package repository

import (
	"reflect"
	"time"
	"wanderer/features/reviews"
	ur "wanderer/features/users/repository"
)

type Review struct {
	Id     uint    `gorm:"column:id; primaryKey;"`
	UserId uint    `gorm:"column:user_id;"`
	User   ur.User `gorm:"foreignkey:UserId;"`
	TourId uint    `gorm:"column:tour_id;"`
	Text   string  `gorm:"column:text; type:text;"`
	Rating float32 `gorm:"column:rating; type:float(8,2);"`

	CreatedAt time.Time
}

func (mod *Review) FromEntity(ent reviews.Review) {
	if ent.Text != "" {
		mod.Text = ent.Text
	}

	if ent.User.Id != 0 {
		mod.UserId = ent.User.Id
	}

	if ent.TourId != 0 {
		mod.TourId = ent.TourId
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

	if mod.UserId != 0 {
		ent.User.Id = mod.UserId
	}

	if !reflect.ValueOf(mod.User).IsZero() {
		ent.User = *mod.User.ToEntity()
	}

	if mod.TourId != 0 {
		ent.TourId = mod.TourId
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

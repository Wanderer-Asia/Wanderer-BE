package repository

import (
	"reflect"
	"time"
	"wanderer/features/users"

	"gorm.io/gorm"
)

type User struct {
	Id          uint   `gorm:"column:id; primaryKey;"`
	Name        string `gorm:"column:fullname; type:varchar(200);"`
	Phone       string `gorm:"column:phone; type:varchar(20);"`
	Email       string `gorm:"column:email; type:varchar(255); unique;"`
	Password    string `gorm:"column:password; type:varchar(72); not null;"`
	Image       string `gorm:"column:image; type:text;"`
	Role        string `gorm:"column:role; type:enum('admin', 'user');"`
	TourCount   int
	ReviewCount int

	Bookings []Booking `gorm:"foreignKey:UserId"`
	Reviews  []Review  `gorm:"foreignKey:UserId"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (mod *User) FromEntity(ent users.User) {
	if ent.Name != "" {
		mod.Name = ent.Name
	}

	if ent.Phone != "" {
		mod.Phone = ent.Phone
	}

	if ent.Email != "" {
		mod.Email = ent.Email
	}

	if ent.Password != "" {
		mod.Password = ent.Password
	}

	if ent.ImageUrl != "" {
		mod.Image = ent.ImageUrl
	}

	if ent.Role != "" {
		mod.Role = ent.Role
	}
}

func (mod *User) ToEntity() *users.User {
	var ent = new(users.User)

	if mod.Id != 0 {
		ent.Id = mod.Id
	}

	if mod.Id != 0 {
		ent.Id = mod.Id
	}

	if mod.Name != "" {
		ent.Name = mod.Name
	}

	if mod.Phone != "" {
		ent.Phone = mod.Phone
	}

	if mod.Email != "" {
		ent.Email = mod.Email
	}

	if mod.Password != "" {
		ent.Password = mod.Password
	}

	if mod.Image != "" {
		ent.ImageUrl = mod.Image
	}

	if mod.Role != "" {
		ent.Role = mod.Role
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

	for _, booking := range mod.Bookings {
		if !reflect.ValueOf(booking).IsZero() {
			ent.Bookings = append(ent.Bookings, *booking.ToEntity())
		}
	}

	for _, review := range mod.Reviews {
		if !reflect.ValueOf(review).IsZero() {
			ent.Reviews = append(ent.Reviews, *review.ToEntity())
		}
	}

	return ent
}

type Booking struct {
	Code   int `gorm:"primaryKey"`
	Status string
	UserId uint

	BookingDetails []BookingDetail `gorm:"foreignKey:BookingCode"`
	DetailCount    int

	TourId uint
	Tour   Tour
}

func (mod *Booking) FromEntity(ent users.Booking) {
	if ent.Status != "" {
		mod.Status = ent.Status
	}

	if ent.UserId != 0 {
		mod.UserId = ent.UserId
	}

	if ent.DetailCount != 0 {
		mod.DetailCount = ent.DetailCount
	}
}

func (mod *Booking) ToEntity() *users.Booking {
	var ent = new(users.Booking)

	if mod.Code != 0 {
		ent.Code = mod.Code
	}

	if mod.Status != "" {
		ent.Status = mod.Status
	}

	if mod.UserId != 0 {
		ent.UserId = mod.UserId
	}

	ent.Tour = *mod.Tour.ToEntity()

	for _, bookingDetail := range mod.BookingDetails {
		if !reflect.ValueOf(bookingDetail).IsZero() {
			ent.BookingDetails = append(ent.BookingDetails, *bookingDetail.ToEntity())
		}
	}

	return ent
}

type Tour struct {
	Id    uint `gorm:"primaryKey;"`
	Title string
}

func (mod *Tour) FromEntity(ent users.Tour) {
	if ent.Title != "" {
		mod.Title = ent.Title
	}
	if ent.Id != 0 {
		mod.Id = ent.Id
	}
}

func (mod *Tour) ToEntity() *users.Tour {
	var ent = new(users.Tour)

	if mod.Id != 0 {
		ent.Id = mod.Id
	}

	if mod.Title != "" {
		ent.Title = mod.Title
	}

	return ent
}

type Review struct {
	Id     uint `gorm:"primaryKey;"`
	UserId uint
}

func (mod *Review) FromEntity(ent users.Review) {
	if ent.UserId != 0 {
		mod.UserId = ent.UserId
	}
	if ent.Id != 0 {
		mod.Id = ent.Id
	}
}

func (mod *Review) ToEntity() *users.Review {
	var ent = new(users.Review)

	if mod.Id != 0 {
		ent.Id = mod.Id
	}

	if mod.UserId != 0 {
		ent.UserId = mod.UserId
	}

	return ent
}

type BookingDetail struct {
	Id          uint `gorm:"primaryKey"`
	BookingCode int
}

func (mod *BookingDetail) FromEntity(ent users.BookingDetail) {
	if ent.BookingCode != 0 {
		mod.BookingCode = ent.BookingCode
	}
	if ent.Id != 0 {
		mod.Id = ent.Id
	}
}

func (mod *BookingDetail) ToEntity() *users.BookingDetail {
	var ent = new(users.BookingDetail)

	if mod.Id != 0 {
		ent.Id = mod.Id
	}

	if mod.BookingCode != 0 {
		ent.BookingCode = mod.BookingCode
	}

	return ent
}

package repository

import (
	"reflect"
	"time"
	"wanderer/features/users"

	"gorm.io/gorm"
)

type User struct {
	Id       uint   `gorm:"column:id; primaryKey;"`
	Name     string `gorm:"column:fullname; type:varchar(200);"`
	Phone    string `gorm:"column:phone; type:varchar(20);"`
	Email    string `gorm:"column:email; type:varchar(255); unique;"`
	Password string `gorm:"column:password; type:varchar(72); not null;"`
	Image    string `gorm:"column:image; type:text; default:null;"`
	Role     string `gorm:"column:role; type:enum('admin', 'user');"`

	TourCount   int       `gorm:"-"`
	ReviewCount int       `gorm:"-"`
	Bookings    []Booking `gorm:"-"`

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

	if mod.ReviewCount != 0 {
		ent.ReviewCount = mod.ReviewCount
	}

	if mod.TourCount != 0 {
		ent.TourCount = mod.TourCount
	}

	for _, booking := range mod.Bookings {
		ent.Bookings = append(ent.Bookings, *booking.ToEntity())
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

type Booking struct {
	Code   int
	Status string

	DetailCount int

	UserId uint

	TourId uint
	Tour   Tour
}

func (mod *Booking) ToEntity() *users.Booking {
	var ent = new(users.Booking)

	if mod.Code != 0 {
		ent.Code = mod.Code
	}

	if mod.Status != "" {
		ent.Status = mod.Status
	}

	if mod.DetailCount != 0 {
		ent.DetailCount = mod.DetailCount
	}

	if !reflect.ValueOf(mod.Tour).IsZero() {
		ent.Tour = *mod.Tour.ToEntity()
	}

	return ent
}

type Tour struct {
	Id    uint
	Title string
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
	Id     uint
	UserId uint
}

package repository

import (
	"io"
	"reflect"
	"time"
	"wanderer/features/tours"

	"gorm.io/gorm"
)

type Tour struct {
	Id          uint      `gorm:"column:id; primaryKey;"`
	Title       string    `gorm:"column:title; type:varchar(200); index;"`
	Description string    `gorm:"column:description; type:text;"`
	Price       float64   `gorm:"column:price; type:decimal(16,2); index;"`
	AdminFee    float64   `gorm:"column:admin_fee; type:decimal(16,2);"`
	Discount    int       `gorm:"column:discount; index;"`
	Start       time.Time `gorm:"column:start; type:timestamp;"`
	Finish      time.Time `gorm:"column:finish; type:timestamp;"`
	Quota       int       `gorm:"column:quota;"`
	Available   int       `gorm:"column:available;"`
	Rating      float32   `gorm:"column:rating; type:float; index;"`

	ThumbnailUrl string    `gorm:"column:thumbnail; type:text;"`
	ThumbnailRaw io.Reader `gorm:"-"`

	Picture []File `gorm:"many2many:tour_attachment"`

	Facility []Facility `gorm:"many2many:tour_facility"`

	Itinerary []Itinerary `gorm:"foreignKey:TourId"`

	AirlineId uint
	Airline   Airline

	LocationId uint
	Location   Location

	Reviews []Review `gorm:"foreignKey:TourId"`

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

	for _, facility := range ent.FacilityInclude {
		if facility.Id != 0 {
			mod.Facility = append(mod.Facility, Facility{Id: facility.Id})
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

func (mod *Tour) ToEntity(excludeFacility []Facility) *tours.Tour {
	var ent = new(tours.Tour)

	if mod.Id != 0 {
		ent.Id = mod.Id
	}

	if mod.Title != "" {
		ent.Title = mod.Title
	}

	if mod.Description != "" {
		ent.Description = mod.Description
	}

	if mod.Price != 0 {
		ent.Price = mod.Price
	}

	if mod.AdminFee != 0 {
		ent.AdminFee = mod.AdminFee
	}

	if mod.Discount != 0 {
		ent.Discount = mod.Discount
	}

	if !mod.Start.IsZero() {
		ent.Start = mod.Start
	}

	if !mod.Finish.IsZero() {
		ent.Finish = mod.Finish
	}

	if mod.Quota != 0 {
		ent.Quota = mod.Quota
	}

	ent.Available = mod.Available
	ent.Rating = mod.Rating

	if mod.ThumbnailUrl != "" {
		ent.Thumbnail.Url = mod.ThumbnailUrl
	}

	for _, pict := range mod.Picture {
		if !reflect.ValueOf(pict).IsZero() {
			ent.Picture = append(ent.Picture, pict.ToEntity())
		}
	}

	for _, fac := range mod.Facility {
		if !reflect.ValueOf(fac).IsZero() {
			ent.FacilityInclude = append(ent.FacilityInclude, *fac.ToEntity())
		}
	}

	for _, fac := range excludeFacility {
		if !reflect.ValueOf(fac).IsZero() {
			ent.FacilityExclude = append(ent.FacilityExclude, *fac.ToEntity())
		}
	}

	for _, it := range mod.Itinerary {
		if !reflect.ValueOf(it).IsZero() {
			ent.Itinerary = append(ent.Itinerary, it.ToEntity())
		}
	}

	if !reflect.ValueOf(mod.Airline).IsZero() {
		ent.Airline = mod.Airline.ToEntity()
	}

	if !reflect.ValueOf(mod.Location).IsZero() {
		ent.Location = mod.Location.ToEntity()
	}

	for _, review := range mod.Reviews {
		ent.Reviews = append(ent.Reviews, *review.ToEntity())
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

func (mod *File) ToEntity() tours.File {
	var ent = new(tours.File)

	if mod.Id != 0 {
		ent.Id = mod.Id
	}

	if mod.Url != "" {
		ent.Url = mod.Url
	}

	if !mod.CreatedAt.IsZero() {
		ent.CreatedAt = mod.CreatedAt
	}

	return *ent
}

type Itinerary struct {
	Id          int    `gorm:"column:id; primaryKey;"`
	Location    string `gorm:"column:location; type:varchar(200);"`
	Description string `gorm:"column:description; type:text;"`

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

func (mod *Itinerary) ToEntity() tours.Itinerary {
	var ent = new(tours.Itinerary)

	if mod.Id != 0 {
		ent.Id = mod.Id
	}

	if mod.Location != "" {
		ent.Location = mod.Location
	}

	if mod.Description != "" {
		ent.Description = mod.Description
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

	return *ent
}

type Facility struct {
	Id   uint
	Name string
}

func (mod *Facility) ToEntity() *tours.Facility {
	var ent = new(tours.Facility)

	if mod.Id != 0 {
		ent.Id = mod.Id
	}

	if mod.Name != "" {
		ent.Name = mod.Name
	}

	return ent
}

type Airline struct {
	Id   uint
	Name string
}

func (mod *Airline) ToEntity() tours.Airline {
	var ent = new(tours.Airline)

	if mod.Id != 0 {
		ent.Id = mod.Id
	}

	if mod.Name != "" {
		ent.Name = mod.Name
	}

	return *ent
}

type Location struct {
	Id   uint
	Name string
}

func (mod *Location) ToEntity() tours.Location {
	var ent = new(tours.Location)

	if mod.Id != 0 {
		ent.Id = mod.Id
	}

	if mod.Name != "" {
		ent.Name = mod.Name
	}

	return *ent
}

type Review struct {
	Id        uint
	Text      string
	Rating    float32
	CreatedAt time.Time
	UserId    uint
	User      User
	TourId    uint
}

func (mod *Review) ToEntity() *tours.Review {
	var ent = new(tours.Review)

	if mod.Id != 0 {
		ent.Id = mod.Id
	}

	if mod.Text != "" {
		ent.Text = mod.Text
	}

	if mod.Rating != 0 {
		ent.Rating = mod.Rating
	}

	if !reflect.ValueOf(mod.User).IsZero() {
		ent.User = mod.User.ToEntity()
	}

	if !mod.CreatedAt.IsZero() {
		ent.CreatedAt = mod.CreatedAt
	}

	return ent
}

type User struct {
	Id    uint
	Name  string `gorm:"column:fullname;"`
	Image string
}

func (mod *User) ToEntity() tours.User {
	var ent = new(tours.User)

	if mod.Id != 0 {
		ent.Id = mod.Id
	}

	if mod.Name != "" {
		ent.Name = mod.Name
	}

	if mod.Image != "" {
		ent.Image = mod.Image
	}

	return *ent
}

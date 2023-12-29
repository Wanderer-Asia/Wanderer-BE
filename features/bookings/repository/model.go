package repository

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
	"wanderer/features/bookings"

	"gorm.io/gorm"
)

type Booking struct {
	Code      int            `gorm:"column:code; primaryKey;"`
	Total     float64        `gorm:"column:total;"`
	Status    string         `gorm:"column:status; type:enum('pending', 'cancel', 'approved', 'refund', 'refunded'); default:'pending'; index;"`
	BookedAt  time.Time      `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	UserId uint
	User   User `gorm:"foreignKey:UserId"`

	TourId uint
	Tour   Tour `gorm:"foreignKey:TourId"`

	Detail  []BookingDetail
	Payment Payment `gorm:"embedded;embeddedPrefix:payment_"`
}

func (mod *Booking) GenerateCode() (err error) {
	mod.Code, err = strconv.Atoi(fmt.Sprintf("%d%d%d", mod.UserId, mod.TourId, time.Now().Unix()))
	if err != nil {
		return err
	}

	return
}

func (mod *Booking) CalcTotal(tour Tour) {
	mod.Total = (float64(len(mod.Detail)) * tour.Price) - (float64(tour.Discount) / 100 * tour.Price * float64(len(mod.Detail))) + tour.AdminFee
}

func (mod *Booking) FromEntity(ent bookings.Booking) {
	if ent.Tour.Id != 0 {
		mod.TourId = ent.Tour.Id
	}

	if ent.User.Id != 0 {
		mod.UserId = ent.User.Id
	}

	for _, detail := range ent.Detail {
		var tmpDetail = new(BookingDetail)
		tmpDetail.FromEntity(detail)

		mod.Detail = append(mod.Detail, *tmpDetail)
	}

	if ent.Status != "" {
		mod.Status = ent.Status
	}

	if !reflect.ValueOf(ent.Payment).IsZero() {
		mod.Payment.FromEntity(ent.Payment)
	}
}

func (mod *Booking) ToEntity() *bookings.Booking {
	var ent = new(bookings.Booking)

	if mod.Code != 0 {
		ent.Code = mod.Code
	}

	if mod.Total != 0 {
		ent.Total = mod.Total
	}

	if mod.Status != "" {
		ent.Status = mod.Status
	}

	if !mod.BookedAt.IsZero() {
		ent.BookedAt = mod.BookedAt
	}

	if !mod.DeletedAt.Time.IsZero() {
		ent.DeletedAt = mod.DeletedAt.Time
	}

	if !reflect.ValueOf(mod.Tour).IsZero() {
		ent.Tour = *mod.Tour.ToEntity(nil)
	}

	if !reflect.ValueOf(mod.User).IsZero() {
		ent.User = *mod.User.ToEntity()
	}

	for _, detail := range mod.Detail {
		if !reflect.ValueOf(detail).IsZero() {
			ent.Detail = append(ent.Detail, detail.ToEntity())
		}
	}

	if !reflect.ValueOf(mod.Payment).IsZero() {
		ent.Payment = *mod.Payment.ToEntity()
	}

	return ent
}

type BookingDetail struct {
	Id             uint      `gorm:"column:id; primaryKey;"`
	DocumentNumber string    `gorm:"column:document_number; type:varchar(200);"`
	Greeting       string    `gorm:"column:greeting; type:varchar(10);"`
	Name           string    `gorm:"column:name; type:varchar(200);"`
	Nationality    string    `gorm:"column:nationality; type:varchar(100);"`
	DOB            time.Time `gorm:"column:dob;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	BookingCode int
}

func (mod *BookingDetail) FromEntity(ent bookings.Detail) {
	if ent.DocumentNumber != "" {
		mod.DocumentNumber = ent.DocumentNumber
	}

	if ent.Greeting != "" {
		mod.Greeting = ent.Greeting
	}

	if ent.Name != "" {
		mod.Name = ent.Name
	}

	if ent.Nationality != "" {
		mod.Nationality = ent.Nationality
	}

	if !ent.DOB.IsZero() {
		mod.DOB = ent.DOB
	}
}

func (mod *BookingDetail) ToEntity() bookings.Detail {
	var ent = new(bookings.Detail)

	if mod.Id != 0 {
		ent.Id = mod.Id
	}

	if mod.DocumentNumber != "" {
		ent.DocumentNumber = mod.DocumentNumber
	}

	if mod.Greeting != "" {
		ent.Greeting = mod.Greeting
	}

	if mod.Name != "" {
		ent.Name = mod.Name
	}

	if mod.Nationality != "" {
		ent.Nationality = mod.Nationality
	}

	if !mod.DOB.IsZero() {
		ent.DOB = mod.DOB
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

	if mod.BookingCode != 0 {
		ent.BookingCode = mod.BookingCode
	}

	return *ent
}

type Payment struct {
	Method        string `gorm:"column:method; type:varchar(20);"`
	Bank          string `gorm:"column:bank; type:varchar(20);"`
	VirtualNumber string `gorm:"column:virtual_number; type:varchar(50);"`
	BillKey       string `gorm:"column:bill_key; type:varchar(50);"`
	BillCode      string `gorm:"column:bill_code; type:varchar(50);"`
	Status        string `gorm:"column:status; type:varchar(20);"`

	CreatedAt time.Time `gorm:"index"`
	ExpiredAt time.Time `gorm:"nullable"`
	PaidAt    time.Time `gorm:"default:null;"`
}

func (mod *Payment) FromEntity(ent bookings.Payment) {
	if ent.Method != "" {
		mod.Method = ent.Method
	}

	if ent.Bank != "" {
		mod.Bank = ent.Bank
	}

	if ent.VirtualNumber != "" {
		mod.VirtualNumber = ent.VirtualNumber
	}

	if ent.BillKey != "" {
		mod.BillKey = ent.BillKey
	}

	if ent.BillCode != "" {
		mod.BillCode = ent.BillCode
	}

	if ent.Status != "" {
		mod.Status = ent.Status
	}

	if !ent.ExpiredAt.IsZero() {
		mod.ExpiredAt = ent.ExpiredAt
	}

	if !ent.PaidAt.IsZero() {
		mod.PaidAt = ent.PaidAt
	}
}

func (mod *Payment) ToEntity() *bookings.Payment {
	var ent = new(bookings.Payment)

	if mod.Method != "" {
		ent.Method = mod.Method
	}

	if mod.Bank != "" {
		ent.Bank = mod.Bank
	}

	if mod.VirtualNumber != "" {
		ent.VirtualNumber = mod.VirtualNumber
	}

	if mod.BillKey != "" {
		ent.BillKey = mod.BillKey
	}

	if mod.BillCode != "" {
		ent.BillCode = mod.BillCode
	}

	if mod.Status != "" {
		ent.Status = mod.Status
	}

	if !mod.CreatedAt.IsZero() {
		ent.CreatedAt = mod.CreatedAt
	}

	if !mod.ExpiredAt.IsZero() {
		ent.ExpiredAt = mod.ExpiredAt
	}

	if !mod.PaidAt.IsZero() {
		ent.PaidAt = mod.PaidAt
	}

	return ent
}

type User struct {
	Id    uint
	Name  string `gorm:"column:fullname;"`
	Email string
	Phone string

	Image string
}

func (mod *User) ToEntity() *bookings.User {
	var ent = new(bookings.User)

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

	return ent
}

type Tour struct {
	Id          uint
	Title       string
	Description string
	Price       float64
	AdminFee    float64
	Discount    int
	Start       time.Time
	Finish      time.Time
	Quota       int
	Available   int
	Rating      *float32

	Picture   []File `gorm:"many2many:tour_attachment;"`
	Itinerary []Itinerary
	Facility  []Facility `gorm:"many2many:tour_facility;"`

	AirlineId uint
	Airline   Airline

	LocationId uint
	Location   Location

	Reviews []Review `gorm:"foreignKey:TourId"`
}

func (mod *Tour) ToEntity(excludeFacility []Facility) *bookings.Tour {
	var ent = new(bookings.Tour)

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

	if !mod.Start.IsZero() {
		ent.Start = mod.Start
	}

	if !mod.Finish.IsZero() {
		ent.Finish = mod.Finish
	}

	if mod.Quota != 0 {
		ent.Quota = mod.Quota
	}

	if mod.Available != 0 {
		ent.Available = mod.Available
	}

	if mod.Rating != nil {
		ent.Rating = mod.Rating
	}

	if !reflect.ValueOf(mod.Airline).IsZero() {
		ent.Airline = *mod.Airline.ToEntity()
	}

	if !reflect.ValueOf(mod.Location).IsZero() {
		ent.Location = *mod.Location.ToEntity()
	}

	for _, pict := range mod.Picture {
		if !reflect.ValueOf(pict).IsZero() {
			ent.Picture = append(ent.Picture, *pict.ToEntity())
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
			ent.Itinerary = append(ent.Itinerary, *it.ToEntity())
		}
	}

	for _, rev := range mod.Reviews {
		if !reflect.ValueOf(rev).IsZero() {
			ent.Reviews = append(ent.Reviews, *rev.ToEntity())
		}
	}

	return ent
}

type File struct {
	Id  int
	Url string `gorm:"column:file;"`
}

func (mod *File) ToEntity() *bookings.File {
	var ent = new(bookings.File)

	if mod.Id != 0 {
		ent.Id = mod.Id
	}

	if mod.Url != "" {
		ent.Url = mod.Url
	}

	return ent
}

type Itinerary struct {
	Id          int
	Location    string
	Description string

	TourId uint
}

func (mod *Itinerary) ToEntity() *bookings.Itinerary {
	var ent = new(bookings.Itinerary)

	if mod.Id != 0 {
		ent.Id = mod.Id
	}

	if mod.Location != "" {
		ent.Location = mod.Location
	}

	if mod.Description != "" {
		ent.Description = mod.Description
	}

	return ent
}

type Facility struct {
	Id   uint
	Name string
}

func (mod *Facility) ToEntity() *bookings.Facility {
	var ent = new(bookings.Facility)

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

func (mod *Airline) ToEntity() *bookings.Airline {
	var ent = new(bookings.Airline)

	if mod.Id != 0 {
		ent.Id = mod.Id
	}

	if mod.Name != "" {
		ent.Name = mod.Name
	}

	return ent
}

type Location struct {
	Id   uint
	Name string
}

func (mod *Location) ToEntity() *bookings.Location {
	var ent = new(bookings.Location)

	if mod.Id != 0 {
		ent.Id = mod.Id
	}

	if mod.Name != "" {
		ent.Name = mod.Name
	}

	return ent
}

type Review struct {
	Id     uint
	UserId uint
	User   User `gorm:"foreignkey:UserId;"`
	TourId uint
	Text   string
	Rating float32

	CreatedAt time.Time
}

func (mod *Review) ToEntity() *bookings.Review {
	var ent = new(bookings.Review)

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
		ent.User = *mod.User.ToEntity()
	}

	if !mod.CreatedAt.IsZero() {
		ent.CreatedAt = mod.CreatedAt
	}

	return ent
}

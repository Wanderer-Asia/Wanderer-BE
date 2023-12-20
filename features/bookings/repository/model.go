package repository

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
	"wanderer/features/bookings"
	tr "wanderer/features/tours/repository"
	ur "wanderer/features/users/repository"

	"gorm.io/gorm"
)

type Booking struct {
	Code      int            `gorm:"column:code; primaryKey;"`
	Total     float64        `gorm:"column:total;"`
	Status    string         `gorm:"column:status; type:enum('pending', 'approved', 'refund', 'refunded'); default:'pending'; index;"`
	BookedAt  time.Time      `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	UserId uint
	User   ur.User `gorm:"foreignKey:UserId"`

	TourId uint
	Tour   tr.Tour `gorm:"foreignKey:TourId"`

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

func (mod *Booking) CalcTotal(tour tr.Tour) {
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
		ent.Payment = mod.Payment.ToEntity()
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
	ExpiredAt time.Time
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

func (mod *Payment) ToEntity() bookings.Payment {
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

	return *ent
}

package handler

import (
	"reflect"
	"time"
	"wanderer/features/bookings"
)

type BookingResponse struct {
	Code        int     `json:"booking_code,omitempty"`
	DetailCount int     `json:"detail_count,omitempty"`
	Status      string  `json:"status,omitempty"`
	Total       float64 `json:"total,omitempty"`

	PaymentBank          string     `json:"payment_method,omitempty"`
	PaymentVirtualNumber string     `json:"virtual_number,omitempty"`
	PaymentBillKey       string     `json:"key_bill,omitempty"`
	PaymentBillCode      string     `json:"code_bill,omitempty"`
	PaymentExpiredAt     *time.Time `json:"payment_expired,omitempty"`

	Tour *TourResponse `json:"tour,omitempty"`

	User *UserResponse `json:"user,omitempty"`
}

func (res *BookingResponse) FromEntity(ent bookings.Booking) {
	if ent.Code != 0 {
		res.Code = ent.Code
	}

	if len(ent.Detail) != 0 {
		res.DetailCount = len(ent.Detail)
	}

	if ent.Status != "" {
		res.Status = ent.Status
	}

	if ent.Total != 0 {
		res.Total = ent.Total
	}

	if !reflect.ValueOf(ent.Payment).IsZero() {
		if ent.Payment.Bank != "" {
			res.PaymentBank = ent.Payment.Bank
		}

		if ent.Payment.VirtualNumber != "" {
			res.PaymentVirtualNumber = ent.Payment.VirtualNumber
		}

		if ent.Payment.BillKey != "" {
			res.PaymentBillKey = ent.Payment.BillKey
		}

		if ent.Payment.BillCode != "" {
			res.PaymentBillCode = ent.Payment.BillCode
		}

		if !ent.Payment.ExpiredAt.IsZero() {
			res.PaymentExpiredAt = &ent.Payment.ExpiredAt
		}
	}

	if !reflect.ValueOf(ent.Tour).IsZero() {
		var tmpTour = new(TourResponse)
		tmpTour.FromEntity(ent.Tour)
		res.Tour = tmpTour
	}

	if !reflect.ValueOf(ent.User).IsZero() {
		var tmpUser = new(UserResponse)
		tmpUser.FromEntity(ent.User)
		res.User = tmpUser
	}
}

type TourResponse struct {
	Id          uint       `json:"tour_id,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Price       float64    `json:"price,omitempty"`
	Start       *time.Time `json:"start,omitempty"`
	Finish      *time.Time `json:"finish,omitempty"`
	Rating      *float32   `json:"rating,omitempty"`

	Picture []string `json:"picture,omitempty"`

	Facility *struct {
		Include []string `json:"include"`
		Exclude []string `json:"exclude"`
	} `json:"facility,omitempty"`

	Itinerary []ItineraryResponse `json:"itinerary,omitempty"`
	Airline   string              `json:"airline,omitempty"`
	Reviews   []ReviewResponse    `json:"reviews,omitempty"`
}

func (res *TourResponse) FromEntity(ent bookings.Tour) {
	if ent.Id != 0 {
		res.Id = ent.Id
	}

	if ent.Title != "" {
		res.Title = ent.Title
	}

	if ent.Description != "" {
		res.Description = ent.Description
	}

	if ent.Price != 0 {
		res.Price = ent.Price
	}

	if !ent.Start.IsZero() {
		res.Start = &ent.Start
	}

	if !ent.Finish.IsZero() {
		res.Finish = &ent.Finish
	}

	if ent.Rating != nil {
		res.Rating = ent.Rating
	}

	if !reflect.ValueOf(ent.Airline).IsZero() {
		res.Airline = ent.Airline.Name
	}

	for _, pict := range ent.Picture {
		if pict.Url != "" {
			res.Picture = append(res.Picture, pict.Url)
		}
	}

	if len(ent.FacilityInclude) != 0 || len(ent.FacilityExclude) != 0 {
		res.Facility = &struct {
			Include []string `json:"include"`
			Exclude []string `json:"exclude"`
		}{}
	}

	for _, fac := range ent.FacilityInclude {
		res.Facility.Include = append(res.Facility.Include, fac.Name)
	}

	for _, fac := range ent.FacilityExclude {
		res.Facility.Exclude = append(res.Facility.Exclude, fac.Name)
	}

	for _, it := range ent.Itinerary {
		var tmpItinerary = new(ItineraryResponse)
		tmpItinerary.FromEntity(it)

		res.Itinerary = append(res.Itinerary, *tmpItinerary)
	}

	for _, rev := range ent.Reviews {
		var tmpReview = new(ReviewResponse)
		tmpReview.FromEntity(rev)

		res.Reviews = append(res.Reviews, *tmpReview)
	}
}

type ItineraryResponse struct {
	Location    string `json:"location"`
	Description string `json:"description"`
}

func (res *ItineraryResponse) FromEntity(ent bookings.Itinerary) {
	res.Location = ent.Location
	res.Description = ent.Description
}

type ReviewResponse struct {
	User      UserResponse `json:"user"`
	Text      string       `json:"text,omitempty"`
	CreatedAt time.Time    `json:"created_at"`
}

func (res *ReviewResponse) FromEntity(ent bookings.Review) {
	if !reflect.ValueOf(ent.User).IsZero() {
		var tmpUser = new(UserResponse)
		tmpUser.FromEntity(ent.User)
		res.User = *tmpUser
	}

	if ent.Text != "" {
		res.Text = ent.Text
	}

	res.CreatedAt = ent.CreatedAt
}

type UserResponse struct {
	Id    uint   `json:"user_id,omitempty"`
	Name  string `json:"fullname,omitempty"`
	Image string `json:"image,omitempty"`
}

func (res *UserResponse) FromEntity(ent bookings.User) {
	if ent.Id != 0 {
		res.Id = ent.Id
	}

	if ent.Name != "" {
		res.Name = ent.Name
	}

	if ent.ImageUrl != "" {
		res.Image = ent.ImageUrl
	} else {
		res.Image = "default"
	}
}

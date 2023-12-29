package handler

import (
	"reflect"
	"wanderer/features/users"
)

type UserResponse struct {
	Id    uint   `json:"user_id,omitempty"`
	Name  string `json:"fullname,omitempty"`
	Phone string `json:"phone,omitempty"`
	Email string `json:"email,omitempty"`
	Image string `json:"image,omitempty"`
	Role  string `json:"role,omitempty"`

	TourCount   int               `json:"tour_count"`
	ReviewCount int               `json:"review_count"`
	Bookings    []BookingResponse `json:"bookings"`
}

func (res *UserResponse) FromEntity(ent users.User) {
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

	if ent.Role != "" {
		res.Role = ent.Role
	}

	if ent.Email != "" {
		res.Email = ent.Email
	}

	if ent.Phone != "" {
		res.Phone = ent.Phone
	}

	if ent.TourCount != 0 {
		res.TourCount = ent.TourCount
	}

	if ent.ReviewCount != 0 {
		res.ReviewCount = ent.ReviewCount
	}

	for _, booking := range ent.Bookings {
		var tmpBooking = new(BookingResponse)
		tmpBooking.FromEntity(booking)

		res.Bookings = append(res.Bookings, *tmpBooking)
	}
}

type BookingResponse struct {
	Code        int          `json:"booking_code,omitempty"`
	Status      string       `json:"status,omitempty"`
	DetailCount int          `json:"detail_count,omitempty"`
	Tour        TourResponse `json:"tour,omitempty"`
}

func (res *BookingResponse) FromEntity(ent users.Booking) {
	if ent.Code != 0 {
		res.Code = ent.Code
	}

	if ent.Status != "" {
		res.Status = ent.Status
	}

	if ent.DetailCount != 0 {
		res.DetailCount = ent.DetailCount
	}

	if !reflect.ValueOf(ent.Tour).IsZero() {
		res.Tour.FromEntity(ent.Tour)
	}
}

type TourResponse struct {
	Id    uint   `json:"tour_id,omitempty"`
	Title string `json:"title,omitempty"`
}

func (res *TourResponse) FromEntity(ent users.Tour) {
	if ent.Id != 0 {
		res.Id = ent.Id
	}

	if ent.Title != "" {
		res.Title = ent.Title
	}
}

type LoginResponse struct {
	Id    uint   `json:"user_id,omitempty"`
	Name  string `json:"fullname,omitempty"`
	Image string `json:"image,omitempty"`
	Role  string `json:"role,omitempty"`
	Token string `json:"token,omitempty"`
}

func (res *LoginResponse) FromEntity(ent users.User) {
	if ent.Id != 0 {
		res.Id = ent.Id
	}

	if ent.Name != "" {
		res.Name = ent.Name
	}

	if ent.ImageUrl != "" {
		res.Image = ent.ImageUrl
	} else {
		res.Image = "https://res.cloudinary.com/dhxzinjxp/image/upload/v1703490558/asset-default/user_d5pwxw.png"
	}

	if ent.Role != "" {
		res.Role = ent.Role
	}
}

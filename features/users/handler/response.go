package handler

import (
	"wanderer/features/users"
)

type UserResponse struct {
	Id    uint   `json:"user_id,omitempty"`
	Name  string `json:"fullname,omitempty"`
	Phone string `json:"phone,omitempty"`
	Email string `json:"email,omitempty"`
	Image string `json:"image,omitempty"`
	Role  string `json:"role,omitempty"`
	Token string `json:"token,omitempty"`
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
		res.Image = "default"
	}

	if ent.Role != "" {
		res.Role = ent.Role
	}
}

type UserByIdResponse struct {
	Id          uint   `json:"user_id,omitempty"`
	Name        string `json:"fullname,omitempty"`
	Email       string `json:"email,omitempty"`
	Image       string `json:"image,omitempty"`
	TourCount   int    `json:"tour_count,omitempty"`
	ReviewCount int    `json:"review_count,omitempty"`

	Bookings []Booking `json:"bookings,omitempty"`
	Reviews  []Review  `json:"-"`
}

func (res *UserByIdResponse) FromEntity(ent users.User) {
	if ent.Id != 0 {
		res.Id = ent.Id
	}

	if ent.Name != "" {
		res.Name = ent.Name
	}

	if ent.Email != "" {
		res.Email = ent.Email
	}

	if ent.ImageUrl != "" {
		res.Image = ent.ImageUrl
	} else {
		res.Image = "default"
	}

	for _, booking := range ent.Bookings {
		var tmpBooking = new(Booking)
		tmpBooking.FromEntity(booking)

		res.Bookings = append(res.Bookings, *tmpBooking)
	}

	res.TourCount = len(res.Bookings)

	for _, review := range ent.Reviews {
		var tmpReview = new(Review)
		tmpReview.FromEntity(review)

		res.Reviews = append(res.Reviews, *tmpReview)
	}

	res.ReviewCount = len(res.Reviews)

}

type Booking struct {
	Code        int    `json:"booking_code,omitempty"`
	DetailCount int    `json:"detail_count,omitempty"`
	Status      string `json:"status,omitempty"`

	Tour           Tour            `json:"tour,omitempty"`
	BookingDetails []BookingDetail `json:"-"`
}

func (res *Booking) FromEntity(ent users.Booking) {
	if ent.Code != 0 {
		res.Code = ent.Code
	}

	if ent.Status != "" {
		res.Status = ent.Status
	}

	res.Tour.FromEntity(ent.Tour)

	for _, bd := range ent.BookingDetails {
		var tmpBd = new(BookingDetail)
		tmpBd.FromEntity(bd)

		res.BookingDetails = append(res.BookingDetails, *tmpBd)
	}

	res.DetailCount = len(res.BookingDetails)
}

type Tour struct {
	Id    uint   `json:"tour_id,omitempty"`
	Title string `json:"title,omitempty"`
}

func (res *Tour) FromEntity(ent users.Tour) {
	if ent.Id != 0 {
		res.Id = ent.Id
	}

	if ent.Title != "" {
		res.Title = ent.Title
	}
}

type Review struct {
	Id     uint `json:"tour_id,omitempty"`
	UserId uint `json:"user_id,omitempty"`
}

func (res *Review) FromEntity(ent users.Review) {
	if ent.Id != 0 {
		res.Id = ent.Id
	}

	if ent.UserId != 0 {
		res.UserId = ent.UserId
	}
}

type BookingDetail struct {
	Id          uint `json:"tour_id,omitempty"`
	BookingCode int  `json:"booking_code,omitempty"`
}

func (res *BookingDetail) FromEntity(ent users.BookingDetail) {
	if ent.Id != 0 {
		res.Id = ent.Id
	}

	if ent.BookingCode != 0 {
		res.BookingCode = ent.BookingCode
	}
}

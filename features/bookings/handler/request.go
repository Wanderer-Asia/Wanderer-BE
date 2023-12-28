package handler

import (
	"time"
	"wanderer/features/bookings"
)

type BookingCreateRequest struct {
	TourId uint                         `json:"tour_id"`
	Detail []BookingDetailCreateRequest `json:"detail"`
	Bank   string                       `json:"payment_method"`
}

func (req *BookingCreateRequest) ToEntity(userId uint) bookings.Booking {
	var ent = new(bookings.Booking)

	if req.TourId != 0 {
		ent.Tour.Id = req.TourId
	}

	if userId != 0 {
		ent.User.Id = userId
	}

	for _, reqDetail := range req.Detail {
		ent.Detail = append(ent.Detail, reqDetail.ToEntity())
	}

	if req.Bank != "" {
		ent.Payment.Bank = req.Bank
	}

	return *ent
}

type BookingUpdateRequest struct {
	Bank   string `json:"payment_method"`
	Status string `json:"status"`
}

func (req *BookingUpdateRequest) ToEntity() bookings.Booking {
	var ent = new(bookings.Booking)

	if req.Bank != "" {
		ent.Payment.Bank = req.Bank
	}

	if req.Status != "" {
		ent.Status = req.Status
	}

	return *ent
}

type BookingDetailCreateRequest struct {
	DocumentNumber string    `json:"document_number"`
	Greeting       string    `json:"greeting"`
	Name           string    `json:"name"`
	Nationality    string    `json:"nationality"`
	DOB            time.Time `json:"dob"`
}

func (req *BookingDetailCreateRequest) ToEntity() bookings.Detail {
	var ent = new(bookings.Detail)

	if req.DocumentNumber != "" {
		ent.DocumentNumber = req.DocumentNumber
	}

	if req.Greeting != "" {
		ent.Greeting = req.Greeting
	}

	if req.Name != "" {
		ent.Name = req.Name
	}

	if req.Nationality != "" {
		ent.Nationality = req.Nationality
	}

	if !req.DOB.IsZero() {
		ent.DOB = req.DOB
	}

	return *ent
}

type PaymentNotificationRequest struct {
	Code   string `json:"order_id"`
	Status string `json:"transaction_status"`
}

package handler

import (
	"reflect"
	"time"
	"wanderer/features/bookings"
	tr "wanderer/features/tours/handler"
)

type BookingResponse struct {
	Code        int     `json:"booking_code,omitempty"`
	DetailCount int     `json:"detail_count,omitempty"`
	Status      string  `json:"status,omitempty"`
	Total       float64 `json:"total"`

	PaymentBank          string    `json:"payment_method,omitempty"`
	PaymentVirtualNumber string    `json:"virtual_number,omitempty"`
	PaymentBillKey       string    `json:"key_bill,omitempty"`
	PaymentBillCode      string    `json:"code_bill,omitempty"`
	PaymentExpiredAt     time.Time `json:"payment_expired"`

	Tour tr.TourResponse `json:"tour,omitempty"`
}

func (res *BookingResponse) FromEntity(ent bookings.Booking) {
	if ent.Code != 0 {
		res.Code = ent.Code
	}

	if ent.Total != 0 {
		res.Total = ent.Total
	}

	if ent.Status != "" {
		res.Status = ent.Status
	}

	if len(ent.Detail) != 0 {
		res.DetailCount = len(ent.Detail)
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
			res.PaymentExpiredAt = ent.Payment.ExpiredAt
		}
	}

	if !reflect.ValueOf(ent.Tour).IsZero() {
		var tour = new(tr.TourResponse)
		tour.FromEntity(ent.Tour)
		res.Tour = *tour
	}
}

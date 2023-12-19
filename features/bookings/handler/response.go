package handler

import (
	"reflect"
	"time"
	"wanderer/features/bookings"
)

type BookingResponse struct {
	Total float64 `json:"total"`

	PaymentVirtualNumber string    `json:"virtual_number,omitempty"`
	PaymentBillKey       string    `json:"key_bill,omitempty"`
	PaymentBillCode      string    `json:"code_bill,omitempty"`
	PaymentExpiredAt     time.Time `json:"payment_expired"`
}

func (res *BookingResponse) FromEntity(ent bookings.Booking) {
	if ent.Total != 0 {
		res.Total = ent.Total
	}

	if !reflect.ValueOf(ent.Payment).IsZero() {
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
}

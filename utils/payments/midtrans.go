package payments

import (
	"errors"
	"fmt"
	"time"
	"wanderer/config"
	"wanderer/features/bookings"

	mdt "github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type Midtrans interface {
	NewBookingPayment(data bookings.Booking) (*bookings.Payment, error)
	CancelBookingPayment(code int) error
}

func NewMidtrans(config config.Midtrans) Midtrans {
	var client coreapi.Client
	client.New(config.ApiKey, config.Env)

	return &midtrans{
		config: config,
		client: client,
	}
}

type midtrans struct {
	config config.Midtrans
	client coreapi.Client
}

func (pay *midtrans) NewBookingPayment(data bookings.Booking) (*bookings.Payment, error) {
	req := new(coreapi.ChargeReq)
	req.TransactionDetails = mdt.TransactionDetails{
		OrderID:  fmt.Sprintf("%d", data.Code),
		GrossAmt: int64(data.Total),
	}

	req.CustomerDetails = &mdt.CustomerDetails{
		FName: data.User.Name,
		Email: data.User.Email,
		Phone: data.User.Phone,
	}

	var reqItem []mdt.ItemDetails
	for _, detail := range data.Detail {
		reqItem = append(reqItem, mdt.ItemDetails{
			ID:    detail.DocumentNumber,
			Name:  detail.Greeting + " " + detail.Name,
			Price: int64(data.Total / float64(len(data.Detail))),
			Qty:   1,
		})
	}
	req.Items = &reqItem

	switch data.Payment.Bank {
	case "bca":
		req.PaymentType = coreapi.PaymentTypeBankTransfer
		req.BankTransfer = &coreapi.BankTransferDetails{
			Bank: mdt.BankBca,
		}
	case "bni":
		req.PaymentType = coreapi.PaymentTypeBankTransfer
		req.BankTransfer = &coreapi.BankTransferDetails{
			Bank: mdt.BankBni,
		}
	case "bri":
		req.PaymentType = coreapi.PaymentTypeBankTransfer
		req.BankTransfer = &coreapi.BankTransferDetails{
			Bank: mdt.BankBri,
		}
	case "permata":
		req.PaymentType = coreapi.PaymentTypeBankTransfer
		req.BankTransfer = &coreapi.BankTransferDetails{
			Bank: mdt.BankPermata,
		}
	case "mandiri":
		req.PaymentType = coreapi.PaymentTypeEChannel
		req.EChannel = &coreapi.EChannelDetail{
			BillInfo1: "Wanderer Booking",
			BillInfo2: fmt.Sprintf("%d person", len(data.Detail)),
			BillKey:   fmt.Sprintf("%d", data.Code),
		}
	default:
		return nil, errors.New("unsupported payment")
	}

	res, _ := pay.client.ChargeTransaction(req)
	if res.StatusCode != "201" {
		return nil, errors.New(res.StatusMessage)
	}

	if res.BillKey != "" {
		data.Payment.BillKey = res.BillKey
	}

	if res.BillerCode != "" {
		data.Payment.BillCode = res.BillerCode
	}

	if len(res.VaNumbers) == 1 {
		data.Payment.VirtualNumber = res.VaNumbers[0].VANumber
	}

	if res.PermataVaNumber != "" {
		data.Payment.VirtualNumber = res.PermataVaNumber
	}

	if res.PaymentType != "" {
		data.Payment.Method = res.PaymentType
	}

	if res.TransactionStatus != "" {
		data.Payment.Status = res.TransactionStatus
	}

	if expiredAt, err := time.Parse("2006-01-02 15:04:05", res.ExpiryTime); err != nil {
		return nil, err
	} else {
		data.Payment.ExpiredAt = expiredAt
	}

	data.Payment.BookingTotal = data.Total

	return &data.Payment, nil
}

func (pay *midtrans) CancelBookingPayment(code int) error {
	res, _ := pay.client.CancelTransaction(fmt.Sprintf("%d", code))
	if res.StatusCode != "200" && res.StatusCode != "412" {
		return errors.New(res.StatusMessage)
	}

	return nil
}

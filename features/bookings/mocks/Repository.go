// Code generated by mockery v2.37.1. DO NOT EDIT.

package mocks

import (
	context "context"
	bookings "wanderer/features/bookings"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, data
func (_m *Repository) Create(ctx context.Context, data bookings.Booking) (*bookings.Booking, error) {
	ret := _m.Called(ctx, data)

	var r0 *bookings.Booking
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, bookings.Booking) (*bookings.Booking, error)); ok {
		return rf(ctx, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, bookings.Booking) *bookings.Booking); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bookings.Booking)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, bookings.Booking) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx
func (_m *Repository) GetAll(ctx context.Context) ([]bookings.Booking, int, error) {
	ret := _m.Called(ctx)

	var r0 []bookings.Booking
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]bookings.Booking, int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []bookings.Booking); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bookings.Booking)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) int); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetDetail provides a mock function with given fields: ctx, code
func (_m *Repository) GetDetail(ctx context.Context, code int) (*bookings.Booking, error) {
	ret := _m.Called(ctx, code)

	var r0 *bookings.Booking
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*bookings.Booking, error)); ok {
		return rf(ctx, code)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *bookings.Booking); ok {
		r0 = rf(ctx, code)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bookings.Booking)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, code, data
func (_m *Repository) Update(ctx context.Context, code int, data bookings.Booking) (*bookings.Booking, error) {
	ret := _m.Called(ctx, code, data)

	var r0 *bookings.Booking
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, bookings.Booking) (*bookings.Booking, error)); ok {
		return rf(ctx, code, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, bookings.Booking) *bookings.Booking); ok {
		r0 = rf(ctx, code, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bookings.Booking)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, bookings.Booking) error); ok {
		r1 = rf(ctx, code, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

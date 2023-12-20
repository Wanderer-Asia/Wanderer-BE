// Code generated by mockery v2.37.1. DO NOT EDIT.

package mocks

import (
	context "context"
	filters "wanderer/helpers/filters"

	mock "github.com/stretchr/testify/mock"

	tours "wanderer/features/tours"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, data
func (_m *Service) Create(ctx context.Context, data tours.Tour) error {
	ret := _m.Called(ctx, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, tours.Tour) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx, flt
func (_m *Service) GetAll(ctx context.Context, flt filters.Filter) ([]tours.Tour, int, error) {
	ret := _m.Called(ctx, flt)

	var r0 []tours.Tour
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, filters.Filter) ([]tours.Tour, int, error)); ok {
		return rf(ctx, flt)
	}
	if rf, ok := ret.Get(0).(func(context.Context, filters.Filter) []tours.Tour); ok {
		r0 = rf(ctx, flt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]tours.Tour)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, filters.Filter) int); ok {
		r1 = rf(ctx, flt)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, filters.Filter) error); ok {
		r2 = rf(ctx, flt)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetDetail provides a mock function with given fields: ctx, id
func (_m *Service) GetDetail(ctx context.Context, id uint) (*tours.Tour, error) {
	ret := _m.Called(ctx, id)

	var r0 *tours.Tour
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) (*tours.Tour, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint) *tours.Tour); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*tours.Tour)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, id, data
func (_m *Service) Update(ctx context.Context, id uint, data tours.Tour) error {
	ret := _m.Called(ctx, id, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint, tours.Tour) error); ok {
		r0 = rf(ctx, id, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
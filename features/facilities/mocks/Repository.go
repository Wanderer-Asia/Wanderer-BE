// Code generated by mockery v2.37.1. DO NOT EDIT.

package mocks

import (
	facilities "wanderer/features/facilities"
	filters "wanderer/helpers/filters"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: newFacility
func (_m *Repository) Create(newFacility facilities.Facility) error {
	ret := _m.Called(newFacility)

	var r0 error
	if rf, ok := ret.Get(0).(func(facilities.Facility) error); ok {
		r0 = rf(newFacility)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *Repository) Delete(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: flt
func (_m *Repository) GetAll(flt filters.Filter) ([]facilities.Facility, error) {
	ret := _m.Called(flt)

	var r0 []facilities.Facility
	var r1 error
	if rf, ok := ret.Get(0).(func(filters.Filter) ([]facilities.Facility, error)); ok {
		return rf(flt)
	}
	if rf, ok := ret.Get(0).(func(filters.Filter) []facilities.Facility); ok {
		r0 = rf(flt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]facilities.Facility)
		}
	}

	if rf, ok := ret.Get(1).(func(filters.Filter) error); ok {
		r1 = rf(flt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, updateFacility
func (_m *Repository) Update(id uint, updateFacility facilities.Facility) error {
	ret := _m.Called(id, updateFacility)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, facilities.Facility) error); ok {
		r0 = rf(id, updateFacility)
	} else {
		r0 = ret.Error(0)
	}

	return r0
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

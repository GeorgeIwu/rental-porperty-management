// Code generated by mockery v2.13.1. DO NOT EDIT.

package apartment

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockIApartmentService is an autogenerated mock type for the IApartmentService type
type MockIApartmentService struct {
	mock.Mock
}

// Create provides a mock function with given fields: c, attributes
func (_m *MockIApartmentService) Create(c context.Context, attributes ApartmentRequest) (*Apartment, error) {
	ret := _m.Called(c, attributes)

	var r0 *Apartment
	if rf, ok := ret.Get(0).(func(context.Context, ApartmentRequest) *Apartment); ok {
		r0 = rf(c, attributes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Apartment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ApartmentRequest) error); ok {
		r1 = rf(c, attributes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: ctx, searchBy, sortBy, pageNumber, itemsPerPage
func (_m *MockIApartmentService) Fetch(ctx context.Context, searchBy string, sortBy string, pageNumber int, itemsPerPage int) ([]*Apartment, error) {
	ret := _m.Called(ctx, searchBy, sortBy, pageNumber, itemsPerPage)

	var r0 []*Apartment
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int, int) []*Apartment); ok {
		r0 = rf(ctx, searchBy, sortBy, pageNumber, itemsPerPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Apartment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, int, int) error); ok {
		r1 = rf(ctx, searchBy, sortBy, pageNumber, itemsPerPage)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchByID provides a mock function with given fields: c, id
func (_m *MockIApartmentService) FetchByID(c context.Context, id int) (*Apartment, error) {
	ret := _m.Called(c, id)

	var r0 *Apartment
	if rf, ok := ret.Get(0).(func(context.Context, int) *Apartment); ok {
		r0 = rf(c, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Apartment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(c, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: c, propertyID
func (_m *MockIApartmentService) Remove(c context.Context, propertyID int) error {
	ret := _m.Called(c, propertyID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(c, propertyID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: c, propertyID, r
func (_m *MockIApartmentService) Update(c context.Context, propertyID int, r ApartmentRequest) (*Apartment, error) {
	ret := _m.Called(c, propertyID, r)

	var r0 *Apartment
	if rf, ok := ret.Get(0).(func(context.Context, int, ApartmentRequest) *Apartment); ok {
		r0 = rf(c, propertyID, r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Apartment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, ApartmentRequest) error); ok {
		r1 = rf(c, propertyID, r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockIApartmentService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockIApartmentService creates a new instance of MockIApartmentService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockIApartmentService(t mockConstructorTestingTNewMockIApartmentService) *MockIApartmentService {
	mock := &MockIApartmentService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

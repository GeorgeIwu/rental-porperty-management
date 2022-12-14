// Code generated by mockery v2.13.1. DO NOT EDIT.

package property

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockIPropertyService is an autogenerated mock type for the IPropertyService type
type MockIPropertyService struct {
	mock.Mock
}

// Create provides a mock function with given fields: c, attributes
func (_m *MockIPropertyService) Create(c context.Context, attributes PropertyRequest) (*Property, error) {
	ret := _m.Called(c, attributes)

	var r0 *Property
	if rf, ok := ret.Get(0).(func(context.Context, PropertyRequest) *Property); ok {
		r0 = rf(c, attributes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Property)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, PropertyRequest) error); ok {
		r1 = rf(c, attributes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateManager provides a mock function with given fields: c, name
func (_m *MockIPropertyService) CreateManager(c context.Context, name string) (*Manager, error) {
	ret := _m.Called(c, name)

	var r0 *Manager
	if rf, ok := ret.Get(0).(func(context.Context, string) *Manager); ok {
		r0 = rf(c, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Manager)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: ctx, searchBy, sortBy, pageNumber, itemsPerPage
func (_m *MockIPropertyService) Fetch(ctx context.Context, searchBy string, sortBy string, pageNumber int, itemsPerPage int) ([]*Property, error) {
	ret := _m.Called(ctx, searchBy, sortBy, pageNumber, itemsPerPage)

	var r0 []*Property
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int, int) []*Property); ok {
		r0 = rf(ctx, searchBy, sortBy, pageNumber, itemsPerPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Property)
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
func (_m *MockIPropertyService) FetchByID(c context.Context, id int) (*Property, error) {
	ret := _m.Called(c, id)

	var r0 *Property
	if rf, ok := ret.Get(0).(func(context.Context, int) *Property); ok {
		r0 = rf(c, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Property)
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
func (_m *MockIPropertyService) Remove(c context.Context, propertyID int) error {
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
func (_m *MockIPropertyService) Update(c context.Context, propertyID int, r PropertyRequest) (*Property, error) {
	ret := _m.Called(c, propertyID, r)

	var r0 *Property
	if rf, ok := ret.Get(0).(func(context.Context, int, PropertyRequest) *Property); ok {
		r0 = rf(c, propertyID, r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Property)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, PropertyRequest) error); ok {
		r1 = rf(c, propertyID, r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockIPropertyService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockIPropertyService creates a new instance of MockIPropertyService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockIPropertyService(t mockConstructorTestingTNewMockIPropertyService) *MockIPropertyService {
	mock := &MockIPropertyService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

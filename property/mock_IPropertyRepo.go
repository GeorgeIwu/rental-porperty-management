// Code generated by mockery v2.13.1. DO NOT EDIT.

package property

import (
	context "context"
	ent "rental-porperty-management/ent"

	mock "github.com/stretchr/testify/mock"
)

// MockIPropertyRepo is an autogenerated mock type for the IPropertyRepo type
type MockIPropertyRepo struct {
	mock.Mock
}

// Create provides a mock function with given fields: c, attributes
func (_m *MockIPropertyRepo) Create(c context.Context, attributes PropertyRequest) (*ent.Property, error) {
	ret := _m.Called(c, attributes)

	var r0 *ent.Property
	if rf, ok := ret.Get(0).(func(context.Context, PropertyRequest) *ent.Property); ok {
		r0 = rf(c, attributes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.Property)
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
func (_m *MockIPropertyRepo) CreateManager(c context.Context, name string) (*ent.Manager, error) {
	ret := _m.Called(c, name)

	var r0 *ent.Manager
	if rf, ok := ret.Get(0).(func(context.Context, string) *ent.Manager); ok {
		r0 = rf(c, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.Manager)
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

// Delete provides a mock function with given fields: c, serviceID
func (_m *MockIPropertyRepo) Delete(c context.Context, serviceID int) error {
	ret := _m.Called(c, serviceID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(c, serviceID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, searchBy, sortBy, pageOffset, itemsPerPage
func (_m *MockIPropertyRepo) Get(ctx context.Context, searchBy string, sortBy string, pageOffset int, itemsPerPage int) ([]*ent.Property, error) {
	ret := _m.Called(ctx, searchBy, sortBy, pageOffset, itemsPerPage)

	var r0 []*ent.Property
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int, int) []*ent.Property); ok {
		r0 = rf(ctx, searchBy, sortBy, pageOffset, itemsPerPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ent.Property)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, int, int) error); ok {
		r1 = rf(ctx, searchBy, sortBy, pageOffset, itemsPerPage)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: c, id
func (_m *MockIPropertyRepo) GetByID(c context.Context, id int) (*ent.Property, error) {
	ret := _m.Called(c, id)

	var r0 *ent.Property
	if rf, ok := ret.Get(0).(func(context.Context, int) *ent.Property); ok {
		r0 = rf(c, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.Property)
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

// Update provides a mock function with given fields: c, serviceID, r
func (_m *MockIPropertyRepo) Update(c context.Context, serviceID int, r PropertyRequest) (*ent.Property, error) {
	ret := _m.Called(c, serviceID, r)

	var r0 *ent.Property
	if rf, ok := ret.Get(0).(func(context.Context, int, PropertyRequest) *ent.Property); ok {
		r0 = rf(c, serviceID, r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.Property)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, PropertyRequest) error); ok {
		r1 = rf(c, serviceID, r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockIPropertyRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockIPropertyRepo creates a new instance of MockIPropertyRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockIPropertyRepo(t mockConstructorTestingTNewMockIPropertyRepo) *MockIPropertyRepo {
	mock := &MockIPropertyRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

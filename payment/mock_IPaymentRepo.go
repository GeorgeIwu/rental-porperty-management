// Code generated by mockery v2.13.1. DO NOT EDIT.

package payment

import (
	context "context"
	ent "rental-porperty-management/ent"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// MockIPaymentRepo is an autogenerated mock type for the IPaymentRepo type
type MockIPaymentRepo struct {
	mock.Mock
}

// Create provides a mock function with given fields: c, rq
func (_m *MockIPaymentRepo) Create(c context.Context, rq PaymentRequest) (*ent.Payment, error) {
	ret := _m.Called(c, rq)

	var r0 *ent.Payment
	if rf, ok := ret.Get(0).(func(context.Context, PaymentRequest) *ent.Payment); ok {
		r0 = rf(c, rq)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.Payment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, PaymentRequest) error); ok {
		r1 = rf(c, rq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, fromDate, toDate, pageOffset, itemsPerPage
func (_m *MockIPaymentRepo) Get(ctx context.Context, fromDate time.Time, toDate time.Time, pageOffset int, itemsPerPage int) ([]*ent.Payment, error) {
	ret := _m.Called(ctx, fromDate, toDate, pageOffset, itemsPerPage)

	var r0 []*ent.Payment
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, time.Time, int, int) []*ent.Payment); ok {
		r0 = rf(ctx, fromDate, toDate, pageOffset, itemsPerPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ent.Payment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, time.Time, time.Time, int, int) error); ok {
		r1 = rf(ctx, fromDate, toDate, pageOffset, itemsPerPage)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: c, id
func (_m *MockIPaymentRepo) GetByID(c context.Context, id int) (*ent.Payment, error) {
	ret := _m.Called(c, id)

	var r0 *ent.Payment
	if rf, ok := ret.Get(0).(func(context.Context, int) *ent.Payment); ok {
		r0 = rf(c, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.Payment)
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

// GetByTenant provides a mock function with given fields: ctx, tenantID, fromDate, pageOffset, itemsPerPage
func (_m *MockIPaymentRepo) GetByTenant(ctx context.Context, tenantID int, fromDate time.Time, pageOffset int, itemsPerPage int) ([]*ent.Payment, error) {
	ret := _m.Called(ctx, tenantID, fromDate, pageOffset, itemsPerPage)

	var r0 []*ent.Payment
	if rf, ok := ret.Get(0).(func(context.Context, int, time.Time, int, int) []*ent.Payment); ok {
		r0 = rf(ctx, tenantID, fromDate, pageOffset, itemsPerPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ent.Payment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, time.Time, int, int) error); ok {
		r1 = rf(ctx, tenantID, fromDate, pageOffset, itemsPerPage)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockIPaymentRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockIPaymentRepo creates a new instance of MockIPaymentRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockIPaymentRepo(t mockConstructorTestingTNewMockIPaymentRepo) *MockIPaymentRepo {
	mock := &MockIPaymentRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

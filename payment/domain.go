package payment

import (
	"context"
	"rental-porperty-management/ent"
	"time"
)

type PaymentModel = ent.Payment
type TenantModel = ent.Tenant

// Payment ...
type Payment struct {
	ID        int       `json:"id"`
	Amount    int       `json:"amount" validate:"required"`
	Date      time.Time `json:"date" validate:"required"`
	TenantID  int       `json:"owner_id"`
	State     string    `json:"state"`
	CreatedAt time.Time `json:"created_at"`
}

// Create Payment Request...
type PaymentRequest struct {
	Amount   int       `json:"amount" validate:"required"`
	Date     time.Time `json:"date" validate:"required"`
	TenantID int       `json:"tenant_id" validate:"required"`
}

// IPaymentService ...
type IPaymentService interface {
	Fetch(ctx context.Context, fromDate time.Time, toDate time.Time, pageNumber int, itemsPerPage int) (res []*Payment, err error)
	FetchByID(c context.Context, id int) (res *Payment, err error)
	FetchByTenant(ctx context.Context, tenantID int, fromDate time.Time, pageNumber int, itemsPerPage int) (res []*Payment, err error)
	Create(c context.Context, attributes PaymentRequest) (res *Payment, err error)
}

// IPaymentRepo ...
type IPaymentRepo interface {
	Get(ctx context.Context, fromDate time.Time, toDate time.Time, pageOffset int, itemsPerPage int) (res []*PaymentModel, err error)
	GetByID(c context.Context, id int) (res *PaymentModel, err error)
	GetByTenant(ctx context.Context, tenantID int, fromDate time.Time, pageOffset int, itemsPerPage int) (res []*PaymentModel, err error)
	Create(c context.Context, rq PaymentRequest) (res *PaymentModel, err error)
}

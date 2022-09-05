package payment

import (
	"context"
	"fmt"
	"rental-porperty-management/ent"
	"rental-porperty-management/ent/payment"
	"rental-porperty-management/ent/tenant"
	"time"
)

type PaymentRepo struct {
	client         *ent.Client
	contextTimeout time.Duration
}

// NewPaymentRepo will create new an PaymentRepo
func NewPaymentRepo(c *ent.Client, timeout time.Duration) *PaymentRepo {
	return &PaymentRepo{
		client:         c,
		contextTimeout: timeout,
	}
}

func (r *PaymentRepo) Get(ctx context.Context, fromDate time.Time, toDate time.Time, pageOffset int, itemsPerPage int) (res []*PaymentModel, err error) {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	paymentEntities, err := r.client.Payment.Query().
		Where(
			payment.And(
				payment.CreatedAtGTE(fromDate),
				payment.CreatedAtLTE(toDate),
			),
		).
		Offset(pageOffset).
		Limit(itemsPerPage).
		Order(ent.Desc(payment.FieldCreatedAt)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return paymentEntities, nil
}

func (r *PaymentRepo) GetByID(c context.Context, id int) (res *PaymentModel, err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	paymentEntity, err := r.client.Payment.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return paymentEntity, nil
}

func (r *PaymentRepo) GetByTenant(ctx context.Context, tenantID int, fromDate time.Time, pageOffset int, itemsPerPage int) (res []*PaymentModel, err error) {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	paymentEntities, err := r.client.Payment.Query().
		Where(payment.OwnerIDEQ(tenantID)).
		Offset(pageOffset).
		Limit(itemsPerPage).
		Order(ent.Desc(payment.FieldCreatedAt)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return paymentEntities, nil
}

func (r *PaymentRepo) Create(c context.Context, rq PaymentRequest) (res *PaymentModel, err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	tx, err := r.client.Tx(ctx)
	defer cancel()

	tenantEntity, err := tx.Tenant.Query().
		Where(
			tenant.IDEQ(rq.TenantID),
		).
		WithApartment().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	// Update the Payment.
	paymentEntity, err := tx.Payment.
		Create().
		SetAmount(rq.Amount).
		SetTenantID(rq.TenantID).
		SetDate(rq.Date).
		SetApartmentID(tenantEntity.Edges.Apartment.ID).
		Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}
	fmt.Println("paymentEntity was created: ", paymentEntity)

	tx.Commit()

	return paymentEntity, nil
}

func (r *PaymentRepo) Update(c context.Context, paymentID int, rq PaymentRequest) (res *PaymentModel, err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	tx, err := r.client.Tx(ctx)
	defer cancel()

	existingPayment, err := tx.Payment.Get(ctx, paymentID)
	if existingPayment == nil {
		return nil, err
	}

	// Update the Payment.
	paymentEntity, err := tx.Payment.UpdateOneID(paymentID).
		SetState(payment.StateProcessed).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	tx.Commit()
	fmt.Println("version was created: ", paymentEntity)

	return paymentEntity, nil
}

func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}

package tenant

import (
	"context"
	"fmt"
	"rental-porperty-management/ent"
	"rental-porperty-management/ent/tenant"
	"strconv"
	"time"
)

type TenantRepo struct {
	client         *ent.Client
	contextTimeout time.Duration
}

// NewTenantRepo will create new an TenantRepo
func NewTenantRepo(c *ent.Client, timeout time.Duration) *TenantRepo {
	return &TenantRepo{
		client:         c,
		contextTimeout: timeout,
	}
}

func (r *TenantRepo) Get(ctx context.Context, searchBy string, sortBy string, pageOffset int, itemsPerPage int) (res []*TenantModel, err error) {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	ssn, err := strconv.Atoi(searchBy)
	if err != nil {
		return nil, err
	}

	tenantEntities, err := r.client.Tenant.Query().
		Where(
			tenant.Or(
				tenant.FirstNameContains(searchBy),
				tenant.LastNameContains(searchBy),
				tenant.SsnEQ(ssn),
			),
		).
		Offset(pageOffset).
		Limit(itemsPerPage).
		Order(ent.Desc(getSortType(sortBy))).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return tenantEntities, nil
}

func (r *TenantRepo) GetByID(c context.Context, id int) (res *TenantModel, err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	tenantEntity, err := r.client.Tenant.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return tenantEntity, nil
}

func (r *TenantRepo) Create(c context.Context, rq TenantRequest) (res *TenantModel, err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	tx, err := r.client.Tx(ctx)
	defer cancel()

	// Get Apartment.
	existingApartment, err := tx.Tenant.Get(ctx, rq.ApartmentID)
	if existingApartment == nil {
		return nil, err
	}

	// Create a new Tenant.
	tenantEntity, err := tx.Tenant.
		Create().
		SetApartmentID(existingApartment.ID).
		SetFirstName(rq.FirstName).
		SetLastName(rq.LastName).
		SetDob(rq.DoB).
		SetSsn(rq.SSN).
		SetLeaseStartAt(rq.LeaseStartAt).
		Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	tx.Commit()

	return tenantEntity, nil
}

func (r *TenantRepo) Update(c context.Context, tenantID int, rq TenantRequest) (res *TenantModel, err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	tx, err := r.client.Tx(ctx)
	defer cancel()

	existingApartment, err := tx.Apartment.Get(ctx, tenantID)
	if existingApartment == nil {
		return nil, err
	}

	// Update the Tenant.
	tenantEntity, err := tx.Tenant.
		UpdateOneID(tenantID).
		SetApartmentID(existingApartment.ID).
		SetFirstName(rq.FirstName).
		SetLastName(rq.LastName).
		SetDob(rq.DoB).
		SetSsn(rq.SSN).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	tx.Commit()
	fmt.Println("version was created: ", tenantEntity)

	return tenantEntity, nil
}

func (r *TenantRepo) Delete(c context.Context, tenantID int) (err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	existingTenant, err := r.client.Tenant.Get(ctx, tenantID)
	if existingTenant == nil {
		return err
	}

	// Delete the Tenant.
	newerr := r.client.Tenant.DeleteOneID(tenantID).Exec(ctx)
	if newerr != nil {
		return newerr
	}

	fmt.Println("version was deleted ")

	return nil
}

func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}

func getSortType(sort string) string {

	switch sort {
	case "first_name":
		return tenant.FieldFirstName
	case "last_name":
		return tenant.FieldLastName
	case "dob":
		return tenant.FieldDob
	default:
		return tenant.FieldCreatedAt
	}
}

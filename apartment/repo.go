package apartment

import (
	"context"
	"fmt"
	"rental-porperty-management/ent"
	"rental-porperty-management/ent/apartment"
	"time"
)

type ApartmentRepo struct {
	client         *ent.Client
	contextTimeout time.Duration
}

// NewApartmentRepo will create new an ApartmentRepo
func NewApartmentRepo(c *ent.Client, timeout time.Duration) *ApartmentRepo {
	return &ApartmentRepo{
		client:         c,
		contextTimeout: timeout,
	}
}

func (r *ApartmentRepo) Get(ctx context.Context, searchBy string, sortBy string, pageOffset int, itemsPerPage int) (res []*ApartmentModel, err error) {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	apartmentEntities, err := r.client.Apartment.Query().
		Where(
			apartment.UnitNumberContains(searchBy),
		).
		Offset(pageOffset).
		Limit(itemsPerPage).
		Order(ent.Desc(getSortType(sortBy))).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return apartmentEntities, nil
}

func (r *ApartmentRepo) GetByID(c context.Context, id int) (res *ApartmentModel, err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	apartmentEntity, err := r.client.Apartment.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return apartmentEntity, nil
}

func (r *ApartmentRepo) Create(c context.Context, rq ApartmentRequest) (res *ApartmentModel, err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	tx, err := r.client.Tx(ctx)
	defer cancel()

	// Get Apartment.
	existingProperty, err := tx.Property.Get(ctx, rq.PropertyID)
	if existingProperty == nil {
		return nil, err
	}

	// Create a new Apartment.
	apartmentEntity, err := tx.Apartment.
		Create().
		SetPropertyID(existingProperty.ID).
		SetUnitNumber(rq.UnitNumber).
		SetCharge(rq.Charge).
		Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	tx.Commit()

	return apartmentEntity, nil
}

func (r *ApartmentRepo) Update(c context.Context, apartmentID int, rq ApartmentRequest) (res *ApartmentModel, err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	tx, err := r.client.Tx(ctx)
	defer cancel()

	existingApartment, err := tx.Apartment.Get(ctx, apartmentID)
	if existingApartment == nil {
		return nil, err
	}

	// Update the Apartment.
	apartmentEntity, err := tx.Apartment.
		UpdateOneID(apartmentID).
		SetUnitNumber(rq.UnitNumber).
		SetCharge(rq.Charge).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	tx.Commit()
	fmt.Println("version was created: ", apartmentEntity)

	return apartmentEntity, nil
}

func (r *ApartmentRepo) Delete(c context.Context, apartmentID int) (err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	existingApartment, err := r.client.Apartment.Get(ctx, apartmentID)
	if existingApartment == nil {
		return err
	}

	// Delete the Apartment.
	newerr := r.client.Apartment.DeleteOneID(apartmentID).Exec(ctx)
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
	case "unit_number":
		return apartment.FieldUnitNumber
	case "charge":
		return apartment.FieldCharge
	default:
		return apartment.FieldCreatedAt
	}
}

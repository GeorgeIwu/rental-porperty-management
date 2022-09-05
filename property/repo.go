package property

import (
	"context"
	"errors"
	"fmt"
	"rental-porperty-management/ent"
	"rental-porperty-management/ent/predicate"
	"rental-porperty-management/ent/property"
	"time"
)

type PropertyRepo struct {
	client         *ent.Client
	contextTimeout time.Duration
}

// NewPropertyRepo will create new an PropertyRepo
func NewPropertyRepo(c *ent.Client, timeout time.Duration) *PropertyRepo {
	return &PropertyRepo{
		client:         c,
		contextTimeout: timeout,
	}
}

func (r *PropertyRepo) Get(ctx context.Context, searchBy string, sortBy string, pageOffset int, itemsPerPage int) (res []*PropertyModel, err error) {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	propertyEntities, err := r.client.Property.Query().
		Where(
			property.Or(
				property.NameContains(searchBy),
				property.Address(searchBy),
			),
		).
		WithManager().
		Offset(pageOffset).
		Limit(itemsPerPage).
		Order(ent.Desc(getSortType(sortBy))).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return propertyEntities, nil
}

func (r *PropertyRepo) GetByID(c context.Context, id int) (res *PropertyModel, err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	propertyEntity, err := r.client.Property.Query().
		Where(
			property.IDEQ(id),
		).
		WithManager().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return propertyEntity, nil
}

func (r *PropertyRepo) Create(c context.Context, rq PropertyRequest) (res *PropertyModel, err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	tx, err := r.client.Tx(ctx)
	defer cancel()

	// Create a manager.
	manager, err := tx.Manager.
		Create().
		SetName(rq.Username).
		SetUpdatedAt(time.Now()).
		SetCreatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}
	fmt.Println("manager was created: ", manager)

	// Create a new Property.
	propertyEntity, err := tx.Property.
		Create().
		SetName(rq.Name).
		SetAddress(rq.Address).
		SetUnitsCount(rq.NumeberOfUnits).
		SetManager(manager).
		SetUpdatedAt(time.Now()).
		SetCreatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	tx.Commit()

	return propertyEntity, nil
}

func (r *PropertyRepo) Update(c context.Context, propertyID int, rq PropertyRequest) (res *PropertyModel, err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	tx, err := r.client.Tx(ctx)
	defer cancel()

	existingProperty, err := tx.Property.Get(ctx, propertyID)
	if existingProperty == nil {
		return nil, err
	}

	// Update the Property.
	propertyEntity, err := tx.Property.UpdateOneID(propertyID).
		SetName(rq.Name).
		SetAddress(rq.Address).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	tx.Commit()
	fmt.Println("version was created: ", propertyEntity)

	return propertyEntity, nil
}

func (r *PropertyRepo) Delete(c context.Context, propertyID int) (err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	existingProperty, err := r.client.Property.Get(ctx, propertyID)
	if existingProperty == nil {
		return err
	}

	// Delete the Property.
	newerr := r.client.Property.DeleteOneID(propertyID).Exec(ctx)
	if newerr != nil {
		return newerr
	}

	fmt.Println("version was deleted ")

	return nil
}

func (r *PropertyRepo) CreateManager(c context.Context, name string) (res *ManagerModel, err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	tx, err := r.client.Tx(ctx)
	defer cancel()

	existingProperty, err := tx.Manager.Query().
		Where(
			predicate.Manager(property.NameEQ(name)),
		).
		All(ctx)
	if existingProperty != nil {
		return nil, errors.New("name already exists")
	}

	// Create a manager.
	managerEntity, err := tx.Manager.Create().SetName(name).Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}
	fmt.Println("manager was created: ", managerEntity)

	tx.Commit()
	fmt.Println("manager was created: ", managerEntity)

	return managerEntity, nil
}

func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}

func getSortType(sort string) string {

	switch sort {
	case "name":
		return property.FieldName
	case "address":
		return property.FieldAddress
	default:
		return property.FieldCreatedAt
	}
}

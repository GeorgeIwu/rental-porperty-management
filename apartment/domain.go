package apartment

import (
	"context"
	"rental-porperty-management/ent"
	"time"
)

type ApartmentModel = ent.Apartment

// Apartment ...
type Apartment struct {
	ID         int       `json:"id"`
	UnitNumber string    `json:"unit_number" validate:"required"`
	Charge     int       `json:"charge" validate:"required"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// Create Apartment Request...
type ApartmentRequest struct {
	PropertyID int    `json:"property_id" validate:"required"`
	UnitNumber string `json:"unit_number" validate:"required"`
	Charge     int    `json:"charge" validate:"required"`
}

// IApartmentService ...
type IApartmentService interface {
	Fetch(ctx context.Context, searchBy string, sortBy string, pageNumber int, itemsPerPage int) (res []*Apartment, err error)
	FetchByID(c context.Context, id int) (res *Apartment, err error)
	Create(c context.Context, attributes ApartmentRequest) (res *Apartment, err error)
	Update(c context.Context, propertyID int, r ApartmentRequest) (res *Apartment, err error)
	Remove(c context.Context, propertyID int) (err error)
}

// IApartmentRepo ...
type IApartmentRepo interface {
	Get(ctx context.Context, searchBy string, sortBy string, pageOffset int, itemsPerPage int) (res []*ApartmentModel, err error)
	GetByID(c context.Context, id int) (res *ApartmentModel, err error)
	Create(c context.Context, attributes ApartmentRequest) (res *ApartmentModel, err error)
	Update(c context.Context, serviceID int, r ApartmentRequest) (res *ApartmentModel, err error)
	Delete(c context.Context, serviceID int) (err error)
}

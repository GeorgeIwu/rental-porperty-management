package tenant

import (
	"context"
	"rental-porperty-management/ent"
	"time"
)

type TenantModel = ent.Tenant

// Tenant ...
type Tenant struct {
	ID            int       `json:"id"`
	FirstName     string    `json:"first_name" validate:"required"`
	LastName      string    `json:"last_name"`
	DoB           time.Time `json:"date_of_birth"`
	LeaseStartAt  time.Time `json:"lease_start_date"`
	LeaseEndAt    time.Time `json:"lease_end_date`
	IsLeaseHolder bool      `json:"is_lease_holder"`
	State         string    `json:"state"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// Create Tenant Request...
type TenantRequest struct {
	ApartmentID   int           `json:"apartment_id" validate:"required"`
	FirstName     string        `json:"first_name" validate:"required"`
	LastName      string        `json:"last_name"`
	DoB           time.Time     `json:"date_of_birth"`
	SSN           int           `json:"ssn" validate:"required"`
	LeaseStartAt  time.Time     `json:"lease_start_date"`
	LeaseEndAt    time.Time     `json:"lease_end_date`
	IsLeaseHolder bool          `json:"is_lease_holder"`
	Duration      time.Duration `json:"duration"`
	IsHolder      bool          `json:"is_holder"`
}

// ITenantService ...
type ITenantService interface {
	Fetch(ctx context.Context, searchBy string, sortBy string, pageNumber int, itemsPerPage int) (res []*Tenant, err error)
	FetchByID(c context.Context, id int) (res *Tenant, err error)
	Create(c context.Context, attributes TenantRequest) (res *Tenant, err error)
	Update(c context.Context, propertyID int, r TenantRequest) (res *Tenant, err error)
	Remove(c context.Context, propertyID int) (err error)
}

// ITenantRepo ...
type ITenantRepo interface {
	Get(ctx context.Context, searchBy string, sortBy string, pageOffset int, itemsPerPage int) (res []*TenantModel, err error)
	GetByID(c context.Context, id int) (res *TenantModel, err error)
	Create(c context.Context, attributes TenantRequest) (res *TenantModel, err error)
	Update(c context.Context, serviceID int, r TenantRequest) (res *TenantModel, err error)
	Delete(c context.Context, serviceID int) (err error)
}

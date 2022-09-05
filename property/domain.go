package property

import (
	"context"
	"rental-porperty-management/ent"
	"time"
)

type PropertyModel = ent.Property
type ManagerModel = ent.Manager

// Manager ...
type Manager struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" validate:"required"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// Property ...
type Property struct {
	ID             int       `json:"id"`
	Name           string    `json:"name" validate:"required"`
	Address        string    `json:"address" validate:"required"`
	NumeberOfUnits int       `json:"number_of_units" validate:"required"`
	ManagerID      int       `json:"manager_id" validate:"required"`
	UpdatedAt      time.Time `json:"updated_at"`
	CreatedAt      time.Time `json:"created_at"`
}

// Create Property Request...
type ManagerRequest struct {
	Name string `json:"name" validate:"required"`
}

// Create Property Request...
type PropertyRequest struct {
	Username       string `json:"username"`
	Name           string `json:"name" validate:"required"`
	Address        string `json:"address" validate:"required"`
	NumeberOfUnits int    `json:"number_of_units" validate:"required"`
}

// IPropertyService ...
type IPropertyService interface {
	Fetch(ctx context.Context, searchBy string, sortBy string, pageNumber int, itemsPerPage int) (res []*Property, err error)
	FetchByID(c context.Context, id int) (res *Property, err error)
	Create(c context.Context, attributes PropertyRequest) (res *Property, err error)
	Update(c context.Context, propertyID int, r PropertyRequest) (res *Property, err error)
	Remove(c context.Context, propertyID int) (err error)
	CreateManager(c context.Context, name string) (res *Manager, err error)
}

// IPropertyRepo ...
type IPropertyRepo interface {
	Get(ctx context.Context, searchBy string, sortBy string, pageOffset int, itemsPerPage int) (res []*PropertyModel, err error)
	GetByID(c context.Context, id int) (res *PropertyModel, err error)
	Create(c context.Context, attributes PropertyRequest) (res *PropertyModel, err error)
	Update(c context.Context, serviceID int, r PropertyRequest) (res *PropertyModel, err error)
	Delete(c context.Context, serviceID int) (err error)
	CreateManager(c context.Context, name string) (res *ManagerModel, err error)
}

package tenant

import (
	"context"
	"errors"
	"fmt"
)

type TenantService struct {
	repo ITenantRepo
}

var (
	NO_RECORD     = errors.New(`Record not found`)
	BAD_REQUEST   = errors.New(`Some parameters are not valid`)
	REQUEST_ERROR = errors.New(`Fialed to process request`)
)

// NewTenantService will create new an TenantService
func NewTenantService(r ITenantRepo) *TenantService {
	return &TenantService{
		repo: r,
	}
}

func (s *TenantService) Fetch(ctx context.Context, searchBy string, sortBy string, pageNumber int, itemsPerPage int) (res []*Tenant, err error) {
	pageOffset := (pageNumber - 1) * itemsPerPage
	tenantEntities, err := s.repo.Get(ctx, searchBy, sortBy, pageOffset, itemsPerPage)
	if err != nil {
		return nil, NO_RECORD
	}

	tenants := []*Tenant{}
	for _, tenantEntity := range tenantEntities {
		newtenant, _ := mapTenant(tenantEntity)
		tenants = append(tenants, newtenant)
	}

	return tenants, nil
}

func (s *TenantService) FetchByID(ctx context.Context, id int) (res *Tenant, err error) {

	tenantEntity, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, NO_RECORD
	}

	tenant, err := mapTenant(tenantEntity)
	if err != nil {
		return nil, REQUEST_ERROR
	}

	return tenant, nil
}

func (s *TenantService) Create(c context.Context, r TenantRequest) (res *Tenant, err error) {
	tenantEntity, err := s.repo.Create(c, r)
	if err != nil {
		return nil, REQUEST_ERROR
	}

	tenant, err := mapTenant(tenantEntity)
	if err != nil {
		return nil, REQUEST_ERROR
	}
	fmt.Println("tenant was created: ", tenant)

	return tenant, nil
}

func (s *TenantService) Update(c context.Context, tenantID int, r TenantRequest) (res *Tenant, err error) {

	// Update the Tenant.
	tenantEntity, err := s.repo.Update(c, tenantID, r)
	if err != nil {
		return nil, REQUEST_ERROR
	}

	tenant, err := mapTenant(tenantEntity)
	if err != nil {
		return nil, REQUEST_ERROR
	}
	fmt.Println("tenant was update: ", tenant)

	return tenant, nil
}

func (s *TenantService) Remove(c context.Context, tenantID int) (err error) {

	newerr := s.repo.Delete(c, tenantID)
	if newerr != nil {
		return BAD_REQUEST
	}

	fmt.Println("tenant was deleted ")

	return nil
}

func mapTenant(data *TenantModel) (*Tenant, error) {

	tenant := &Tenant{
		ID:            data.ID,
		FirstName:     data.FirstName,
		LastName:      data.LastName,
		DoB:           data.Dob,
		IsLeaseHolder: data.IsLeaseHolder,
		State:         string(data.State),
		CreatedAt:     data.CreatedAt,
		UpdatedAt:     data.UpdatedAt,
	}

	return tenant, nil
}

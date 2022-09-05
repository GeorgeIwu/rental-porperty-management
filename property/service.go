package property

import (
	"context"
	"errors"
	"fmt"
)

type PropertyService struct {
	repo IPropertyRepo
}

var (
	NO_RECORD     = errors.New(`Record not found`)
	BAD_REQUEST   = errors.New(`Some parameters are not valid`)
	REQUEST_ERROR = errors.New(`Fialed to process request`)
)

// NewPropertyService will create new an PropertyService
func NewPropertyService(r IPropertyRepo) *PropertyService {
	return &PropertyService{
		repo: r,
	}
}

func (s *PropertyService) Fetch(ctx context.Context, searchBy string, sortBy string, pageNumber int, itemsPerPage int) (res []*Property, err error) {
	pageOffset := (pageNumber - 1) * itemsPerPage
	propertyEntities, err := s.repo.Get(ctx, searchBy, sortBy, pageOffset, itemsPerPage)
	if err != nil {
		return nil, NO_RECORD
	}

	properties := []*Property{}
	for _, propertyEntity := range propertyEntities {
		newproperty, _ := mapProperty(*propertyEntity)
		properties = append(properties, newproperty)
	}

	return properties, nil
}

func (s *PropertyService) FetchByID(ctx context.Context, id int) (res *Property, err error) {

	propertyEntity, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, NO_RECORD
	}

	property, err := mapProperty(*propertyEntity)
	if err != nil {
		return nil, REQUEST_ERROR
	}

	return property, nil
}

func (s *PropertyService) Create(c context.Context, r PropertyRequest) (res *Property, err error) {
	propertyEntity, err := s.repo.Create(c, r)
	if err != nil {
		return nil, REQUEST_ERROR
	}

	property, err := mapProperty(*propertyEntity)
	if err != nil {
		return nil, REQUEST_ERROR
	}
	fmt.Println("property was created: ", property)

	return property, nil
}

func (s *PropertyService) Update(c context.Context, propertyID int, r PropertyRequest) (res *Property, err error) {

	// Update the Property.
	propertyEntity, err := s.repo.Update(c, propertyID, r)
	if err != nil {
		return nil, REQUEST_ERROR
	}

	property, err := mapProperty(*propertyEntity)
	if err != nil {
		return nil, REQUEST_ERROR
	}
	fmt.Println("property was update: ", property)

	return property, nil
}

func (s *PropertyService) Remove(c context.Context, propertyID int) (err error) {

	newerr := s.repo.Delete(c, propertyID)
	if newerr != nil {
		return BAD_REQUEST
	}

	fmt.Println("property was deleted ")

	return nil
}

func (s *PropertyService) CreateManager(c context.Context, managerName string) (res *Manager, err error) {

	managerEntity, err := s.repo.CreateManager(c, managerName)
	if err != nil {
		return nil, REQUEST_ERROR
	}

	manager, err := mapManager(managerEntity)
	if err != nil {
		return nil, REQUEST_ERROR
	}
	fmt.Println("manager was created: ", manager)

	return manager, nil
}

func mapProperty(data PropertyModel) (*Property, error) {

	messageID := 0 //strconv.ParseInt(data.Edges.Manager.ID)

	property := &Property{
		ID:             data.ID,
		Name:           data.Name,
		Address:        data.Address,
		NumeberOfUnits: data.UnitsCount,
		ManagerID:      messageID,
		CreatedAt:      data.CreatedAt,
		UpdatedAt:      data.UpdatedAt,
	}

	return property, nil
}

func mapManager(data *ManagerModel) (*Manager, error) {

	manager := &Manager{
		ID:        data.ID,
		Name:      data.Name,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return manager, nil
}

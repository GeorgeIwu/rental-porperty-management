package apartment

import (
	"context"
	"errors"
	"fmt"
)

type ApartmentService struct {
	repo IApartmentRepo
}

var (
	NO_RECORD     = errors.New(`Record not found`)
	BAD_REQUEST   = errors.New(`Some parameters are not valid`)
	REQUEST_ERROR = errors.New(`Fialed to process request`)
)

// NewApartmentService will create new an ApartmentService
func NewApartmentService(r IApartmentRepo) *ApartmentService {
	return &ApartmentService{
		repo: r,
	}
}

func (s *ApartmentService) Fetch(ctx context.Context, searchBy string, sortBy string, pageNumber int, itemsPerPage int) (res []*Apartment, err error) {
	pageOffset := (pageNumber - 1) * itemsPerPage
	apartmentEntities, err := s.repo.Get(ctx, searchBy, sortBy, pageOffset, itemsPerPage)
	if err != nil {
		return nil, NO_RECORD
	}

	apartments := []*Apartment{}
	for _, apartmentEntity := range apartmentEntities {
		newapartment, _ := mapApartment(apartmentEntity)
		apartments = append(apartments, newapartment)
	}

	return apartments, nil
}

func (s *ApartmentService) FetchByID(ctx context.Context, id int) (res *Apartment, err error) {

	apartmentEntity, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, NO_RECORD
	}

	apartment, err := mapApartment(apartmentEntity)
	if err != nil {
		return nil, err
	}

	return apartment, nil
}

func (s *ApartmentService) Create(c context.Context, r ApartmentRequest) (res *Apartment, err error) {
	apartmentEntity, err := s.repo.Create(c, r)
	if err != nil {
		return nil, REQUEST_ERROR
	}

	apartment, err := mapApartment(apartmentEntity)
	if err != nil {
		return nil, REQUEST_ERROR
	}
	fmt.Println("apartment was created: ", apartment)

	return apartment, nil
}

func (s *ApartmentService) Update(c context.Context, apartmentID int, r ApartmentRequest) (res *Apartment, err error) {

	// Update the Apartment.
	apartmentEntity, err := s.repo.Update(c, apartmentID, r)
	if err != nil {
		return nil, REQUEST_ERROR
	}

	apartment, err := mapApartment(apartmentEntity)
	if err != nil {
		return nil, REQUEST_ERROR
	}
	fmt.Println("apartment was update: ", apartment)

	return apartment, nil
}

func (s *ApartmentService) Remove(c context.Context, apartmentID int) (err error) {

	newerr := s.repo.Delete(c, apartmentID)
	if newerr != nil {
		return BAD_REQUEST
	}

	fmt.Println("apartment was deleted ")

	return nil
}

func mapApartment(data *ApartmentModel) (*Apartment, error) {

	apartment := &Apartment{
		ID:         data.ID,
		UnitNumber: data.UnitNumber,
		Charge:     data.Charge,
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
	}

	return apartment, nil
}

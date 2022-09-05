package payment

import (
	"context"
	"errors"
	"fmt"
	"rental-porperty-management/utils"
	"time"
)

type PaymentService struct {
	repo IPaymentRepo
}

var (
	NO_RECORD     = errors.New(`Record not found`)
	BAD_REQUEST   = errors.New(`Some parameters are not valid`)
	REQUEST_ERROR = errors.New(`Failed to process request`)
)

// NewPaymentService will create new an PaymentService
func NewPaymentService(r IPaymentRepo) *PaymentService {
	return &PaymentService{
		repo: r,
	}
}

func (s *PaymentService) Fetch(ctx context.Context, fromDate time.Time, toDate time.Time, pageNumber int, itemsPerPage int) (res []*Payment, err error) {
	pageOffset := (pageNumber - 1) * itemsPerPage
	paymentEntities, err := s.repo.Get(ctx, fromDate, toDate, pageOffset, itemsPerPage)
	if err != nil {
		return nil, NO_RECORD
	}

	payments := []*Payment{}
	for _, paymentEntity := range paymentEntities {
		newpayment, _ := mapPayment(paymentEntity)
		payments = append(payments, newpayment)
	}

	return payments, nil
}

func (s *PaymentService) FetchByID(ctx context.Context, id int) (res *Payment, err error) {

	paymentEntity, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, NO_RECORD
	}

	payment, err := mapPayment(paymentEntity)
	if err != nil {
		return nil, REQUEST_ERROR
	}

	return payment, nil
}

func (s *PaymentService) FetchByTenant(ctx context.Context, tenantID int, fromDate time.Time, pageNumber int, itemsPerPage int) (res []*Payment, err error) {
	pageOffset := (pageNumber - 1) * itemsPerPage
	paymentEntities, err := s.repo.GetByTenant(ctx, tenantID, fromDate, pageOffset, itemsPerPage)
	if err != nil {
		return nil, NO_RECORD
	}

	payments := []*Payment{}
	for _, paymentEntity := range paymentEntities {
		if len(payments) > 0 {
			prevPayment := payments[len(payments)-1]

			if utils.AreSameMonth(&prevPayment.Date, &paymentEntity.Date) == true {
				prevPayment.Amount += paymentEntity.Amount
				payments[len(payments)-1] = prevPayment
				continue
			}
		}

		newpayment, _ := mapPayment(paymentEntity)
		payments = append(payments, newpayment)
	}

	return payments, nil
}

func (s *PaymentService) Create(c context.Context, r PaymentRequest) (res *Payment, err error) {

	paymentEntity, err := s.repo.Create(c, r)
	if err != nil {
		return nil, REQUEST_ERROR
	}

	payment, err := mapPayment(paymentEntity)
	if err != nil {
		return nil, REQUEST_ERROR
	}
	fmt.Println("payment was created: ", payment)

	return payment, nil
}

func mapPayment(data *PaymentModel) (*Payment, error) {

	payment := &Payment{
		ID:        data.ID,
		Amount:    data.Amount,
		Date:      data.Date,
		TenantID:  data.OwnerID,
		State:     string(data.State),
		CreatedAt: data.CreatedAt,
	}

	return payment, nil
}

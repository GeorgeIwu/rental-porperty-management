package payment

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"rental-porperty-management/ent"
	"rental-porperty-management/ent/payment"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gopkg.in/go-playground/assert.v1"
)

type RequestData struct {
	searchBy     string
	sortBy       string
	pageNumber   int
	itemsPerPage int
}

//Unit Tests
type UnitTestSuite struct {
	suite.Suite

	repo   *MockIPaymentRepo
	server *httptest.Server
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, &UnitTestSuite{})
}

func (uts *UnitTestSuite) SetupTest() {
	ec := echo.New()
	repoMock := MockIPaymentRepo{}
	service := NewPaymentService(&repoMock)
	controller := NewPaymentController(ec, service)

	ec.GET("/payments", controller.Fetch)
	ec.POST("/payments", controller.Create)
	ec.GET("/tenant/:id/payments", controller.FetchByTenant)

	server := httptest.NewServer(ec)

	uts.repo = &repoMock
	uts.server = server
}

func (uts *UnitTestSuite) TearDownSuite() {
	defer uts.server.Close()
}

func (uts *UnitTestSuite) TestModelGet() {
	payments := getResponseData()
	uts.repo.On("Get", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(payments, nil)

	resp, err := http.Get(uts.server.URL + "/payments")
	if err != nil {
		uts.T().Fatal(err)
	}
	defer resp.Body.Close()

	var result []Payment
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		uts.T().Fatal(err)
	}

	assert.Equal(uts.T(), http.StatusOK, resp.StatusCode)
	assert.Equal(uts.T(), payments[0].Amount, result[0].Amount)
}

func (uts *UnitTestSuite) TestModelGetByTenant() {
	payments := getResponseData()

	secondPayment := *payments[0]
	secondPayment.Amount = 20
	newDate, err := time.Parse("2006-01-02", "2022-02-14")
	if err != nil {
		uts.T().Fatal(err)
	}
	secondPayment.Date = newDate

	thirdPayment := *payments[0]
	thirdPayment.Amount = 30
	thirdPayment.Date = newDate.AddDate(0, 0, 10)
	payments = append(payments, &secondPayment, &thirdPayment)

	uts.repo.On("GetByTenant", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(payments, nil)

	resp, err := http.Get(uts.server.URL + "/tenant/1/payments")
	if err != nil {
		uts.T().Fatal(err)
	}
	defer resp.Body.Close()

	var result []Payment
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		uts.T().Fatal(err)
	}

	assert.Equal(uts.T(), http.StatusOK, resp.StatusCode)
	assert.Equal(uts.T(), len(payments)-1, len(result))
	assert.Equal(uts.T(), secondPayment.Amount+thirdPayment.Amount, result[1].Amount)
}

func (uts *UnitTestSuite) TestModelCreateSuccess() {
	payments := getResponseData()
	requesData := getResquestData()
	uts.repo.On("Create", mock.Anything, mock.Anything).Return(payments[0], nil)

	requestBody, err := json.Marshal(requesData)
	if err != nil {
		uts.T().Fatal(err)
	}

	resp, err := http.Post(uts.server.URL+"/payments", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		uts.T().Fatal(err)
	}
	defer resp.Body.Close()

	var result Payment
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		uts.T().Fatal(err)
	}

	assert.Equal(uts.T(), http.StatusCreated, resp.StatusCode)
	assert.Equal(uts.T(), payments[0].ID, result.ID)
}

func (uts *UnitTestSuite) TestModelCreateFail() {
	payments := getResponseData()
	requestData := getResquestData()
	uts.repo.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(payments[0], nil)

	requestData.Amount = 0
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		uts.T().Fatal(err)
	}

	resp, err := http.Post(uts.server.URL+"/payments", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		uts.T().Fatal(err)
	}
	defer resp.Body.Close()

	assert.Equal(uts.T(), http.StatusBadRequest, resp.StatusCode)
}

func getResponseData() []*ent.Payment {
	newPaymentData := ent.Payment{
		ID:        3,
		Amount:    10,
		Date:      time.Now(),
		OwnerID:   1,
		State:     payment.StateUnprocessed,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	payments := []*ent.Payment{}
	payments = append(payments, &newPaymentData)
	return payments
}

func getResquestData() PaymentRequest {
	requestBody := PaymentRequest{
		Amount:   25,
		Date:     time.Now(),
		TenantID: 35,
	}

	return requestBody
}

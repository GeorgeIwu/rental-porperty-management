package apartment

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"rental-porperty-management/ent"
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

	repo   *MockIApartmentRepo
	server *httptest.Server
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, &UnitTestSuite{})
}

func (uts *UnitTestSuite) SetupTest() {
	ec := echo.New()
	repoMock := MockIApartmentRepo{}
	service := NewApartmentService(&repoMock)
	controller := NewApartmentController(ec, service)

	ec.GET("/apartments", controller.Fetch)
	ec.POST("/apartments", controller.Create)

	server := httptest.NewServer(ec)

	uts.repo = &repoMock
	uts.server = server
}

func (uts *UnitTestSuite) TearDownSuite() {
	defer uts.server.Close()
}

func (uts *UnitTestSuite) TestModelGet() {
	apartments := getResponseData()
	uts.repo.On("Get", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(apartments, nil)

	resp, err := http.Get(uts.server.URL + "/apartments")
	if err != nil {
		uts.T().Fatal(err)
	}
	defer resp.Body.Close()

	var result []Apartment
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		uts.T().Fatal(err)
	}

	assert.Equal(uts.T(), http.StatusOK, resp.StatusCode)
	assert.Equal(uts.T(), apartments[0].UnitNumber, result[0].UnitNumber)
}

func (uts *UnitTestSuite) TestModelCreateSuccess() {
	apartments := getResponseData()
	requesData := getResquestData()
	uts.repo.On("Create", mock.Anything, mock.Anything).Return(apartments[0], nil)

	requestBody, err := json.Marshal(requesData)
	if err != nil {
		uts.T().Fatal(err)
	}

	resp, err := http.Post(uts.server.URL+"/apartments", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		uts.T().Fatal(err)
	}
	defer resp.Body.Close()

	var result Apartment
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		uts.T().Fatal(err)
	}

	assert.Equal(uts.T(), http.StatusCreated, resp.StatusCode)
	assert.Equal(uts.T(), apartments[0].ID, result.ID)
}

func (uts *UnitTestSuite) TestModelCreateFail() {
	apartments := getResponseData()
	requestData := getResquestData()
	uts.repo.On("Create", mock.Anything, mock.Anything).Return(apartments[0], nil)

	requestData.UnitNumber = ""
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		uts.T().Fatal(err)
	}

	resp, err := http.Post(uts.server.URL+"/apartments", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		uts.T().Fatal(err)
	}
	defer resp.Body.Close()

	assert.Equal(uts.T(), http.StatusBadRequest, resp.StatusCode)
}

func getResponseData() []*ent.Apartment {
	newApartmentData := ent.Apartment{
		ID:         35,
		UnitNumber: "44",
		Charge:     200,
		UpdatedAt:  time.Now(),
		CreatedAt:  time.Now(),
	}

	apartments := []*ent.Apartment{}
	apartments = append(apartments, &newApartmentData)
	return apartments
}

func getResquestData() ApartmentRequest {
	requestBody := ApartmentRequest{
		PropertyID: 1,
		UnitNumber: "4",
		Charge:     200,
	}

	return requestBody
}

package tenant

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"rental-porperty-management/ent"
	"rental-porperty-management/ent/tenant"
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

	repo   *MockITenantRepo
	server *httptest.Server
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, &UnitTestSuite{})
}

func (uts *UnitTestSuite) SetupTest() {
	ec := echo.New()
	repoMock := MockITenantRepo{}
	service := NewTenantService(&repoMock)
	controller := NewTenantController(ec, service)

	ec.GET("/tenants", controller.Fetch)
	ec.POST("/tenants", controller.Create)

	server := httptest.NewServer(ec)

	uts.repo = &repoMock
	uts.server = server
}

func (uts *UnitTestSuite) TearDownSuite() {
	defer uts.server.Close()
}

func (uts *UnitTestSuite) TestModelGet() {
	tenants := getResponseData()
	uts.repo.On("Get", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tenants, nil)

	resp, err := http.Get(uts.server.URL + "/tenants")
	if err != nil {
		uts.T().Fatal(err)
	}
	defer resp.Body.Close()

	var result []Tenant
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		uts.T().Fatal(err)
	}

	assert.Equal(uts.T(), http.StatusOK, resp.StatusCode)
	assert.Equal(uts.T(), tenants[0].FirstName, result[0].FirstName)
}

func (uts *UnitTestSuite) TestModelCreateSuccess() {
	tenants := getResponseData()
	requesData := getResquestData()
	uts.repo.On("Create", mock.Anything, mock.Anything).Return(tenants[0], nil)

	requestBody, err := json.Marshal(requesData)
	if err != nil {
		uts.T().Fatal(err)
	}

	resp, err := http.Post(uts.server.URL+"/tenants", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		uts.T().Fatal(err)
	}
	defer resp.Body.Close()

	var result Tenant
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		uts.T().Fatal(err)
	}

	assert.Equal(uts.T(), http.StatusCreated, resp.StatusCode)
	assert.Equal(uts.T(), tenants[0].ID, result.ID)
}

func (uts *UnitTestSuite) TestModelCreateFail() {
	tenants := getResponseData()
	requestData := getResquestData()
	uts.repo.On("Create", mock.Anything, mock.Anything).Return(tenants[0], nil)

	requestData.FirstName = ""
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		uts.T().Fatal(err)
	}

	resp, err := http.Post(uts.server.URL+"/tenants", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		uts.T().Fatal(err)
	}
	defer resp.Body.Close()

	assert.Equal(uts.T(), http.StatusBadRequest, resp.StatusCode)
}

func getResponseData() []*ent.Tenant {
	newTenantData := ent.Tenant{
		ID:            35,
		FirstName:     "Don",
		LastName:      "Capo",
		Dob:           time.Now(),
		Ssn:           2323,
		IsLeaseHolder: true,
		State:         tenant.StateActive,
		UpdatedAt:     time.Now(),
		CreatedAt:     time.Now(),
	}

	tenants := []*ent.Tenant{}
	tenants = append(tenants, &newTenantData)
	return tenants
}

func getResquestData() TenantRequest {
	requestBody := TenantRequest{
		ApartmentID:  1,
		FirstName:    "Doe",
		LastName:     "Lom",
		DoB:          time.Now(),
		SSN:          2323,
		LeaseStartAt: time.Now(),
		Duration:     time.Duration(120394034),
		IsHolder:     true,
	}

	return requestBody
}

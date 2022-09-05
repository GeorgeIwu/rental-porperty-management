package property

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

	repo   *MockIPropertyRepo
	server *httptest.Server
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, &UnitTestSuite{})
}

func (uts *UnitTestSuite) SetupTest() {
	ec := echo.New()
	repoMock := MockIPropertyRepo{}
	service := NewPropertyService(&repoMock)
	controller := NewPropertyController(ec, service)

	ec.GET("/properties", controller.Fetch)
	ec.POST("/properties", controller.Create)

	server := httptest.NewServer(ec)

	uts.repo = &repoMock
	uts.server = server
}

func (uts *UnitTestSuite) TearDownSuite() {
	defer uts.server.Close()
}

func (uts *UnitTestSuite) TestModelGet() {
	properties := getResponseData()
	uts.repo.On("Get", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(properties, nil)

	resp, err := http.Get(uts.server.URL + "/properties")
	if err != nil {
		uts.T().Fatal(err)
	}
	defer resp.Body.Close()

	var result []Property
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		uts.T().Fatal(err)
	}

	assert.Equal(uts.T(), http.StatusOK, resp.StatusCode)
	assert.Equal(uts.T(), properties[0].Name, result[0].Name)
}

func (uts *UnitTestSuite) TestModelCreateSuccess() {
	properties := getResponseData()
	requesData := getResquestData()
	uts.repo.On("Create", mock.Anything, mock.Anything).Return(properties[0], nil)

	requestBody, err := json.Marshal(requesData)
	if err != nil {
		uts.T().Fatal(err)
	}

	resp, err := http.Post(uts.server.URL+"/properties", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		uts.T().Fatal(err)
	}
	defer resp.Body.Close()

	var result Property
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		uts.T().Fatal(err)
	}

	assert.Equal(uts.T(), http.StatusCreated, resp.StatusCode)
	assert.Equal(uts.T(), properties[0].Name, result.Name)
}

func (uts *UnitTestSuite) TestModelCreateFail() {
	properties := getResponseData()
	requestData := getResquestData()
	uts.repo.On("Create", mock.Anything, mock.Anything).Return(properties[0], nil)

	requestData.Name = ""
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		uts.T().Fatal(err)
	}

	resp, err := http.Post(uts.server.URL+"/properties", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		uts.T().Fatal(err)
	}
	defer resp.Body.Close()

	assert.Equal(uts.T(), http.StatusBadRequest, resp.StatusCode)
}

func getResponseData() []*ent.Property {
	newPropertyData := ent.Property{
		ID:         1,
		Name:       "Obi",
		Address:    "34",
		UnitsCount: 4,
		CreatedAt:  time.Now(),
	}

	properties := []*ent.Property{}
	properties = append(properties, &newPropertyData)
	return properties
}

func getResquestData() PropertyRequest {
	requestBody := PropertyRequest{
		Username:       "John",
		Name:           "KLS",
		Address:        "ksdol",
		NumeberOfUnits: 2,
	}

	return requestBody
}

package property

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"rental-porperty-management/ent"
	"rental-porperty-management/ent/enttest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func SetupManagerTest(uts *UnitTestSuite) func(uts *UnitTestSuite) {
	t := uts.T()
	ec := echo.New()
	duration := time.Duration(120) * time.Second

	client := getClient(t)
	repo := NewPropertyRepo(client, duration)
	service := NewPropertyService(repo)
	controller := NewPropertyController(ec, service)

	ec.GET("/manager", controller.CreateManager)

	server := httptest.NewServer(ec)
	backup_server := uts.server
	uts.server = server

	return func(uts *UnitTestSuite) {
		uts.server = backup_server
		server.Close()
		client.Close()
	}
}

func (uts *UnitTestSuite) TestManagerCreateSuccess() {
	teardownTest := SetupManagerTest(uts)
	defer teardownTest(uts)
	requesData := getResquestData()

	requestBody, err := json.Marshal(requesData)
	if err != nil {
		uts.T().Fatal(err)
	}

	resp, err := http.Post(uts.server.URL+"/manager", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		uts.T().Fatal(err)
	}
	defer resp.Body.Close()

	var result Manager
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		uts.T().Fatal(err)
	}
}

func getManagerResquestData() ManagerRequest {
	requestBody := ManagerRequest{
		Name: "KLS",
	}

	return requestBody
}

func getClient(t *testing.T) *ent.Client {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	return client
}

// func getClient(t *testing.T) *ent.Client {
// 	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// Run the auto migration tool.
// 	if err := client.Schema.Create(context.Background()); err != nil {
// 		t.Fatal(err)
// 	}

// 	return client
// }

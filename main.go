package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	// _ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"

	"rental-porperty-management/apartment"
	"rental-porperty-management/ent"
	"rental-porperty-management/payment"
	"rental-porperty-management/property"
	"rental-porperty-management/tenant"
	"rental-porperty-management/utils"
)

//move to environment variable
const (
	dbHost     = "localhost"
	dbUser     = "kong"
	dbPassword = "password"
	dbType     = "mysql"
	dbProtocol = "tcp"
	dbName     = "catalog"
)

func main() {
	// sqlInfo := fmt.Sprintf("%s:%s@%s(%s)/%s?parseTime=True", dbUser, dbPassword, dbProtocol, dbHost, dbName)
	// client, err := ent.Open(dbType, sqlInfo)
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sql: %v", err)
	}
	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	ec := echo.New()
	middleware := utils.InitMiddleware()
	ec.Use(middleware.Auth)
	timeoutContext := time.Duration(60) * time.Second

	paymentRepo := payment.NewPaymentRepo(client, timeoutContext)
	paymentService := payment.NewPaymentService(paymentRepo)
	payment.NewPaymentController(ec, paymentService)

	tenantRepo := tenant.NewTenantRepo(client, timeoutContext)
	tenantService := tenant.NewTenantService(tenantRepo)
	tenant.NewTenantController(ec, tenantService)

	apartmentRepo := apartment.NewApartmentRepo(client, timeoutContext)
	apartmentService := apartment.NewApartmentService(apartmentRepo)
	apartment.NewApartmentController(ec, apartmentService)

	propertyRepo := property.NewPropertyRepo(client, timeoutContext)
	propertyService := property.NewPropertyService(propertyRepo)
	property.NewPropertyController(ec, propertyService)

	fmt.Printf("Running app on port 8000")
	log.Fatal(http.ListenAndServe(":8000", ec))
}

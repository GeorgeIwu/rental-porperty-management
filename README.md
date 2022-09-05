# Rental-porperty-management

## Description
This is an example of implementation of Rental Property Management. The entity relationship are shown below.
<img width="801" alt="Screenshot 2022-09-05 at 6 12 00 AM" src="https://user-images.githubusercontent.com/28821928/188426913-7fd8b6c1-e312-4214-a7bb-9f931500e251.png">



## APIs
  ### Property
  - GET "/properties"
  - GET "/properties/:id"
  - POST "/properties"
  - PUT "/properties/:id"
  - POST "/manager"
  ### Apartment
  - GET "/apartments"
  - GET "/apartments/:id"
  - POST "/apartments"
  - PUT "/apartments/:id"
  ### Tenant
  - GET "/tenants"
  - GET "/tenants/:id"
  - POST "/tenants"
  - PUT "/tenants/:id"
  ### Payment
  - GET "/payments"
  - GET "/payments/:id"
  - GET "/tenant/:id/payments"


## Setup
  * clone repository
  * change into project directory
  * Run `go generate ./ent` to generate the models
  * Run `go mod tidy` to tidy
  * RRun application with `go run main.go`
  * Run tests with`go test ./...` or specific package tests with `go test ./property -v`
  <!-- go mod init rental-porperty-management -->
  <!-- go get ./...    -->

## Prerequisite

The following needs to be installed to run application
 * install sqlite3 on system, it is needed by dependency `github.com/mattn/go-sqlite3` (OR OPTIONAL mysql `go get github.com/go-sql-driver/mysql`)

## Considerations

Following the Clean Architecture by Uncle Bob
 * Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
 * Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
 * Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
 * Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
 * Independent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.

More at https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html

This project has  4 Domain layer :
 * Domain Layer
 * Repository(DB) Layer
 * Usecase(Service) Layer
 * Handler(Controller) Layer


## Test Considerations
  * Installed Mockery with `brew install mockery`
  * Ran `mockery --all --inpackage` to generate mocks for repo and service
  * To run all tests with `go test ./...`
  * To run specific test with package name, example `go test ./property -v`

## Scaling Considerations
  * Payment GET API has `from_date` and `to_date` which can query with a historical period
  * Payment entity has state `processed` and `unprocessed` which maybe updated to indicate payment processing (if applicable)
  * Payment entity has a mapping to Apartment entity with `owner_id` and Apartment entity has a mapping Property entity
    so one can query ```Get Property -> Apartments -> Payments```
  * To extend service to get history payments for property, we can build another Payment GET API `/property/:id/payments`
    which takes in query params `from_date` and `to_date`. In the repo this would make join query to fetch the Apartments
    for the property `id`, and get the Payments for each Apartment with the dates sent.

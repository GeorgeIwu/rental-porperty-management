package apartment

import (
	"net/http"
	"rental-porperty-management/utils"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// ApartmentHandler  represent the httphandler for Apartment
type ApartmentController struct {
	service IApartmentService
}

// NewApartmentHandler will initialize the Apartments/ resources endpoint
func NewApartmentController(ec *echo.Echo, s IApartmentService) *ApartmentController {
	apartmentController := &ApartmentController{
		service: s,
	}

	ec.GET("/apartments", apartmentController.Fetch)
	ec.GET("/apartments/:id", apartmentController.FetchByID)
	ec.POST("/apartments", apartmentController.Create)
	ec.PUT("/apartments/:id", apartmentController.Update)
	return apartmentController
}

// FetchApartment will fetch the Apartment based on given params
func (c *ApartmentController) Fetch(ec echo.Context) error {
	ITEMS_PER_PAGE := 12

	ctx := ec.Request().Context()
	searchBy := ec.QueryParam("search")
	sortBy := ec.QueryParam("sort")
	page := ec.QueryParam("page")

	pageNumber, _ := strconv.Atoi(page)
	if pageNumber == 0 {
		pageNumber = 1
	}

	err := validator.New().Var(pageNumber, "gt=0")
	if err != nil {
		return ec.JSON(http.StatusNotFound, "record not found")
	}

	properties, err := c.service.Fetch(ctx, searchBy, sortBy, pageNumber, ITEMS_PER_PAGE)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return ec.JSON(http.StatusOK, properties)
}

// FetchByID will get Apartment by given id
func (c *ApartmentController) FetchByID(ec echo.Context) error {
	ctx := ec.Request().Context()
	apartmentID, err := strconv.Atoi(ec.Param("id"))
	if err != nil {
		return ec.JSON(http.StatusNotFound, "record not found")
	}

	id := int(apartmentID)

	apartment, err := c.service.FetchByID(ctx, id)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return ec.JSON(http.StatusOK, apartment)
}

// Create will create the Apartment by given request body
func (c *ApartmentController) Create(ec echo.Context) (err error) {
	ctx := ec.Request().Context()
	var request ApartmentRequest
	if err := ec.Bind(&request); err != nil {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}

	if ok, err := utils.IsRequestValid(request); !ok {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}

	newapartment, err := c.service.Create(ctx, request)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return ec.JSON(http.StatusCreated, newapartment)
}

// Update will update the Apartment by given request body
func (c *ApartmentController) Update(ec echo.Context) (err error) {
	ctx := ec.Request().Context()
	apartmentID, err := strconv.Atoi(ec.Param("id"))

	var request ApartmentRequest
	if err := ec.Bind(&request); err != nil {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}

	if ok, err := utils.IsRequestValid(request); !ok {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}

	newapartment, err := c.service.Update(ctx, apartmentID, request)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return ec.JSON(http.StatusCreated, newapartment)
}

// Delete will update the Apartment by given request body
func (c *ApartmentController) Delete(ec echo.Context) (err error) {
	ctx := ec.Request().Context()
	apartmentID, err := strconv.Atoi(ec.Param("id"))

	if err != nil {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}

	err = c.service.Remove(ctx, apartmentID)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return ec.JSON(http.StatusCreated, nil)
}

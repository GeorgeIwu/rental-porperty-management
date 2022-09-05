package property

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

// PropertyHandler  represent the httphandler for Property
type PropertyController struct {
	service IPropertyService
}

// NewPropertyHandler will initialize the Propertys/ resources endpoint
func NewPropertyController(ec *echo.Echo, s IPropertyService) *PropertyController {
	propertyController := &PropertyController{
		service: s,
	}

	ec.GET("/properties", propertyController.Fetch)
	ec.GET("/properties/:id", propertyController.FetchByID)
	ec.POST("/properties", propertyController.Create)
	ec.PUT("/properties/:id", propertyController.Update)
	ec.POST("/manager", propertyController.CreateManager)
	return propertyController
}

// FetchProperty will fetch the Property based on given params
func (c *PropertyController) Fetch(ec echo.Context) error {
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

// FetchByID will get Property by given id
func (c *PropertyController) FetchByID(ec echo.Context) error {
	ctx := ec.Request().Context()
	propertyID, err := strconv.Atoi(ec.Param("id"))
	if err != nil {
		return ec.JSON(http.StatusNotFound, "record not found")
	}

	id := int(propertyID)

	property, err := c.service.FetchByID(ctx, id)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return ec.JSON(http.StatusOK, property)
}

// Create will create the Property by given request body
func (c *PropertyController) Create(ec echo.Context) (err error) {
	ctx := ec.Request().Context()
	var request PropertyRequest
	if err := ec.Bind(&request); err != nil {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}

	if ok, err := utils.IsRequestValid(request); !ok {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}

	newproperty, err := c.service.Create(ctx, request)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return ec.JSON(http.StatusCreated, newproperty)
}

// Update will update the Property by given request body
func (c *PropertyController) Update(ec echo.Context) (err error) {
	ctx := ec.Request().Context()
	propertyID, err := strconv.Atoi(ec.Param("id"))

	var request PropertyRequest
	if err := ec.Bind(&request); err != nil {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}

	if ok, err := utils.IsRequestValid(request); !ok {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}

	newproperty, err := c.service.Update(ctx, propertyID, request)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return ec.JSON(http.StatusCreated, newproperty)
}

// Delete will update the Property by given request body
func (c *PropertyController) Delete(ec echo.Context) (err error) {
	ctx := ec.Request().Context()
	propertyID, err := strconv.Atoi(ec.Param("id"))

	if err != nil {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}

	err = c.service.Remove(ctx, propertyID)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return ec.JSON(http.StatusCreated, nil)
}

// CreateVersion will create the Version by given request body
func (c *PropertyController) CreateManager(ec echo.Context) (err error) {
	ctx := ec.Request().Context()
	var request ManagerRequest
	if err := ec.Bind(&request); err != nil {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}

	if ok, err := utils.IsRequestValid(request); !ok {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}

	newproperty, err := c.service.CreateManager(ctx, request.Name)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return ec.JSON(http.StatusCreated, newproperty)
}

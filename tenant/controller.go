package tenant

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

// TenantHandler  represent the httphandler for Tenant
type TenantController struct {
	service ITenantService
}

// NewTenantHandler will initialize the Tenants/ resources endpoint
func NewTenantController(ec *echo.Echo, s ITenantService) *TenantController {
	tenantController := &TenantController{
		service: s,
	}

	ec.GET("/tenants", tenantController.Fetch)
	ec.GET("/tenants/:id", tenantController.FetchByID)
	ec.POST("/tenants", tenantController.Create)
	ec.PUT("/tenants/:id", tenantController.Update)
	return tenantController
}

// FetchTenant will fetch the Tenant based on given params
func (c *TenantController) Fetch(ec echo.Context) error {
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

// FetchByID will get Tenant by given id
func (c *TenantController) FetchByID(ec echo.Context) error {
	ctx := ec.Request().Context()
	tenantID, err := strconv.Atoi(ec.Param("id"))
	if err != nil {
		return ec.JSON(http.StatusNotFound, "record not found")
	}

	id := int(tenantID)

	tenant, err := c.service.FetchByID(ctx, id)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return ec.JSON(http.StatusOK, tenant)
}

// Create will create the Tenant by given request body
func (c *TenantController) Create(ec echo.Context) (err error) {
	ctx := ec.Request().Context()
	var request TenantRequest
	if err := ec.Bind(&request); err != nil {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}

	if ok, err := utils.IsRequestValid(request); !ok {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}

	newtenant, err := c.service.Create(ctx, request)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return ec.JSON(http.StatusCreated, newtenant)
}

// Update will update the Tenant by given request body
func (c *TenantController) Update(ec echo.Context) (err error) {
	ctx := ec.Request().Context()
	tenantID, err := strconv.Atoi(ec.Param("id"))

	var request TenantRequest
	if err := ec.Bind(&request); err != nil {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}

	if ok, err := utils.IsRequestValid(request); !ok {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}

	newtenant, err := c.service.Update(ctx, tenantID, request)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return ec.JSON(http.StatusCreated, newtenant)
}

// Delete will update the Tenant by given request body
func (c *TenantController) Delete(ec echo.Context) (err error) {
	ctx := ec.Request().Context()
	tenantID, err := strconv.Atoi(ec.Param("id"))

	if err != nil {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}

	err = c.service.Remove(ctx, tenantID)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return ec.JSON(http.StatusCreated, nil)
}

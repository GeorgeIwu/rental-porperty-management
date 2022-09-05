package payment

import (
	"errors"
	"net/http"
	"rental-porperty-management/utils"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// PaymentHandler  represent the httphandler for Payment
type PaymentController struct {
	service IPaymentService
}

// NewPaymentHandler will initialize the Payments/ resources endpoint
func NewPaymentController(ec *echo.Echo, s IPaymentService) *PaymentController {
	paymentController := &PaymentController{
		service: s,
	}

	ec.GET("/payments", paymentController.Fetch)
	ec.GET("/payments/:id", paymentController.FetchByID)
	ec.GET("/tenant/:id/payments", paymentController.FetchByTenant)
	ec.POST("/payments", paymentController.Create)
	return paymentController
}

// FetchPayment will fetch the Payment based on given params
func (c *PaymentController) Fetch(ec echo.Context) error {
	FORMAT := "2006-01-02"
	ITEMS_PER_PAGE := 12

	ctx := ec.Request().Context()
	page := ec.QueryParam("page")
	pageNumber, _ := strconv.Atoi(page)
	if pageNumber == 0 {
		pageNumber = 1
	}
	from := ec.QueryParam("from_date")
	if from == "" {
		from = "1960-02-02"
	}
	to := ec.QueryParam("to_date")
	if to == "" {
		to = time.Now().Format(FORMAT)
	}

	fromDate, err := time.Parse(FORMAT, from)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}

	toDate, err := time.Parse(FORMAT, to)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}

	err = validator.New().Var(pageNumber, "gt=0")
	if err != nil {
		return ec.JSON(http.StatusNotFound, err)
	}

	properties, err := c.service.Fetch(ctx, fromDate, toDate, pageNumber, ITEMS_PER_PAGE)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return ec.JSON(http.StatusOK, properties)
}

// FetchByID will get Payment by given id
func (c *PaymentController) FetchByID(ec echo.Context) error {
	ctx := ec.Request().Context()
	paymentID, err := strconv.Atoi(ec.Param("id"))
	if err != nil {
		return ec.JSON(http.StatusNotFound, "record not found")
	}

	id := int(paymentID)

	payment, err := c.service.FetchByID(ctx, id)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return ec.JSON(http.StatusOK, payment)
}

// FetchPayment will fetch the Payment based on given params
func (c *PaymentController) FetchByTenant(ec echo.Context) error {
	FORMAT := "2006-01-02"
	ITEMS_PER_PAGE := 12

	ctx := ec.Request().Context()
	id, err := strconv.Atoi(ec.Param("id"))
	if err != nil {
		return ec.JSON(http.StatusNotFound, "record not found")
	}

	tenantID := int(id)

	page := ec.QueryParam("page")
	pageNumber, _ := strconv.Atoi(page)
	if pageNumber == 0 {
		pageNumber = 1
	}
	from := ec.QueryParam("from_date")
	if from == "" {
		from = "1960-02-02"
	}

	fromDate, err := time.Parse(FORMAT, from)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}

	err = validator.New().Var(pageNumber, "gt=0")
	if err != nil {
		return ec.JSON(http.StatusNotFound, err)
	}

	properties, err := c.service.FetchByTenant(ctx, tenantID, fromDate, pageNumber, ITEMS_PER_PAGE)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return ec.JSON(http.StatusOK, properties)
}

// Create will create the Payment by given request body
func (c *PaymentController) Create(ec echo.Context) (err error) {
	ctx := ec.Request().Context()
	var request PaymentRequest
	if err := ec.Bind(&request); err != nil {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}

	if ok, err := utils.IsRequestValid(request); !ok {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}

	err = utils.ValidatePaymentDate(&request.Date)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, errors.New("Invalid Date"))
	}

	newpayment, err := c.service.Create(ctx, request)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return ec.JSON(http.StatusCreated, newpayment)
}

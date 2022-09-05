package utils

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
)

// Help validate request structure
func IsRequestValid(m interface{}) (bool, error) {
	// if err := c.Bind(&request); err != nil {
	// 	return c.JSON(http.StatusUnprocessableEntity, err.Error())
	// }

	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ValidatePaymentDate(date *time.Time) (err error) {
	today := time.Now()

	if today.Before(*date) {
		return errors.New("Invalid")
	}

	return nil
}

func AreSameMonth(date1 *time.Time, date2 *time.Time) (isValid bool) {
	_, month1, _ := date1.Date()
	_, month2, _ := date2.Date()

	if month1 == month2 {
		return true
	}

	return false
}

package router

import (
	"github.com/go-playground/validator/v10"
)

// customValidator implements echo.Validator
type customValidator struct {
	validator *validator.Validate
}

func (cv customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

package router

import (
	"strconv"

	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func newBinder() *binder {
	v := validator.New()
	v.RegisterValidation("date", validateDate)

	return &binder{
		binder:    &echo.DefaultBinder{},
		validator: v,
	}
}

// binder implements echo.Binder
type binder struct {
	binder    *echo.DefaultBinder
	validator *validator.Validate
}

func (cb binder) Bind(i interface{}, c echo.Context) error {
	if err := defaults.Set(i); err != nil {
		return err
	}
	if err := cb.binder.Bind(i, c); err != nil {
		return err
	}
	return cb.validator.Struct(i)
}

// yyyy-MM-dd
func validateDate(fl validator.FieldLevel) bool {
	f := fl.Field().String()

	if len(f) != 10 {
		return false
	}
	year, err := strconv.Atoi(f[:4])
	if err != nil {
		return false
	}
	if year < 1900 || year > 3000 { // TODO: this will break after 976 years
		return false
	}
	month, err := strconv.Atoi(f[5:7])
	if err != nil {
		return false
	}
	if month < 1 || month > 12 {
		return false
	}
	day, err := strconv.Atoi(f[8:10])
	if err != nil {
		return false
	}
	if day < 1 || day > 31 {
		return false
	}
	return true
}

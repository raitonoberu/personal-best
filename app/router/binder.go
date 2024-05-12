package router

import (
	"time"

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
	_, err := time.Parse("2006-01-02", fl.Field().String())
	return err == nil
}

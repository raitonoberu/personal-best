package router

import (
	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func newBinder() *binder {
	return &binder{
		binder:    &echo.DefaultBinder{},
		validator: validator.New(),
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

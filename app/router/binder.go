package router

import (
	"github.com/creasty/defaults"
	"github.com/labstack/echo/v4"
)

// customBinder implements echo.Binder
type customBinder struct {
	binder *echo.DefaultBinder
}

func (cb customBinder) Bind(i interface{}, c echo.Context) error {
	if err := defaults.Set(i); err != nil {
		return err
	}
	return cb.binder.Bind(i, c)
}

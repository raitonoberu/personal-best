package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/raitonoberu/personal-best/app/middleware"
)

// New creates a router with all the middlewares registered.
func New() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HTTPErrorHandler = errorHandler

	e.Validator = customValidator{validator.New()}
	e.Binder = customBinder{&echo.DefaultBinder{}}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Auth)
	return e
}

package router

import (
	"github.com/labstack/echo/v4"
	"github.com/raitonoberu/personal-best/app/middleware"
	"github.com/swaggo/echo-swagger"

	_ "github.com/raitonoberu/personal-best/docs"
)

// New creates a router with all the middlewares registered.
func New() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HTTPErrorHandler = errorHandler

	e.Binder = newBinder()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Auth)

	e.GET("/api/docs/*", echoSwagger.WrapHandler)
	return e
}

package middleware

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Logger sets the logger middleware.
func Logger() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		LogError:  true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				log.Printf("[%d] %s %s\n", v.Status, c.Request().Method, v.URI)
			} else {
				log.Printf("[%d] %s %s error=%s\n", v.Status, c.Request().Method, v.URI, v.Error)
			}
			return nil
		},
	})
}

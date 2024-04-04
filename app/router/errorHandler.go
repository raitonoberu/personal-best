package router

import (
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/raitonoberu/personal-best/app/model"
)

func errorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	msg := err.Error()
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message.(string)
	}
	if ve, ok := err.(validator.ValidationErrors); ok {
		code = http.StatusBadRequest
		sb := strings.Builder{}
		for i, v := range ve {
			sb.WriteString(v.Field())
			sb.WriteString(": tag '")
			sb.WriteString(v.Tag())
			sb.WriteString("' failed")
			if i < len(ve)-1 {
				sb.WriteString(", ")
			}
		}
		msg = sb.String()
	}

	c.JSON(code, model.NewError(msg))
}

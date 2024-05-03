package router

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/raitonoberu/personal-best/app/handler"
	"github.com/raitonoberu/personal-best/app/model"
)

func errorHandler(err error, c echo.Context) {
	if ve, ok := err.(validator.ValidationErrors); ok {
		fields := make([]string, len(ve))
		for i, v := range ve {
			fields[i] = v.Field()
		}
		err = handler.NewValidationErr(fields)
	}

	var code int
	var msg string
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message.(string)
	} else {
		code = http.StatusInternalServerError
		msg = fmt.Sprintf("Произошла внутренняя ошибка: %s", err.Error())
	}

	c.JSON(code, model.NewError(msg))
}

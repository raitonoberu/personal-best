package handler

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

var (
	ErrNotAuthorized       = echo.NewHTTPError(401, "Вы не авторизованы")
	ErrWrongPassword       = echo.NewHTTPError(401, "Неверный пароль")
	ErrAccessDenied        = echo.NewHTTPError(403, "Недостаточно прав")
	ErrUserNotFound        = echo.NewHTTPError(404, "Пользователь не найден")
	ErrCompetitionNotFound = echo.NewHTTPError(404, "Соревнование не найдено")
	ErrInternalError       = echo.NewHTTPError(500, "Произошла внутренняя ошибка") // unused til release
)

func NewValidationErr(fields []string) *echo.HTTPError {
	return echo.NewHTTPError(400, fmt.Sprintf("Неправильно введены поля: %s", strings.Join(fields, ", ")))
}

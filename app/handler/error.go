package handler

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

var (
	ErrNotEnoughFields     = echo.NewHTTPError(400, "Для создания профиля игрока не хватает полей")
	ErrStartBeforeClose    = echo.NewHTTPError(400, "Соревнование не может начинаться до конца регистрации")
	ErrEndBeforeStart      = echo.NewHTTPError(400, "Время начала должно быть раньше времени конца")
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

package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/raitonoberu/personal-best/app/model"
)

// @Summary Get user
// @Security Bearer
// @Description Return single user by ID
// @Description "player" may not be present (trainer / admin)
// @Description player.preparation, player.position may not be present
// @Tags user
// @Produce json
// @Param request path model.GetUserRequest true "path"
// @Success 200 {object} model.GetUserResponse
// @Router /api/users/{id} [get]
func (h Handler) GetUser(c echo.Context) error {
	var req model.GetUserRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	user, err := h.service.GetUser(c.Request().Context(), req.ID)
	if err != nil {
		return err
	}
	return c.JSON(200, user)
}

// @Summary Update user
// @Security Bearer
// @Description Update user info.
// @Description Player fields will be added later.
// @Tags user
// @Param request body model.UpdateUserRequest true "body"
// @Success 204
// @Router /api/users [patch]
func (h Handler) UpdateUser(c echo.Context) error {
	var req model.UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	req.ID = getUserID(c)
	if err := h.service.UpdateUser(c.Request().Context(), req); err != nil {
		return err
	}
	return c.NoContent(204)
}

// @Summary Delete user
// @Security Bearer
// @Description Delete current user
// @Tags user
// @Success 204
// @Router /api/users [delete]
func (h Handler) DeleteUser(c echo.Context) error {
	if err := h.service.DeleteUser(c.Request().Context(), getUserID(c)); err != nil {
		return err
	}
	return c.NoContent(204)
}

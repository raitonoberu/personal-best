package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/raitonoberu/personal-best/app/model"
)

// @Summary Create user
// @Security Bearer
// @Description Create new user with desired params.
// @Description Player-related params only required when creating player
// @Description (is_male, phone, telegram, birth_date).
// @Tags admin
// @Accept json
// @Produce json
// @Param request body model.AdminCreateUserRequest true "body"
// @Success 201 {object} model.AuthResponse
// @Router /api/users [post]
func (h Handler) AdminCreateUser(c echo.Context) error {
	if err := h.ensureAdmin(c); err != nil {
		return err
	}

	var req model.AdminCreateUserRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	result, err := h.service.CreateUser(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(201, result)
}

// @Summary List users
// @Security Bearer
// @Description List users with specified role.
// @Description Can be used for checking players before approving.
// @Tags admin
// @Produce json
// @Param request query model.AdminListUsersRequest true "query"
// @Success 200 {object} model.ListUsersResponse
// @Router /api/users [get]
func (h Handler) AdminListUsers(c echo.Context) error {
	if err := h.ensureAdmin(c); err != nil {
		return err
	}

	var req model.AdminListUsersRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	resp, err := h.service.ListUsers(c.Request().Context(), req.RoleID, req.Limit, req.Offset)
	if err != nil {
		return err
	}

	return c.JSON(200, resp)
}

// @Summary Update user
// @Security Bearer
// @Description Update user.
// @Description Player-related params only changed when updating player
// @Description (is_male, phone, telegram, birth_date).
// @Tags admin
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param request body model.AdminUpdateUserRequest true "body"
// @Success 204
// @Router /api/users/{id} [patch]
func (h Handler) AdminUpdateUser(c echo.Context) error {
	if err := h.ensureAdmin(c); err != nil {
		return err
	}

	var req model.AdminUpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := h.service.AdminUpdateUser(c.Request().Context(), req); err != nil {
		return err
	}

	return c.NoContent(204)
}

// @Summary Delete user
// @Security Bearer
// @Description Delete user by id
// @Tags admin
// @Param id path int true "id"
// @Success 204
// @Router /api/users/{id} [delete]
func (h Handler) AdminDeleteUser(c echo.Context) error {
	if err := h.ensureAdmin(c); err != nil {
		return err
	}

	var req model.AdminDeleteUserRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := h.service.DeleteUser(c.Request().Context(), req.ID); err != nil {
		return err
	}

	return c.NoContent(204)
}

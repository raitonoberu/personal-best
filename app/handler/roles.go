package handler

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/raitonoberu/personal-best/app/model"
	"github.com/raitonoberu/personal-best/db/sqlc"
)

// @Summary List roles
// @Security Bearer
// @Description List all available roles
// @Tags roles
// @Produce json
// @Success 200 {object} []model.RoleResponse
// @Router /api/roles [get]
func (h Handler) ListRoles(c echo.Context) error {
	roles, err := h.queries.ListRoles(c.Request().Context())
	if err != nil {
		return err
	}

	resp := make([]model.RoleResponse, len(roles))
	for i, r := range roles {
		resp[i] = model.RoleResponse(r)
	}

	return c.JSON(200, resp)
}

func (h Handler) getUserRole(c echo.Context) sqlc.Role {
	if r := c.Get("role"); r != nil {
		return r.(sqlc.Role)
	}

	role, err := h.queries.GetUserRole(c.Request().Context(),
		getUserID(c))
	if err != nil {
		log.Printf("[ERROR] Couldn't get role for user %d: %s", getUserID(c), err)
	}
	c.Set("role", role)
	return role
}

func (h Handler) ensureCanView(c echo.Context) error {
	role := h.getUserRole(c)
	if !role.CanView {
		return ErrAccessDenied
	}
	return nil
}

func (h Handler) ensureCanParticipate(c echo.Context) error {
	role := h.getUserRole(c)
	if !role.CanParticipate {
		return ErrAccessDenied
	}
	return nil
}

func (h Handler) ensureCanCreate(c echo.Context) error {
	role := h.getUserRole(c)
	if !role.CanCreate {
		return ErrAccessDenied
	}
	return nil
}

func (h Handler) ensureAdmin(c echo.Context) error {
	role := h.getUserRole(c)
	if !role.IsAdmin {
		return ErrAccessDenied
	}
	return nil
}

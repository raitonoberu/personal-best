package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/raitonoberu/personal-best/app/model"
)

// @Summary List roles
// @Security Bearer
// @Description List all available roles
// @Tags admin
// @Produce json
// @Success 200 {object} []model.RoleResponse
// @Router /api/admin/roles [get]
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

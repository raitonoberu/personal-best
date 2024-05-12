package handler

import (
	"context"

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

func (h Handler) getUserRole(ctx context.Context, id int64) *sqlc.Role {
	role, err := h.queries.GetUserRole(ctx, id)
	if err != nil {
		return nil
	}
	return &role
}

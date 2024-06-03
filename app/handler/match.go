package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/raitonoberu/personal-best/app/model"
)

// @Summary List competition matches
// @Security Bearer
// @Description List all matches with all players
// @Tags match
// @Produce json
// @Param id path int true "competition id"
// @Param request query model.ListMatchesRequest true "query"
// @Success 200 {object} model.ListMatchesResponse
// @Router /api/competitions/{id}/matches [get]
func (h Handler) ListMatches(c echo.Context) error {
	if err := h.ensureCanView(c); err != nil {
		return err
	}

	var req model.ListMatchesRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	matches, err := h.service.ListMatches(c.Request().Context(), req)
	if err != nil {
		return err
	}
	return c.JSON(200, matches)
}

// @Summary Update match
// @Security Bearer
// @Description Update match score
// @Description Match must have players & score must NOT be set already
// @Description This will fill next match's players
// @Tags match
// @Param comp_id path int true "comp_id"
// @Param id path int true "id"
// @Param request body model.UpdateMatchRequest true "body"
// @Success 204
// @Router /api/competitions/{comp_id}/matches/{id} [patch]
func (h Handler) UpdateMatch(c echo.Context) error {
	if err := h.ensureCanCreate(c); err != nil {
		return err
	}

	var req model.UpdateMatchRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	err := h.service.SetMatchScore(c.Request().Context(), req.ID, req.LeftScore, req.RightScore)
	if err != nil {
		return err
	}

	return c.NoContent(204)
}

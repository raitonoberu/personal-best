package handler

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/raitonoberu/personal-best/app/model"
)

// @Summary Create competition
// @Security Bearer
// @Description Create new competition.
// @Description Days must be different (no same day twice).
// @Description Time must be in format HH:MM.
// @Tags competition
// @Accept json
// @Produce json
// @Param request body model.CreateCompetitionRequest true "body"
// @Success 201 {object} model.CreateCompetitionResponse
// @Router /api/competitions [post]
func (h Handler) CreateCompetition(c echo.Context) error {
	if err := h.ensureCanCreate(c); err != nil {
		return err
	}

	var req model.CreateCompetitionRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	req.UserID = getUserID(c)
	id, err := h.service.CreateCompetition(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(201, model.CreateCompetitionResponse{ID: id})
}

// @Summary Get competition
// @Security Bearer
// @Description Return competition by ID
// @Tags competition
// @Produce json
// @Param request path model.GetCompetitionRequest true "path"
// @Success 200 {object} model.GetCompetitionResponse
// @Router /api/competitions/{id} [get]
func (h Handler) GetCompetition(c echo.Context) error {
	var req model.GetCompetitionRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	comp, err := h.service.GetCompetition(c.Request().Context(), req.ID)
	if err != nil {
		return err
	}

	return c.JSON(200, comp)
}

// @Summary List competition
// @Security Bearer
// @Description List all competitions from new to old
// @Description For now there is no way to get start/end but im working on it :)
// @Tags competition
// @Produce json
// @Param request query model.ListCompetitionsRequest true "query"
// @Success 200 {object} model.ListCompetitionsResponse
// @Router /api/competitions [get]
func (h Handler) ListCompetitions(c echo.Context) error {
	var req model.ListCompetitionsRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	resp, err := h.service.ListCompetitions(c.Request().Context(), req.Limit, req.Offset)
	if err != nil {
		return err
	}

	return c.JSON(200, resp)
}

// @Summary Update competition
// @Security Bearer
// @Tags competition
// @Param id path int true "comp id"
// @Param request body model.UpdateCompetitionRequest true "body"
// @Success 204
// @Router /api/competitions/{id} [patch]
func (h Handler) UpdateCompetition(c echo.Context) error {
	if err := h.ensureCanCreate(c); err != nil {
		return err
	}

	var req model.UpdateCompetitionRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := h.service.UpdateCompetition(c.Request().Context(), req); err != nil {
		return err
	}

	return c.NoContent(204)
}

// @Summary Delete competition
// @Security Bearer
// @Tags competition
// @Param id path int true "id of competition"
// @Success 204
// @Router /api/competitions/{id} [delete]
func (h Handler) DeleteCompetition(c echo.Context) error {
	if err := h.ensureCanCreate(c); err != nil {
		return err
	}

	var req model.DeleteCompetitionRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := h.service.DeleteCompetition(c.Request().Context(), req.ID); err != nil {
		return err
	}

	return c.NoContent(204)
}

// @Summary Get competition scores
// @Security Bearer
// @Description Get final scores for all competition players
// @Tags competition
// @Produce json
// @Param id path int true "competition id"
// @Success 200 {object} []model.CompetitionScore
// @Router /api/competitions/{id}/scores [get]
func (h Handler) ListScores(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	scores, err := h.service.GetScores(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(200, scores)
}

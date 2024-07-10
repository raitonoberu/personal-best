package handler

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/raitonoberu/personal-best/app/model"
)

// @Summary List competition registrations
// @Security Bearer
// @Description List registrations for competition
// @Tags registration
// @Produce json
// @Param request path model.ListCompetitionRegistrationsRequest true "path"
// @Success 200 {object} []model.CompetitionRegistration
// @Router /api/competitions/{id}/registrations [get]
func (h Handler) ListCompetitionRegistrations(c echo.Context) error {
	var req model.ListCompetitionRegistrationsRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	regs, err := h.service.ListCompetitionRegistrations(c.Request().Context(), req.ID)
	if err != nil {
		return err
	}

	return c.JSON(200, regs)
}

// @Summary List player registrations
// @Security Bearer
// @Description List competitions where player is registered
// @Tags registration
// @Produce json
// @Param user_id path int true "id of user"
// @Param query query model.ListPlayerRegistrationsRequest true "query"
// @Success 200 {object} model.ListPlayerRegistrationsResponse
// @Router /api/users/{user_id}/registrations [get]
func (h Handler) ListPlayerRegistrations(c echo.Context) error {
	var req model.ListPlayerRegistrationsRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	id, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)

	regs, err := h.service.ListPlayerRegistrations(c.Request().Context(), id, req.Limit, req.Offset)
	if err != nil {
		return err
	}

	return c.JSON(200, regs)
}

// @Summary Register for competition
// @Security Bearer
// @Description Competition must not be closed yet
// @Tags registration
// @Param request path model.RegisterForCompetitionRequest true "path"
// @Success 204
// @Router /api/competitions/{id}/registrations [post]
func (h Handler) RegisterForCompetition(c echo.Context) error {
	if err := h.ensureCanParticipate(c); err != nil {
		return err
	}

	var req model.RegisterForCompetitionRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	err := h.service.RegisterForCompetition(c.Request().Context(), getUserID(c), req.ID)
	if err != nil {
		return err
	}

	return c.NoContent(204)
}

// @Summary Unregister for competition
// @Security Bearer
// @Description Competition must not be closed yet
// @Tags registration
// @Param request path model.UnregisterForCompetitionRequest true "path"
// @Success 204
// @Router /api/competitions/{id}/registrations [delete]
func (h Handler) UnregisterForCompetition(c echo.Context) error {
	if err := h.ensureCanParticipate(c); err != nil {
		return err
	}

	var req model.UnregisterForCompetitionRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	err := h.service.UnregisterForCompetition(c.Request().Context(), getUserID(c), req.ID)
	if err != nil {
		return err
	}

	return c.NoContent(204)
}

// @Summary Update registration
// @Security Bearer
// @Description This is made for trainers/admins
// @Description Here you can approve or drop players
// @Tags registration
// @Param player_id path int true "player_id"
// @Param comp_id path int true "comp_id"
// @Param request body model.UpdateRegistrationRequest true "body"
// @Success 204
// @Router /api/competitions/{comp_id}/registrations/{player_id} [patch]
func (h Handler) UpdateRegistration(c echo.Context) error {
	if err := h.ensureCanCreate(c); err != nil {
		return err
	}

	var req model.UpdateRegistrationRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	err := h.service.UpdateRegistration(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.NoContent(204)
}

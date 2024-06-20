package handler

import (
	"database/sql"
	"errors"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raitonoberu/personal-best/app/model"
	"github.com/raitonoberu/personal-best/app/service"
	"github.com/raitonoberu/personal-best/db/sqlc"
)

// @Summary List competition registrations
// @Security Bearer
// @Description List registrations for competition
// @Tags registration
// @Produce json
// @Param request path model.ListCompetitionRegistrationsRequest true "path"
// @Success 200 {object} model.ListCompetitionRegistrationsResponse
// @Router /api/competitions/{id}/registrations [get]
func (h Handler) ListCompetitionRegistrations(c echo.Context) error {
	var req model.ListCompetitionRegistrationsRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	rows, err := h.queries.ListCompetitionRegistrations(c.Request().Context(), req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrCompetitionNotFound
		}
		return err
	}

	return c.JSON(200, model.NewListCompetitionRegistrationsResponse(rows))
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

	competition, err := h.queries.GetCompetition(c.Request().Context(), req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrCompetitionNotFound
		}
		return err
	}

	// here we add 24h cuz stored ClosesAt is the last day available to register
	if competition.Competition.ClosesAt.Add(time.Hour * 24).Before(time.Now()) {
		return ErrCompetitionClosed
	}

	if err := h.queries.CreateRegistration(c.Request().Context(),
		sqlc.CreateRegistrationParams{
			CompetitionID: req.ID,
			PlayerID:      getUserID(c),
			IsApproved:    false,
			IsDropped:     false,
		}); err != nil {
		return err
	}

	if err := h.service.UpdateMatches(c.Request().Context(), req.ID); err != nil &&
		err != service.ErrNotEnoughPlayers {
		// ignoring if not enough players
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

	competition, err := h.queries.GetCompetition(c.Request().Context(), req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrCompetitionNotFound
		}
		return err
	}

	// here we add 24h cuz stored ClosesAt is the last day available to register
	if competition.Competition.ClosesAt.Add(time.Hour * 24).Before(time.Now()) {
		return ErrCompetitionClosed
	}

	reg, err := h.queries.GetRegistration(c.Request().Context(),
		sqlc.GetRegistrationParams{
			CompetitionID: req.ID,
			PlayerID:      getUserID(c),
		})
	if err != nil {
		return err
	}
	if reg.IsDropped {
		return ErrPlayerDropped
	}

	if err := h.queries.DeleteRegistration(c.Request().Context(),
		sqlc.DeleteRegistrationParams{
			CompetitionID: req.ID,
			PlayerID:      getUserID(c),
		}); err != nil {
		return err
	}

	if err := h.service.UpdateMatches(c.Request().Context(), req.ID); err != nil &&
		err != service.ErrNotEnoughPlayers {
		// ignoring if not enough players
		// TODO: Here we actually should not xD
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

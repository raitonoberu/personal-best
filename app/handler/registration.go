package handler

import (
	"database/sql"
	"errors"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raitonoberu/personal-best/app/model"
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
	if err := h.ensureCanView(c); err != nil {
		return err
	}

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
// @Success 201
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

	return c.NoContent(201)
}

// @Summary Unregister for competition
// @Security Bearer
// @Description Competition must not be closed yet
// @Tags registration
// @Param request path model.UnregisterForCompetitionRequest true "path"
// @Success 200
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

	if err := h.queries.DeleteRegistration(c.Request().Context(),
		sqlc.DeleteRegistrationParams{
			CompetitionID: req.ID,
			PlayerID:      getUserID(c),
		}); err != nil {
		return err
	}

	return c.NoContent(200)
}

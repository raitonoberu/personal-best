package handler

import (
	"database/sql"
	"errors"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raitonoberu/personal-best/app/model"
	"github.com/raitonoberu/personal-best/db/sqlc"
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

	closesAt := parseDate(req.ClosesAt)
	for _, d := range req.Days {
		day := parseDate(d.Date)
		if day.Equal(closesAt) || day.Before(closesAt) {
			return ErrStartBeforeClose
		}

		startTime := parseTime(d.StartTime, day)
		endTime := parseTime(d.EndTime, day)

		if startTime.After(endTime) {
			return ErrEndBeforeStart
		}
	}

	// TODO: here we need MUCH MORE CHECKS BITCH

	tx, err := h.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	qtx := h.queries.WithTx(tx)

	comp, err := qtx.CreateCompetition(c.Request().Context(),
		sqlc.CreateCompetitionParams{
			TrainerID:   getUserID(c),
			Name:        req.Name,
			Description: req.Description,
			Tours:       req.Tours,
			Age:         req.Age,
			Size:        req.Size,
			ClosesAt:    closesAt,
		})
	if err != nil {
		return err
	}

	for _, d := range req.Days {
		date := parseDate(d.Date)
		_, err := qtx.CreateCompetitionDay(c.Request().Context(),
			sqlc.CreateCompetitionDayParams{
				CompetitionID: comp.ID,
				Date:          date,
				StartTime:     parseTime(d.StartTime, date),
				EndTime:       parseTime(d.EndTime, date),
			})
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	// TODO: make it in the SAME TRANSACTION
	if err := h.service.GenerateMatches(c.Request().Context(), comp.ID); err != nil {
		return err
	}

	return c.JSON(201, model.NewCreateCompetitionResponse(comp))
}

// @Summary Get competition
// @Security Bearer
// @Description Return competition by ID
// @Descripiton Not much yet :)
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

	competition, err := h.queries.GetCompetition(c.Request().Context(), req.ID)
	if err != nil {
		return err
	}

	days, err := h.queries.GetCompetitionDays(c.Request().Context(), req.ID)
	if err != nil {
		return err
	}

	return c.JSON(200, model.NewGetCompetitionResponse(competition, days))
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

	competitions, err := h.queries.ListCompetitions(c.Request().Context(),
		sqlc.ListCompetitionsParams(req))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrCompetitionNotFound
		}
		return err
	}
	return c.JSON(200, model.NewListCompetitionsResponse(competitions))
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

	rows, err := h.queries.ListPlayerRegistrations(c.Request().Context(),
		sqlc.ListPlayerRegistrationsParams{
			PlayerID: id,
			Limit:    req.Limit,
			Offset:   req.Offset,
		})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrCompetitionNotFound
		}
		return err
	}
	return c.JSON(200, model.NewListPlayerRegistrationsResponse(rows))
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

	var closesAt *time.Time
	if req.ClosesAt != nil {
		date := parseDate(*req.ClosesAt)
		closesAt = &date

		// TODO: check if it's not after first match and not in past
	}

	err := h.queries.UpdateCompetition(c.Request().Context(),
		sqlc.UpdateCompetitionParams{
			ID:          req.ID,
			Name:        req.Name,
			Description: req.Description,
			ClosesAt:    closesAt,
		})
	if err != nil {
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

	if err := h.queries.DeleteCompetition(c.Request().Context(), req.ID); err != nil {
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

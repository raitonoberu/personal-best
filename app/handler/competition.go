package handler

import (
	"database/sql"
	"errors"
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
	if err := h.ensureCanView(c); err != nil {
		return err
	}

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
	if err := h.ensureCanView(c); err != nil {
		return err
	}

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

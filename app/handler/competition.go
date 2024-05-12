package handler

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raitonoberu/personal-best/app/model"
	"github.com/raitonoberu/personal-best/db/sqlc"
)

func (h Handler) CreateCompetition(c echo.Context) error {
	var req model.CreateCompetitionRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	comp, err := h.queries.CreateCompetition(c.Request().Context(),
		sqlc.CreateCompetitionParams{
			TrainerID:   getUserID(c),
			Name:        req.Name,
			Description: req.Description,
			StartDate:   parseDate(req.StartDate),
			Tours:       req.Tours,
			Age:         req.Age,
			Size:        req.Size,
			ClosesAt:    parseDate(req.ClosesAt),
		})
	if err != nil {
		return err
	}
	return c.JSON(201, model.NewCreateCompetitionResponse(comp))
}

func (h Handler) GetCompetition(c echo.Context) error {
	var req model.GetCompetitionRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	competition, err := h.queries.GetCompetition(c.Request().Context(), req.ID)
	if err != nil {
		return err
	}
	return c.JSON(200, model.NewGetCompetitionResponse(competition))
}

func (h Handler) ListCompetitions(c echo.Context) error {
	var req model.ListCompetitionsRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	competitions, err := h.queries.ListCompetitions(c.Request().Context(),
		sqlc.ListCompetitionsParams(req))
	if err != nil {
		return err
	}
	return c.JSON(200, model.NewListCompetitionsResponse(competitions))
}

func (h Handler) UpdateCompetition(c echo.Context) error {
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
	var req model.DeleteCompetitionRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := h.queries.DeleteCompetition(c.Request().Context(), req.ID); err != nil {
		return err
	}
	return c.NoContent(204)
}

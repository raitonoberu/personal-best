package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/raitonoberu/personal-best/app/model"
	"github.com/raitonoberu/personal-best/db/sqlc"
)

func (h Handler) CreateCompetition(c echo.Context) error {
	var req model.CreateCompetitionRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	req.TrainerID = getUserID(c)

	comp, err := h.db.CreateCompetition(c.Request().Context(),
		sqlc.CreateCompetitionParams(req))
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

	competition, err := h.db.GetCompetition(c.Request().Context(), req.ID)
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

	competitions, err := h.db.ListCompetitions(c.Request().Context(),
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

	err := h.db.UpdateCompetition(c.Request().Context(),
		sqlc.UpdateCompetitionParams(req))
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

	if err := h.db.DeleteCompetition(c.Request().Context(), req.ID); err != nil {
		return err
	}
	return c.NoContent(204)
}

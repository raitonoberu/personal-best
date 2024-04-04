package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/raitonoberu/personal-best/app/model"
)

func (h Handler) GetCompetition(c echo.Context) error {
	var req model.GetCompetitionRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := c.Validate(&req); err != nil {
		return err
	}

	competition, err := h.db.GetCompetition(c.Request().Context(), req.ID)
	if err != nil {
		return err
	}
	return c.JSON(200, model.NewGetCompetitionResponse(competition))
}

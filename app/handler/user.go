package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/raitonoberu/personal-best/app/model"
	"github.com/raitonoberu/personal-best/db/sqlc"
)

func (h Handler) GetUser(c echo.Context) error {
	var req model.GetUserRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := c.Validate(&req); err != nil {
		return err
	}

	user, err := h.db.GetUser(c.Request().Context(), req.ID)
	if err != nil {
		return err
	}
	return c.JSON(200, model.NewGetUserResponse(user))
}

func (h Handler) ListUsers(c echo.Context) error {
	var req model.ListUsersRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := c.Validate(&req); err != nil {
		return err
	}

	users, err := h.db.ListUsers(c.Request().Context(), sqlc.ListUsersParams(req))
	if err != nil {
		return err
	}
	return c.JSON(200, model.NewListUsersResponse(users))
}

func (h Handler) DeleteUser(c echo.Context) error {
	if err := h.db.DeleteUser(c.Request().Context(), getUserID(c)); err != nil {
		return err
	}
	return c.NoContent(204)
}

package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/raitonoberu/personal-best/app/model"
	"github.com/raitonoberu/personal-best/db/sqlc"
	"golang.org/x/crypto/bcrypt"
)

func (h Handler) GetUser(c echo.Context) error {
	var req model.GetUserRequest
	if err := c.Bind(&req); err != nil {
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

	users, err := h.db.ListUsers(c.Request().Context(),
		sqlc.ListUsersParams(req))
	if err != nil {
		return err
	}
	return c.JSON(200, model.NewListUsersResponse(users))
}

func (h Handler) UpdateUser(c echo.Context) error {
	var req model.UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	req.ID = getUserID(c)
	if req.Password != nil {
		hash, err := bcrypt.GenerateFromPassword(
			[]byte(*req.Password), bcrypt.DefaultCost,
		)
		if err != nil {
			return err
		}
		hashStr := string(hash)
		req.Password = &hashStr
	}

	if err := h.db.UpdateUser(c.Request().Context(),
		sqlc.UpdateUserParams(req)); err != nil {
		return err
	}
	return c.NoContent(204)
}

func (h Handler) DeleteUser(c echo.Context) error {
	if err := h.db.DeleteUser(c.Request().Context(), getUserID(c)); err != nil {
		return err
	}
	return c.NoContent(204)
}

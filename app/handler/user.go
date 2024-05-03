package handler

import (
	"database/sql"
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/raitonoberu/personal-best/app/model"
	"github.com/raitonoberu/personal-best/db/sqlc"
	"golang.org/x/crypto/bcrypt"
)

// @Summary Get user
// @Security Bearer
// @Description Return single user by ID
// @Description "player" may not be present (trainer / admin)
// @Description player.preparation, player.position may not be present
// @Tags user
// @Produce json
// @Param request path model.GetUserRequest true "path"
// @Success 200 {object} model.GetUserResponse
// @Router /api/users/{id} [get]
func (h Handler) GetUser(c echo.Context) error {
	var req model.GetUserRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	user, err := h.queries.GetUser(c.Request().Context(), req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrUserNotFound
		}
		return err
	}
	return c.JSON(200, model.NewGetUserResponse(user.User, user.Player))
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

	if err := h.queries.UpdateUser(c.Request().Context(),
		sqlc.UpdateUserParams(req),
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrUserNotFound
		}
		return err
	}
	return c.NoContent(204)
}

// @Summary Delete user
// @Security Bearer
// @Description Delete current user
// @Tags user
// @Success 204
// @Router /api/users [delete]
func (h Handler) DeleteUser(c echo.Context) error {
	if err := h.queries.DeleteUser(c.Request().Context(), getUserID(c)); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrUserNotFound
		}
		return err
	}
	return c.NoContent(204)
}

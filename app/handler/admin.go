package handler

import (
	"reflect"

	"github.com/labstack/echo/v4"
	"github.com/raitonoberu/personal-best/app/model"
	"github.com/raitonoberu/personal-best/db/sqlc"
)

// @Summary Create user
// @Security Bearer
// @Description Create new user with desired params.
// @Description Player-related params only required when creating player
// @Description (is_male, phone, telegram, birth_date).
// @Tags admin
// @Accept json
// @Produce json
// @Param request body model.AdminCreateUserRequest true "body"
// @Success 200 {object} model.AuthResponse
// @Router /api/admin/users [post]
func (h Handler) AdminCreateUser(c echo.Context) error {
	var req model.AdminCreateUserRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := h.ensureAdmin(c); err != nil {
		return err
	}

	role, err := h.queries.GetRole(c.Request().Context(), req.RoleID)
	if err != nil {
		return err
	}

	playerFields := []any{
		req.BirthDate,
		req.IsMale,
		req.Phone,
		req.Telegram,
	}

	missingPlayerFields := 0
	for _, f := range playerFields {
		// TODO: это было бы хорошо делать на уровне валидатора
		if reflect.ValueOf(f).IsNil() {
			missingPlayerFields += 1
		}
	}

	if missingPlayerFields != 0 && role.CanParticipate {
		return ErrNotEnoughFields
	}

	password, err := generateHash(req.Password)
	if err != nil {
		return err
	}

	tx, err := h.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	qtx := h.queries.WithTx(tx)

	user, err := qtx.CreateUser(c.Request().Context(),
		sqlc.CreateUserParams{
			RoleID:     req.RoleID,
			Email:      req.Email,
			Password:   password,
			FirstName:  req.FirstName,
			LastName:   req.LastName,
			MiddleName: req.MiddleName,
		},
	)
	if err != nil {
		return err
	}

	if role.CanParticipate {
		// creating player
		_, err = qtx.CreatePlayer(c.Request().Context(),
			sqlc.CreatePlayerParams{
				UserID:    user.ID,
				IsMale:    *req.IsMale,
				Phone:     *req.Phone,
				Telegram:  *req.Telegram,
				BirthDate: parseDate(*req.BirthDate),
			},
		)
		if err != nil {
			return err
		}
	}

	token, err := generateToken(user)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return c.JSON(200, model.NewAuthResponse(user.ID, token))
}

func (h Handler) ensureAdmin(c echo.Context) error {
	role := h.getUserRole(c.Request().Context(), getUserID(c))
	if role == nil {
		return ErrUserNotFound
	}
	if !role.IsAdmin {
		return ErrAccessDenied
	}
	return nil
}

package handler

import (
	"reflect"
	"time"

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
// @Success 201 {object} model.AuthResponse
// @Router /api/users [post]
func (h Handler) AdminCreateUser(c echo.Context) error {
	if err := h.ensureAdmin(c); err != nil {
		return err
	}

	var req model.AdminCreateUserRequest
	if err := c.Bind(&req); err != nil {
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
	return c.JSON(201, model.NewAuthResponse(user.ID, token))
}

// @Summary List users
// @Security Bearer
// @Description List users with specified role.
// @Description Can be used for checking players before approving.
// @Tags admin
// @Produce json
// @Param request query model.AdminListUsersRequest true "query"
// @Success 200 {object} model.ListUsersResponse
// @Router /api/users [get]
func (h Handler) AdminListUsers(c echo.Context) error {
	if err := h.ensureAdmin(c); err != nil {
		return err
	}

	var req model.AdminListUsersRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	resp, err := h.service.ListUsers(c.Request().Context(), req.RoleID, req.Limit, req.Offset)
	if err != nil {
		return err
	}
	return c.JSON(200, resp)
}

// @Summary Update user
// @Security Bearer
// @Description Update user.
// @Description Player-related params only changed when updating player
// @Description (is_male, phone, telegram, birth_date).
// @Tags admin
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param request body model.AdminUpdateUserRequest true "body"
// @Success 204
// @Router /api/users/{id} [patch]
func (h Handler) AdminUpdateUser(c echo.Context) error {
	if err := h.ensureAdmin(c); err != nil {
		return err
	}

	var req model.AdminUpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	if req.RoleID != nil {
		if _, err := h.queries.GetRole(c.Request().Context(), *req.RoleID); err != nil {
			return err
		}
	}

	tx, err := h.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	qtx := h.queries.WithTx(tx)

	if err := qtx.UpdateUser(c.Request().Context(),
		sqlc.UpdateUserParams{
			ID:         req.ID,
			RoleID:     req.RoleID,
			Email:      req.Email,
			Password:   req.Password,
			FirstName:  req.FirstName,
			MiddleName: req.MiddleName,
			LastName:   req.LastName,
		},
	); err != nil {
		return err
	}

	var birthDate *time.Time
	if req.BirthDate != nil {
		date := parseDate(*req.BirthDate)
		birthDate = &date
	}

	if err := qtx.UpdatePlayer(c.Request().Context(),
		sqlc.UpdatePlayerParams{
			UserID:    req.ID,
			BirthDate: birthDate,
			IsMale:    req.IsMale,
			Phone:     req.Phone,
			Telegram:  req.Telegram,
		},
	); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return c.NoContent(204)
}

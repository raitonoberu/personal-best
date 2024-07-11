package service

import (
	"context"
	"reflect"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raitonoberu/personal-best/app/model"
	"github.com/raitonoberu/personal-best/db/sqlc"
	"golang.org/x/crypto/bcrypt"
)

var ErrNotEnoughFields = echo.NewHTTPError(400, "Для создания профиля игрока не хватает полей")

func (s Service) CreateUser(ctx context.Context, req model.AdminCreateUserRequest) (*model.AuthResponse, error) {
	role, err := s.queries.GetRole(ctx, req.RoleID)
	if err != nil {
		return nil, err
	}

	playerFields := []any{
		req.BirthDate,
		req.IsMale,
		req.Phone,
		req.Telegram,
		req.Position,
		req.Preparation,
	}

	missingPlayerFields := 0
	for _, f := range playerFields {
		// TODO: это было бы хорошо делать на уровне валидатора
		if reflect.ValueOf(f).IsNil() {
			missingPlayerFields += 1
		}
	}

	if missingPlayerFields != 0 && !role.CanCreate {
		return nil, ErrNotEnoughFields
	}

	password, err := generateHash(req.Password)
	if err != nil {
		return nil, err
	}

	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	qtx := s.queries.WithTx(tx)

	user, err := qtx.CreateUser(ctx,
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
		return nil, err
	}

	if role.CanParticipate {
		// creating player
		_, err = qtx.CreatePlayer(ctx,
			sqlc.CreatePlayerParams{
				UserID:      user.ID,
				IsMale:      *req.IsMale,
				Phone:       *req.Phone,
				Telegram:    *req.Telegram,
				BirthDate:   parseDate(*req.BirthDate),
				Preparation: *req.Preparation,
				Position:    *req.Position,
			},
		)
		if err != nil {
			return nil, err
		}
	}

	token, err := generateToken(user)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return &model.AuthResponse{
		ID:    user.ID,
		Token: token,
	}, nil
}

func (s Service) AdminUpdateUser(ctx context.Context, req model.AdminUpdateUserRequest) error {
	if req.RoleID != nil {
		if _, err := s.queries.GetRole(ctx, *req.RoleID); err != nil {
			return err
		}
	}

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

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	qtx := s.queries.WithTx(tx)

	if err := qtx.UpdateUser(ctx,
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

	if err := qtx.UpdatePlayer(ctx,
		sqlc.UpdatePlayerParams{
			UserID:      req.ID,
			BirthDate:   birthDate,
			IsMale:      req.IsMale,
			Phone:       req.Phone,
			Telegram:    req.Telegram,
			Preparation: req.Preparation,
			Position:    req.Position,
		},
	); err != nil {
		return err
	}

	return tx.Commit()
}

func (s Service) ListUsers(ctx context.Context, roleID, limit, offset int64) (*model.ListUsersResponse, error) {
	rows, err := s.queries.ListUsers(ctx, sqlc.ListUsersParams{
		RoleID: roleID,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	items := make([]model.GetUserResponse, len(rows))
	for i, r := range rows {
		items[i] = model.NewGetUserResponse(sqlc.GetUserRow{
			User:       r.User,
			UserPlayer: r.UserPlayer,
		})
	}

	var total int
	if len(items) != 0 {
		total = int(rows[0].Total)
	}

	return &model.ListUsersResponse{
		Count: len(items),
		Total: total,
		Users: items,
	}, nil
}

package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raitonoberu/personal-best/app/model"
	"github.com/raitonoberu/personal-best/db/sqlc"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound = echo.NewHTTPError(404, "Пользователь не найден")
)

func (s Service) GetUser(ctx context.Context, id int64) (*model.GetUserResponse, error) {
	userRow, err := s.queries.GetUser(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	user := model.NewGetUserResponse(userRow)
	return &user, nil
}

func (s Service) UpdateUser(ctx context.Context, req model.UpdateUserRequest) error {
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
			UserID:    req.ID,
			BirthDate: birthDate,
			IsMale:    req.IsMale,
			Phone:     req.Phone,
			Telegram:  req.Telegram,
		},
	); err != nil {
		return err
	}
	return tx.Commit()
}

func (s Service) DeleteUser(ctx context.Context, id int64) error {
	err := s.queries.DeleteUser(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrUserNotFound
	}
	return err
}

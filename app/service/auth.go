package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/raitonoberu/personal-best/app/model"
	"github.com/raitonoberu/personal-best/db/sqlc"
	"golang.org/x/crypto/bcrypt"
)

var ErrWrongPassword = echo.NewHTTPError(401, "Неверный пароль")

func (s Service) Register(ctx context.Context, req model.RegisterRequest) (*model.AuthResponse, error) {
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
			RoleID:     3, // Unverified User
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

	_, err = qtx.CreatePlayer(ctx,
		sqlc.CreatePlayerParams{
			UserID:    user.ID,
			IsMale:    *req.IsMale,
			Phone:     req.Phone,
			Telegram:  req.Telegram,
			BirthDate: parseDate(req.BirthDate),
		},
	)
	if err != nil {
		return nil, err
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

func (s Service) Login(ctx context.Context, req model.LoginRequest) (*model.AuthResponse, error) {
	user, err := s.queries.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	// check if password matches the hash in the database
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		return nil, ErrWrongPassword
	}

	token, err := generateToken(user)
	if err != nil {
		return nil, err
	}
	return &model.AuthResponse{
		ID:    user.ID,
		Token: token,
	}, nil
}

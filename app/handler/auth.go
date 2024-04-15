package handler

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/raitonoberu/personal-best/app/model"
	"github.com/raitonoberu/personal-best/db/sqlc"
	"golang.org/x/crypto/bcrypt"
)

var secret = os.Getenv("SECRET")

// Register creates a new user and sets a JWT token cookie.
func (h Handler) Register(c echo.Context) error {
	var req model.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	// generate password hash to store
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password), bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}
	req.Password = string(hash)

	tx, err := h.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	qtx := h.queries.WithTx(tx)

	user, err := qtx.CreateUser(c.Request().Context(),
		sqlc.CreateUserParams{
			RoleID:     3, // TODO
			Email:      req.Email,
			Password:   req.Password,
			FirstName:  req.FirstName,
			LastName:   req.LastName,
			MiddleName: req.MiddleName,
		},
	)
	if err != nil {
		return err
	}

	_, err = qtx.CreatePlayer(c.Request().Context(),
		sqlc.CreatePlayerParams{
			UserID:    user.ID,
			IsMale:    req.IsMale,
			Phone:     req.Phone,
			Telegram:  req.Telegram,
			BirthDate: req.BirthDate,
		},
	)
	if err != nil {
		return err
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

// Login creates a JWT token for the user and sets it as a cookie.
func (h Handler) Login(c echo.Context) error {
	var req model.LoginRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	user, err := h.queries.GetUserByEmail(c.Request().Context(), req.Email)
	if err != nil {
		return err
	}

	// check if password matches the hash in the database
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		return c.NoContent(403)
	}

	token, err := generateToken(user)
	if err != nil {
		return err
	}
	return c.JSON(200, model.NewAuthResponse(user.ID, token))
}

// generateToken generates a JWT token for the given user.
func generateToken(user sqlc.User) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{"userID": user.ID},
	)
	return token.SignedString([]byte(secret))
}

// getUserID returns the user ID from the context.
// If the user is not authenticated, it returns 0.
func getUserID(c echo.Context) int64 {
	if userID, ok := c.Get("userID").(int64); ok {
		return userID
	}
	return 0
}

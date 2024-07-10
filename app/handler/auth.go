package handler

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/raitonoberu/personal-best/app/model"
)

var secret = os.Getenv("SECRET")

// @Summary Register user
// @Description Register new unverified player
// @Description "birth_date" must have format 1889-04-20
// @Description "phone" must start with +
// @Description "telegram" must start with @
// @Tags auth
// @Accept json
// @Produce json
// @Param request body model.RegisterRequest true "body"
// @Success 201 {object} model.AuthResponse
// @Router /api/register [post]
func (h Handler) Register(c echo.Context) error {
	var req model.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	result, err := h.service.Register(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(201, result)
}

// @Summary Login user
// @Description Login user, return JWT token & ID
// @Tags auth
// @Accept json
// @Produce json
// @Param request body model.LoginRequest true "body"
// @Success 200 {object} model.AuthResponse
// @Router /api/login [post]
func (h Handler) Login(c echo.Context) error {
	var req model.LoginRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	result, err := h.service.Login(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(200, result)
}

// getUserID returns the user ID from the context.
// If the user is not authenticated, it returns 0.
func getUserID(c echo.Context) int64 {
	if userID, ok := c.Get("userID").(int64); ok {
		return userID
	}
	return 0
}

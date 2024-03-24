package middleware

import (
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/raitonoberu/personal-best/app/model"
	"gorm.io/gorm"
)

var secret = os.Getenv("SECRET")

func Auth(c fiber.Ctx) error {
	user := extractUser(c)
	if user != nil {
		c.Locals("user", user)
	}
	return c.Next()
}

func MustAuth(c fiber.Ctx) error {
	if c.Locals("user") == nil {
		return redirectToLogin(c)
	}
	return c.Next()
}

func extractUser(c fiber.Ctx) *model.User {
	tokenCookie := c.Cookies("token")
	if tokenCookie == "" {
		return nil
	}
	token, err := jwt.Parse(tokenCookie, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil
	}
	user := &model.User{
		Model: gorm.Model{ID: uint(claims["id"].(float64))},
		Name:  claims["name"].(string),
		Email: claims["email"].(string),
	}
	return user
}

func redirectToLogin(c fiber.Ctx) error {
	c.Response().Header.Set("HX-Redirect", "/login")
	return c.SendStatus(fiber.StatusUnauthorized)
}

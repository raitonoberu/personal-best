package middleware

import (
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var secret = os.Getenv("SECRET")

// Auth sets the user ID in the context.
// If the user is not authenticated, it sets the user ID to 0.
func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := extractUserID(c)
		c.Set("userID", userID)
		return next(c)
	}
}

// MustAuth sets the user ID in the context.
// If the user is not authenticated, it returns 401.
func MustAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Get("userID") == 0 {
			return c.NoContent(401)
		}
		return next(c)
	}
}

// extractUserID extracts the user ID from the JWT token in the header.
func extractUserID(c echo.Context) int64 {
	header := c.Request().Header.Get("Authorization")
	if header == "" {
		return 0
	}

	tokenString := strings.TrimPrefix(header, "Bearer ")
	token, err := jwt.Parse(tokenString,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
	if err != nil || !token.Valid {
		return 0
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return int64(claims["userID"].(float64))
	}
	return 0
}

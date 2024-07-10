package service

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/raitonoberu/personal-best/db/sqlc"
	"golang.org/x/crypto/bcrypt"
)

var secret = os.Getenv("SECRET")

// parse date in format YYYY-MM-DD
func parseDate(date string) time.Time {
	time, _ := time.Parse("2006-01-02", date)
	// we are using validator to ensure it's in proper format
	return time
}

// parse time in format HH:MM
func parseTime(timeStr string, date time.Time) time.Time {
	time, _ := time.Parse("15:04", timeStr)
	// we are using validator to ensure it's in proper format
	return time.AddDate(date.Year(), int(date.Month()), date.Day())
}

// generateToken generates a JWT token for the given user.
func generateToken(user sqlc.User) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{"userID": user.ID},
	)
	return token.SignedString([]byte(secret))
}

// generate password hash to store
func generateHash(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(password), bcrypt.DefaultCost,
	)
	return string(passwordHash), err
}

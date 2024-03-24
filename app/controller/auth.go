package controller

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/raitonoberu/personal-best/app/model"
	"github.com/raitonoberu/personal-best/db"
	"github.com/raitonoberu/personal-best/view"
	"golang.org/x/crypto/bcrypt"
)

var secret = os.Getenv("SECRET")

func LoginPage(c fiber.Ctx) error {
	return view.Render(c, view.Login())
}

func RegisterPage(c fiber.Ctx) error {
	return view.Render(c, view.Register())
}

func Register(c fiber.Ctx) error {
	name := strings.Clone(c.FormValue("name"))
	email := strings.Clone(c.FormValue("email"))
	password := strings.Clone(c.FormValue("password"))

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password), bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	// TODO: make sure email is unique

	user := model.User{
		Name:     name,
		Email:    email,
		Password: string(hash),
	}
	result := db.Get().Create(&user)
	if result.Error != nil {
		return result.Error
	}
	if err := setTokenCookie(c, user); err != nil {
		return err
	}
	return c.Redirect().To("/")
}

func Login(c fiber.Ctx) error {
	email := strings.Clone(c.FormValue("email"))
	password := strings.Clone(c.FormValue("password"))

	user := model.User{}

	result := db.Get().First(&user, "email = ?", email)
	if result.Error != nil {
		return result.Error
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return c.SendStatus(fiber.StatusForbidden)
	}
	if err := setTokenCookie(c, user); err != nil {
		return err
	}
	return c.Redirect().To("/")
}

func Logout(c fiber.Ctx) error {
	c.ClearCookie("token")
	return c.Redirect().To("/")
}

func setTokenCookie(c fiber.Ctx, user model.User) error {
	token, err := generateToken(user)
	if err != nil {
		return err
	}
	c.Cookie(&fiber.Cookie{
		Name:  "token",
		Value: token,
		// Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		SameSite: "strict",
	})
	return nil
}

func generateToken(user model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	})

	return token.SignedString([]byte(secret))
}

func getUser(c fiber.Ctx) *model.User {
	if user, ok := c.Locals("user").(*model.User); ok {
		return user
	}
	return nil
}

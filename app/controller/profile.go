package controller

import (
	"github.com/gofiber/fiber/v3"
	"github.com/raitonoberu/personal-best/app/model"
	"github.com/raitonoberu/personal-best/db"
	"github.com/raitonoberu/personal-best/view"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetProfile(c fiber.Ctx) error {
	userId := getUser(c).ID

	user := model.User{}

	result := db.Get().Preload("Competitions", func(tx *gorm.DB) *gorm.DB {
		return tx.Preload("Games").Preload("Players")
	}).Find(&user, userId)
	if result.Error != nil {
		return result.Error
	}
	return view.Render(c, view.Profile(user))
}

func UpdateProfile(c fiber.Ctx) error {
	updates := make(map[string]any, 3)
	if name := c.FormValue("name"); name != "" {
		updates["name"] = name
	}
	if email := c.FormValue("email"); email != "" {
		updates["email"] = email
	}
	if password := c.FormValue("password"); password != "" {
		hash, err := bcrypt.GenerateFromPassword(
			[]byte(password), bcrypt.DefaultCost,
		)
		if err != nil {
			return err
		}
		updates["password"] = string(hash)
	}

	user := model.User{}
	result := db.Get().Model(&user).
		Where("id = ?", getUser(c).ID).
		Clauses(clause.Returning{}).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	return view.Render(c, view.ProfileCard(user))
}

func DeleteProfile(c fiber.Ctx) error {
	result := db.Get().Delete(&model.User{}, getUser(c).ID)

	// TODO: delete all competitions & shit

	if result.Error != nil {
		return result.Error
	}
	c.Response().Header.Set("HX-Redirect", "/")
	return nil
}

package controller

import (
	"fmt"
	"strconv"
	"strings"

	"gorm.io/gorm/clause"

	"github.com/gofiber/fiber/v3"
	"github.com/raitonoberu/personal-best/app/model"
	"github.com/raitonoberu/personal-best/db"
	"github.com/raitonoberu/personal-best/view"
	"gorm.io/gorm"
)

func NewCompetition(c fiber.Ctx) error {
	name := strings.Clone(c.FormValue("name"))
	desc := strings.Clone(c.FormValue("description"))

	// TODO: verify

	comp := model.Competition{
		UserID:      getUser(c).ID,
		Name:        name,
		Description: desc,
	}

	result := db.Get().Create(&comp)
	if result.Error != nil {
		return result.Error
	}
	return c.Redirect().To(fmt.Sprintf("/competitions/%d", comp.ID))
}

func ListCompetitions(c fiber.Ctx) error {
	comps := []model.Competition{}

	result := db.Get().
		Preload("User").
		Preload("Games").
		Order("created_at desc").
		Find(&comps)
	if result.Error != nil {
		return result.Error
	}
	return view.Render(c, view.CompetitionList(comps))
}

func ViewCompetition(c fiber.Ctx) error {
	comp := model.Competition{}

	result := db.Get().
		Preload("User").
		Preload("Players", func(db *gorm.DB) *gorm.DB {
			return db.Order("players.created_at desc")
		}).
		Preload("Games").
		Find(&comp, c.Params("id"))
	if result.Error != nil {
		return result.Error
	}
	return view.Render(c, view.Competition(comp))
}

func UpdateCompetition(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	// TODO: verify

	updates := make(map[string]any, 3)
	if name := c.FormValue("name"); name != "" {
		updates["name"] = name
	}
	if description := c.FormValue("description"); description != "" {
		updates["description"] = description
	}
	if rounds := c.FormValue("rounds"); rounds != "" {
		// TODO: check if ready
		c.Response().Header.Set("HX-Refresh", "true")
		updates["rounds"], _ = strconv.Atoi(rounds)
	}

	comp := model.Competition{}
	result := db.Get().Model(&comp).
		Where("id = ?", id).
		Clauses(clause.Returning{}).
		Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteCompetition(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	// TODO: verify

	result := db.Get().
		Where("id = ?", id).
		Delete(&model.Competition{})
	if result.Error != nil {
		return result.Error
	}
	c.Response().Header.Set("HX-Redirect", "/")
	return nil
}

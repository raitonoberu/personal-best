package controller

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/raitonoberu/personal-best/app/model"
	"github.com/raitonoberu/personal-best/db"
	"github.com/raitonoberu/personal-best/view"
	"gorm.io/gorm/clause"
)

func AddPlayer(c fiber.Ctx) error {
	compId, _ := strconv.Atoi(c.Params("comp"))
	name := strings.Clone(c.FormValue("name"))
	isKeeper := c.FormValue("is-keeper") == "1"

	// TODO: verify

	player := model.Player{
		CompetitionID: uint(compId),
		Name:          name,
		IsKeeper:      isKeeper,
	}
	result := db.Get().Create(&player)
	if result.Error != nil {
		return result.Error
	}

	comp := model.Competition{}
	result = db.Get().
		Preload("Players").
		Find(&comp, compId)
	if result.Error != nil {
		return result.Error
	}

	return view.Render(c, view.PlayerWithUpdate(comp, player))
}

func UpdatePlayer(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	compId, _ := strconv.Atoi(c.Params("comp"))

	// TODO: verify

	updates := make(map[string]any, 2)
	if name := c.FormValue("name"); name != "" {
		updates["name"] = name
	}
	keeperField := fmt.Sprintf("is-keeper-%d", id)
	if isKeeper := c.FormValue(keeperField); isKeeper != "" {
		updates["is_keeper"] = isKeeper == "1"
	}

	player := model.Player{}
	result := db.Get().Model(&player).
		Where("id = ? AND competition_id = ?", id, compId).
		Clauses(clause.Returning{}).Updates(updates)
	if result.Error != nil {
		return result.Error
	}

	comp := model.Competition{}
	result = db.Get().
		Preload("Players").
		Find(&comp, compId)
	if result.Error != nil {
		return result.Error
	}

	return view.Render(c, view.PlayerWithUpdate(comp, player))
}

func DeletePlayer(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	compId, _ := strconv.Atoi(c.Params("comp"))

	// TODO: verify

	result := db.Get().
		Where("id = ? AND competition_id = ?", id, compId).
		Delete(&model.Player{})
	if result.Error != nil {
		return result.Error
	}

	comp := model.Competition{}
	result = db.Get().
		Preload("Players").
		Find(&comp, compId)
	if result.Error != nil {
		return result.Error
	}

	return view.Render(c, view.PlayerCount(comp))
}

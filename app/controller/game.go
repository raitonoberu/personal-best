package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/raitonoberu/personal-best/app/model"
	"github.com/raitonoberu/personal-best/db"
)

func AddGame(c fiber.Ctx) error {
	compId, _ := strconv.Atoi(c.Params("comp"))
	number, _ := strconv.Atoi(c.FormValue("number"))
	score1, _ := strconv.Atoi(c.FormValue("score1"))
	score2, _ := strconv.Atoi(c.FormValue("score2"))

	// TODO: verify

	game := model.Game{
		CompetitionID: uint(compId),
		Number:        number,
		Score1:        score1,
		Score2:        score2,
	}

	result := db.Get().Create(&game)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

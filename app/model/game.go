package model

import "gorm.io/gorm"

type Game struct {
	gorm.Model

	CompetitionID uint
	Number        int
	Score1        int
	Score2        int
}

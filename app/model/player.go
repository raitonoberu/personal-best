package model

import "gorm.io/gorm"

type Player struct {
	gorm.Model

	CompetitionID uint
	Name          string
	IsKeeper      bool
}

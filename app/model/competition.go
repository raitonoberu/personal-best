package model

import (
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type Competition struct {
	gorm.Model

	UserID uint
	User   User

	Name        string
	Description string
	Rounds      int

	Players []Player
	Games   []Game
}

func (c Competition) Started() bool {
	return c.Rounds > 0
}

func (c Competition) CanStart() bool {
	players := 0
	keepers := 0
	for _, p := range c.Players {
		if p.IsKeeper {
			keepers += 1
		} else {
			players += 1
		}
	}
	return players >= 10 && keepers >= 2
}

func (c Competition) Status() string {
	if len(c.Games) == 0 {
		return "Ещё не начался"
	}
	if len(c.Games) == c.Rounds {
		return "Завершён"
	}
	return "Идёт"
}

func (c Competition) MinRounds() int {
	players := 0
	for _, p := range c.Players {
		if !p.IsKeeper {
			players += 1
		}
	}
	return players - 5
}

func (c Competition) PlayerCount() string {
	players := 0
	keepers := 0
	for _, p := range c.Players {
		if p.IsKeeper {
			keepers += 1
		} else {
			players += 1
		}
	}
	sb := strings.Builder{}
	sb.WriteString(strconv.Itoa(players))
	sb.WriteString(" игроков")
	if players < 10 {
		sb.WriteString(" (минимум 10)")
	}
	sb.WriteString(", ")

	sb.WriteString(strconv.Itoa(keepers))
	sb.WriteString(" вратарей")
	if keepers < 2 {
		sb.WriteString(" (минимум 2)")
	}
	return sb.String()
}

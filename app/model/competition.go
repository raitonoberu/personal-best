package model

import (
	"time"

	"github.com/raitonoberu/personal-best/db/sqlc"
)

type GetCompetitionRequest struct {
	ID int64 `param:"id" validate:"required,gt=0"`
}

type GetCompetitionResponse struct {
	ID          int64           `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	StartDate   time.Time       `json:"start_date"`
	Trainer     GetUserResponse `json:"trainer"`
}

func NewGetCompetitionResponse(row sqlc.GetCompetitionRow) GetCompetitionResponse {
	return GetCompetitionResponse{
		ID:          row.Competition.ID,
		Name:        row.Competition.Name,
		Description: row.Competition.Description,
		StartDate:   row.Competition.StartDate,
		Trainer:     NewGetUserResponse(row.User),
	}
}

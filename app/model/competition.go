package model

import (
	"time"

	"github.com/raitonoberu/personal-best/db/sqlc"
)

type CreateCompetitionRequest struct {
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	StartDate   time.Time `json:"start_date" validate:"required"`
	TrainerID   int64     `json:"-"`
}

type CreateCompetitionResponse struct {
	ID int64 `json:"id"`
}

func NewCreateCompetitionResponse(c sqlc.Competition) CreateCompetitionResponse {
	return CreateCompetitionResponse{
		ID: c.ID,
	}
}

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

type ListCompetitionsRequest struct {
	Limit  int64 `query:"limit" validate:"gte=1,lte=100" default:"10"`
	Offset int64 `query:"offset" validate:"gte=0"`
}

type ListCompetitionsResponse struct {
	Count        int                      `json:"count"`
	Total        int                      `json:"total"`
	Competitions []GetCompetitionResponse `json:"competitions"`
}

func NewListCompetitionsResponse(rows []sqlc.ListCompetitionsRow) ListCompetitionsResponse {
	competitions := make([]GetCompetitionResponse, len(rows))
	for i, row := range rows {
		competitions[i] = NewGetCompetitionResponse(
			sqlc.GetCompetitionRow{
				Competition: row.Competition,
				User:        row.User,
			})
	}

	var total int
	if len(rows) > 0 {
		total = int(rows[0].Total)
	}

	return ListCompetitionsResponse{
		Count:        len(rows),
		Total:        total,
		Competitions: competitions,
	}
}

type DeleteCompetitionRequest struct {
	ID int64 `param:"id" validate:"required,gt=0"`
}

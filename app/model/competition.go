package model

import (
	"github.com/raitonoberu/personal-best/db/sqlc"
)

type CompetitionDay struct {
	Date      string `json:"date" validate:"required,date"`
	StartTime string `json:"start_time" validate:"required,time"`
	EndTime   string `json:"end_time" validate:"required,time"`
}

type CreateCompetitionRequest struct {
	UserID      int64  `json:"-"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Tours       int64  `json:"tours" validate:"required"`
	Age         int64  `json:"age" validate:"required"`
	Size        int64  `json:"size" validate:"required"`
	ClosesAt    string `json:"closes_at" validate:"required,date"`

	Days []CompetitionDay `json:"days" validate:"required,notblank,dive"`
}

type CreateCompetitionResponse struct {
	ID int64 `json:"id"`
}

type GetCompetitionRequest struct {
	ID int64 `param:"id" validate:"required"`
}

type GetCompetitionResponse struct {
	ID          int64            `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Tours       int64            `json:"tours"`
	Age         int64            `json:"age"`
	Size        int64            `json:"size"`
	ClosesAt    string           `json:"closes_at"`
	Trainer     GetUserResponse  `json:"trainer"`
	Days        []CompetitionDay `json:"days,omitempty"`
}

func NewGetCompetitionResponse(row sqlc.GetCompetitionRow, days []sqlc.CompetitionDay) GetCompetitionResponse {
	dayModels := make([]CompetitionDay, len(days))
	for i, d := range days {
		dayModels[i] = CompetitionDay{
			Date:      d.Date.Format("2006-01-02"),
			StartTime: d.StartTime.Format("15:04"),
			EndTime:   d.EndTime.Format("15:04"),
		}
	}

	return GetCompetitionResponse{
		ID:          row.Competition.ID,
		Name:        row.Competition.Name,
		Description: row.Competition.Description,
		Tours:       row.Competition.Tours,
		Age:         row.Competition.Age,
		Size:        row.Competition.Size,
		ClosesAt:    row.Competition.ClosesAt.Format("2006-01-02"),
		Trainer: GetUserResponse{
			ID:         row.User.ID,
			RoleID:     row.User.RoleID,
			FirstName:  row.User.FirstName,
			LastName:   row.User.LastName,
			MiddleName: row.User.MiddleName,
			Email:      row.User.Email,
		},
		Days: dayModels,
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

type UpdateCompetitionRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	ClosesAt    *string `json:"closes_at" validate:"omitempty,date"`
	ID          int64   `json:"-" param:"id" validate:"required"`
}

type DeleteCompetitionRequest struct {
	ID int64 `param:"id" validate:"required"`
}

type CompetitionRegistration struct {
	IsApproved bool            `json:"is_approved"`
	IsDropped  bool            `json:"is_dropped"`
	User       GetUserResponse `json:"user"`
}

type PlayerRegistration struct {
	IsApproved  bool                   `json:"is_approved"`
	IsDropped   bool                   `json:"is_dropped"`
	Competition GetCompetitionResponse `json:"competition"`
}

type ListCompetitionRegistrationsRequest struct {
	ID int64 `param:"id" validate:"required"`
}

type ListPlayerRegistrationsRequest struct {
	Limit  int64 `query:"limit" validate:"gte=1,lte=100" default:"10"`
	Offset int64 `query:"offset" validate:"gte=0"`
}

type ListPlayerRegistrationsResponse struct {
	Count         int                  `json:"count"`
	Total         int                  `json:"total"`
	Registrations []PlayerRegistration `json:"registrations"`
}

type RegisterForCompetitionRequest struct {
	ID int64 `param:"id" validate:"required"`
}

type UnregisterForCompetitionRequest struct {
	ID int64 `param:"id" validate:"required"`
}

type UpdateRegistrationRequest struct {
	IsApproved    *bool `json:"is_approved"`
	IsDropped     *bool `json:"is_dropped"`
	PlayerID      int64 `json:"-" param:"player_id" validate:"required"`
	CompetitionID int64 `json:"-" param:"comp_id" validate:"required"`
}

type CompetitionScore struct {
	WinScore  int             `json:"win_score"`
	LoseScore int             `json:"lose_score"`
	User      GetUserResponse `json:"user"`
}

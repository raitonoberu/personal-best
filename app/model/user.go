package model

import (
	"time"

	"github.com/raitonoberu/personal-best/db/sqlc"
)

type GetUserRequest struct {
	ID int64 `param:"id" validate:"required,gt=0"`
}

type GetUserResponse struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	IsTrainer bool       `json:"is_trainer"`
	BirthDate *time.Time `json:"birth_date,omitempty"`
}

func NewGetUserResponse(m sqlc.User) GetUserResponse {
	return GetUserResponse{
		ID:        m.ID,
		Name:      m.Name,
		Email:     m.Email,
		IsTrainer: m.IsTrainer,
		BirthDate: m.BirthDate,
	}
}

type ListUsersRequest struct {
	Limit  int64 `query:"limit" validate:"gte=1,lte=100" default:"10"`
	Offset int64 `query:"offset" validate:"gte=0"`
}

type ListUsersResponse struct {
	Count int               `json:"count"`
	Total int               `json:"total"`
	Users []GetUserResponse `json:"users"`
}

func NewListUsersResponse(rows []sqlc.ListUsersRow) ListUsersResponse {
	users := make([]GetUserResponse, len(rows))
	for i, row := range rows {
		users[i] = NewGetUserResponse(row.User)
	}

	var total int
	if len(rows) > 0 {
		total = int(rows[0].Total)
	}

	return ListUsersResponse{
		Count: len(rows),
		Total: total,
		Users: users,
	}
}

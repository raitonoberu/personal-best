package model

import (
	"github.com/raitonoberu/personal-best/db/sqlc"
)

type GetUserRequest struct {
	ID int64 `param:"id" validate:"required,gt=0"`
}

type PlayerResponse struct {
	BirthDate string `json:"birth_date"`
	IsMale    bool   `json:"is_male"`
	Phone     string `json:"phone"`
	Telegram  string `json:"telegram"`

	Preparation *string `json:"preparation,omitempty"`
	Position    *string `json:"position,omitempty"`
}

type GetUserResponse struct {
	ID         int64  `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	Email      string `json:"email"`

	Role   RoleResponse    `json:"role"`
	Player *PlayerResponse `json:"player,omitempty"`
}

func NewGetUserResponse(u sqlc.User, p sqlc.Player, r sqlc.Role) GetUserResponse {
	var player *PlayerResponse
	if p.UserID != 0 {
		player = &PlayerResponse{
			BirthDate:   p.BirthDate.Format("2006-01-02"),
			IsMale:      p.IsMale,
			Phone:       p.Phone,
			Telegram:    p.Telegram,
			Preparation: p.Preparation,
			Position:    p.Position,
		}
	}

	return GetUserResponse{
		ID:         u.ID,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		MiddleName: u.MiddleName,
		Email:      u.Email,
		Role:       RoleResponse(r),
		Player:     player,
	}
}

type UpdateUserRequest struct {
	FirstName  *string `json:"first_name"`
	LastName   *string `json:"last_name"`
	MiddleName *string `json:"middle_name"`
	Email      *string `json:"email"`
	Password   *string `json:"password"`
	ID         int64   `json:"-"`
}

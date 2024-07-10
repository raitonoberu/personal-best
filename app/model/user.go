package model

import (
	"github.com/raitonoberu/personal-best/db/sqlc"
)

type GetUserRequest struct {
	ID int64 `param:"id" validate:"required"`
}

type PlayerResponse struct {
	BirthDate string `json:"birth_date"`
	IsMale    bool   `json:"is_male"`
	Phone     string `json:"phone"`
	Telegram  string `json:"telegram"`

	Preparation *string `json:"preparation,omitempty"`
	Position    *string `json:"position,omitempty"`
}

type ListUsersResponse struct {
	Count int               `json:"count"`
	Total int               `json:"total"`
	Users []GetUserResponse `json:"users"`
}

type GetUserResponse struct {
	ID         int64  `json:"id"`
	RoleID     int64  `json:"role_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	Email      string `json:"email"`

	Player *PlayerResponse `json:"player,omitempty"`
}

func NewGetUserResponse(row sqlc.GetUserRow) GetUserResponse {
	var player *PlayerResponse
	if row.UserPlayer.UserID != nil {
		player = &PlayerResponse{
			BirthDate:   (*row.UserPlayer.BirthDate).Format("2006-01-02"),
			IsMale:      *row.UserPlayer.IsMale,
			Phone:       *row.UserPlayer.Phone,
			Telegram:    *row.UserPlayer.Telegram,
			Preparation: row.UserPlayer.Preparation,
			Position:    row.UserPlayer.Position,
		}
	}

	return GetUserResponse{
		ID:         row.User.ID,
		RoleID:     row.User.RoleID,
		FirstName:  row.User.FirstName,
		LastName:   row.User.LastName,
		MiddleName: row.User.MiddleName,
		Email:      row.User.Email,
		Player:     player,
	}
}

func NewGetPlayerResponse(u sqlc.User, p sqlc.Player) GetUserResponse {
	var player *PlayerResponse
	if p.UserID != 0 {
		player = &PlayerResponse{
			BirthDate:   (p.BirthDate).Format("2006-01-02"),
			IsMale:      p.IsMale,
			Phone:       p.Phone,
			Telegram:    p.Telegram,
			Preparation: p.Preparation,
			Position:    p.Position,
		}
	}

	return GetUserResponse{
		ID:         u.ID,
		RoleID:     u.RoleID,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		MiddleName: u.MiddleName,
		Email:      u.Email,
		Player:     player,
	}
}

type UpdateUserRequest struct {
	ID         int64   `json:"-"`
	FirstName  *string `json:"first_name"`
	LastName   *string `json:"last_name"`
	MiddleName *string `json:"middle_name"`
	Email      *string `json:"email" validate:"omitempty,email"`
	Password   *string `json:"password"`

	BirthDate *string `json:"birth_date" validate:"omitempty,date"`
	IsMale    *bool   `json:"is_male"`
	Phone     *string `json:"phone" validate:"omitempty,e164"`
	Telegram  *string `json:"telegram" validate:"omitempty,startswith=@"`
}

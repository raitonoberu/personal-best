package model

import (
	"github.com/raitonoberu/personal-best/db/sqlc"
)

type GetUserRequest struct {
	ID int64 `param:"id" validate:"required,gt=0"`
}

type Player struct {
	BirthDate  string `json:"birth_date"`
	IsMale     bool   `json:"is_male"`
	Phone      string `json:"phone"`
	Telegram   string `json:"telegram"`
	IsVerified bool   `json:"is_verified"`

	Preparation *string `json:"preparation,omitempty"`
	Position    *string `json:"position,omitempty"`
}

type GetUserResponse struct {
	ID         int64  `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	Email      string `json:"email"`

	Player *Player `json:"player,omitempty"`
}

func NewGetUserResponse(u sqlc.User, p sqlc.Player) GetUserResponse {
	var player *Player
	if p.UserID != 0 {
		player = &Player{
			BirthDate:   p.BirthDate.Format("2006-01-02"),
			IsMale:      p.IsMale,
			Phone:       p.Phone,
			Telegram:    p.Telegram,
			IsVerified:  p.IsVerified,
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
		Player:     player,
	}
}

// type ListUsersRequest struct {
// 	Limit  int64 `query:"limit" validate:"gte=1,lte=100" default:"10"`
// 	Offset int64 `query:"offset" validate:"gte=0"`
// }
//
// type ListUsersResponse struct {
// 	Count int               `json:"count"`
// 	Total int               `json:"total"`
// 	Users []GetUserResponse `json:"users"`
// }
//
// func NewListUsersResponse(rows []sqlc.ListUsersRow) ListUsersResponse {
// 	users := make([]GetUserResponse, len(rows))
// 	for i, row := range rows {
// 		users[i] = NewGetUserResponse(row.User)
// 	}
//
// 	var total int
// 	if len(rows) > 0 {
// 		total = int(rows[0].Total)
// 	}
//
// 	return ListUsersResponse{
// 		Count: len(rows),
// 		Total: total,
// 		Users: users,
// 	}
// }

type UpdateUserRequest struct {
	FirstName  *string `json:"first_name"`
	LastName   *string `json:"last_name"`
	MiddleName *string `json:"middle_name"`
	Email      *string `json:"email"`
	Password   *string `json:"password"`
	ID         int64   `json:"-"`
}

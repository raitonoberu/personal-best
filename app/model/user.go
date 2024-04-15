package model

import (
	"github.com/raitonoberu/personal-best/db/sqlc"
)

type GetUserRequest struct {
	ID int64 `param:"id" validate:"required,gt=0"`
}

type GetUserResponse struct {
	ID         int64  `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	Email      string `json:"email"`
}

func NewGetUserResponse(m sqlc.User) GetUserResponse {
	return GetUserResponse{
		ID:         m.ID,
		FirstName:  m.FirstName,
		LastName:   m.LastName,
		MiddleName: m.MiddleName,
		Email:      m.Email,
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

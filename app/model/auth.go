package model

import "time"

type RegisterRequest struct {
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required"`
	FirstName  string `json:"first_name" validate:"required"`
	LastName   string `json:"last_name" validate:"required"`
	MiddleName string `json:"middle_name" validate:"required"`

	BirthDate time.Time `json:"birth_date" validate:"required"`
	IsMale    bool      `json:"is_male" validate:"required"`
	Phone     string    `json:"phone" validate:"required,e164"`
	Telegram  string    `json:"telegrem" validate:"required,startswith=@"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	ID    int64  `json:"id"`
	Token string `json:"token"`
}

func NewAuthResponse(id int64, token string) AuthResponse {
	return AuthResponse{
		ID:    id,
		Token: token,
	}
}

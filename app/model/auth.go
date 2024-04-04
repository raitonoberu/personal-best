package model

import "time"

type RegisterRequest struct {
	Name      string     `json:"name" validate:"required"`
	Email     string     `json:"email" validate:"required,email"`
	Password  string     `json:"password" validate:"required"`
	IsTrainer bool       `json:"is_trainer"`
	BirthDate *time.Time `json:"birth_date" validate:"required_if=IsTrainer false,excluded_if=IsTrainer true"`
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

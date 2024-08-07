package model

type RegisterRequest struct {
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required"`
	FirstName  string `json:"first_name" validate:"required"`
	LastName   string `json:"last_name" validate:"required"`
	MiddleName string `json:"middle_name" validate:"required"`

	BirthDate   string `json:"birth_date" validate:"required,date"`
	IsMale      *bool  `json:"is_male" validate:"required"`
	Phone       string `json:"phone" validate:"required,e164"`
	Telegram    string `json:"telegram" validate:"required,startswith=@"` // TODO: more validations
	Preparation string `json:"preparation"`                               // TODO: limit to options
	Position    string `json:"position"`                                  // TODO: limit to options
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	ID    int64  `json:"id"`
	Token string `json:"token"`
}

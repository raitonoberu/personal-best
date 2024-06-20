package model

type AdminCreateUserRequest struct {
	RoleID     int64  `json:"role_id" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required"`
	FirstName  string `json:"first_name" validate:"required"`
	LastName   string `json:"last_name" validate:"required"`
	MiddleName string `json:"middle_name" validate:"required"`

	BirthDate *string `json:"birth_date" validate:"omitempty,date"`
	IsMale    *bool   `json:"is_male"`
	Phone     *string `json:"phone" validate:"omitempty,e164"`
	Telegram  *string `json:"telegram" validate:"omitempty,startswith=@"`
}

type AdminListUsersRequest struct {
	Limit  int64 `query:"limit" validate:"gte=1,lte=100" default:"10"`
	Offset int64 `query:"offset" validate:"gte=0"`
	RoleID int64 `query:"role_id" json:"role_id" validate:"gte=1" default:"3"`
}

type AdminUpdateUserRequest struct {
	ID         int64   `json:"-" param:"id" validate:"required"`
	RoleID     *int64  `json:"role_id"`
	Email      *string `json:"email" validate:"omitempty,email"`
	Password   *string `json:"password"`
	FirstName  *string `json:"first_name"`
	LastName   *string `json:"last_name"`
	MiddleName *string `json:"middle_name"`

	BirthDate *string `json:"birth_date" validate:"omitempty,date"`
	IsMale    *bool   `json:"is_male"`
	Phone     *string `json:"phone" validate:"omitempty,e164"`
	Telegram  *string `json:"telegram" validate:"omitempty,startswith=@"`
}

type AdminDeleteUserRequest struct {
	ID int64 `param:"id" validate:"required"`
}

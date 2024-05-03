package model

type RoleResponse struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	CanView        bool   `json:"can_view"`
	CanParticipate bool   `json:"can_participate"`
	CanCreate      bool   `json:"can_create"`
	IsFree         bool   `json:"is_free"`
	IsAdmin        bool   `json:"is_admin"`
}

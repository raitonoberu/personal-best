package service

import (
	"context"

	"github.com/raitonoberu/personal-best/app/model"
	"github.com/raitonoberu/personal-best/db/sqlc"
)

func (s Service) GetUserRole(ctx context.Context, id int64) (*sqlc.Role, error) {
	role, err := s.queries.GetUserRole(ctx, id)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (s Service) ListRoles(ctx context.Context) ([]model.RoleResponse, error) {
	roles, err := s.queries.ListRoles(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]model.RoleResponse, len(roles))
	for i, r := range roles {
		resp[i] = model.RoleResponse(r)
	}
	return resp, err
}

package service

import (
	"context"

	"github.com/raitonoberu/personal-best/app/model"
	"github.com/raitonoberu/personal-best/db/sqlc"
)

func (s Service) ListUsers(ctx context.Context, roleID, limit, offset int64) (*model.ListUsersResponse, error) {
	rows, err := s.queries.ListUsers(ctx, sqlc.ListUsersParams{
		RoleID: roleID,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	items := make([]model.GetUserResponse, len(rows))
	for i, r := range rows {
		items[i] = model.NewGetUserResponse(sqlc.GetUserRow{
			User:       r.User,
			UserPlayer: r.UserPlayer,
		})
	}

	var total int
	if len(items) != 0 {
		total = int(rows[0].Total)
	}

	return &model.ListUsersResponse{
		Count: len(items),
		Total: total,
		Users: items,
	}, nil
}

package service

import (
	"context"

	"github.com/raitonoberu/personal-best/app/model"
	"github.com/raitonoberu/personal-best/db/sqlc"
)

func (s Service) UpdateRegistration(ctx context.Context, req model.UpdateRegistrationRequest) error {
	err := s.queries.UpdateRegistration(ctx, sqlc.UpdateRegistrationParams(req))
	if err != nil {
		return err
	}
	err = s.UpdateMatches(ctx, req.CompetitionID)
	if err != nil && err != ErrNotEnoughPlayers {
		return err
	}
	return nil
}

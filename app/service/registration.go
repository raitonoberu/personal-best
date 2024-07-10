package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raitonoberu/personal-best/app/model"
	"github.com/raitonoberu/personal-best/db/sqlc"
)

var (
	ErrCompetitionNotFound = echo.NewHTTPError(404, "Соревнование не найдено")
	ErrCompetitionClosed   = echo.NewHTTPError(400, "Запись на сорвенование закрыта")
	ErrPlayerDropped       = echo.NewHTTPError(400, "Вы исключены из соревнования")
)

func (s Service) ListCompetitionRegistrations(ctx context.Context, id int64) ([]model.CompetitionRegistration, error) {
	rows, err := s.queries.ListCompetitionRegistrations(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrCompetitionNotFound
		}
		return nil, err
	}
	regs := make([]model.CompetitionRegistration, len(rows))
	for i := 0; i < len(rows); i++ {
		regs[i] = model.CompetitionRegistration{
			IsApproved: rows[i].Registration.IsApproved,
			IsDropped:  rows[i].Registration.IsDropped,
			User:       model.NewGetPlayerResponse(rows[i].User, rows[i].Player),
		}
	}
	return regs, nil
}

func (s Service) ListPlayerRegistrations(ctx context.Context, id, limit, offset int64) (*model.ListPlayerRegistrationsResponse, error) {
	rows, err := s.queries.ListPlayerRegistrations(ctx,
		sqlc.ListPlayerRegistrationsParams{
			PlayerID: id,
			Limit:    limit,
			Offset:   offset,
		})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrCompetitionNotFound
		}
		return nil, err
	}

	regs := make([]model.PlayerRegistration, len(rows))
	for i := 0; i < len(rows); i++ {
		regs[i] = model.PlayerRegistration{
			IsApproved: rows[i].Registration.IsApproved,
			IsDropped:  rows[i].Registration.IsDropped,
			Competition: model.NewGetCompetitionResponse(sqlc.GetCompetitionRow{
				Competition: rows[i].Competition,
				User:        rows[i].User,
			}, nil),
		}
	}
	var total int
	if len(rows) > 0 {
		total = int(rows[0].Total)
	}

	return &model.ListPlayerRegistrationsResponse{
		Count:         len(rows),
		Total:         total,
		Registrations: regs,
	}, nil
}

func (s Service) RegisterForCompetition(ctx context.Context, userId, compId int64) error {
	competition, err := s.queries.GetCompetition(ctx, compId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrCompetitionNotFound
		}
		return err
	}

	// here we add 24h cuz stored ClosesAt is the last day available to register
	if competition.Competition.ClosesAt.Add(time.Hour * 24).Before(time.Now()) {
		return ErrCompetitionClosed
	}

	if err := s.queries.CreateRegistration(ctx,
		sqlc.CreateRegistrationParams{
			CompetitionID: compId,
			PlayerID:      userId,
			IsApproved:    false,
			IsDropped:     false,
		}); err != nil {
		return err
	}

	if err := s.UpdateMatches(ctx, compId); err != nil && err != ErrNotEnoughPlayers {
		// ignoring if not enough players
		return err
	}
	return nil
}

func (s Service) UnregisterForCompetition(ctx context.Context, userId, compId int64) error {
	competition, err := s.queries.GetCompetition(ctx, compId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrCompetitionNotFound
		}
		return err
	}

	// here we add 24h cuz stored ClosesAt is the last day available to register
	if competition.Competition.ClosesAt.Add(time.Hour * 24).Before(time.Now()) {
		return ErrCompetitionClosed
	}

	reg, err := s.queries.GetRegistration(ctx,
		sqlc.GetRegistrationParams{
			CompetitionID: compId,
			PlayerID:      userId,
		})
	if err != nil {
		return err
	}
	if reg.IsDropped {
		return ErrPlayerDropped
	}

	if err := s.queries.DeleteRegistration(ctx,
		sqlc.DeleteRegistrationParams{
			CompetitionID: compId,
			PlayerID:      userId,
		}); err != nil {
		return err
	}

	if err := s.UpdateMatches(ctx, compId); err != nil &&
		err != ErrNotEnoughPlayers {
		// ignoring if not enough players
		// TODO: Here we actually should do something xD
		return err
	}
	return nil
}

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

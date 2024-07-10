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
	ErrStartBeforeClose = echo.NewHTTPError(400, "Соревнование не может начинаться до конца регистрации")
	ErrEndBeforeStart   = echo.NewHTTPError(400, "Время начала должно быть раньше времени конца")
)

func (s Service) CreateCompetition(ctx context.Context, req model.CreateCompetitionRequest) (int64, error) {
	closesAt := parseDate(req.ClosesAt)
	for _, d := range req.Days {
		day := parseDate(d.Date)
		if day.Equal(closesAt) || day.Before(closesAt) {
			return 0, ErrStartBeforeClose
		}

		startTime := parseTime(d.StartTime, day)
		endTime := parseTime(d.EndTime, day)

		if startTime.After(endTime) {
			return 0, ErrEndBeforeStart
		}
	}

	// TODO: here we need MUCH MORE CHECKS

	tx, err := s.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	qtx := s.queries.WithTx(tx)

	comp, err := qtx.CreateCompetition(ctx,
		sqlc.CreateCompetitionParams{
			TrainerID:   req.UserID,
			Name:        req.Name,
			Description: req.Description,
			Tours:       req.Tours,
			Age:         req.Age,
			Size:        req.Size,
			ClosesAt:    closesAt,
		})
	if err != nil {
		return 0, err
	}

	for _, d := range req.Days {
		date := parseDate(d.Date)
		_, err := qtx.CreateCompetitionDay(ctx,
			sqlc.CreateCompetitionDayParams{
				CompetitionID: comp.ID,
				Date:          date,
				StartTime:     parseTime(d.StartTime, date),
				EndTime:       parseTime(d.EndTime, date),
			})
		if err != nil {
			return 0, err
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	// TODO: make it in the SAME TRANSACTION
	return comp.ID, s.GenerateMatches(ctx, comp.ID)
}

func (s Service) GetCompetition(ctx context.Context, id int64) (*model.GetCompetitionResponse, error) {
	competition, err := s.queries.GetCompetition(ctx, id)
	if err != nil {
		return nil, err
	}

	days, err := s.queries.GetCompetitionDays(ctx, id)
	if err != nil {
		return nil, err
	}
	c := model.NewGetCompetitionResponse(competition, days)
	return &c, nil
}

func (s Service) ListCompetitions(ctx context.Context, limit, offset int64) (*model.ListCompetitionsResponse, error) {
	rows, err := s.queries.ListCompetitions(ctx,
		sqlc.ListCompetitionsParams{
			Limit:  limit,
			Offset: offset,
		})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrCompetitionNotFound
		}
		return nil, err
	}

	competitions := make([]model.GetCompetitionResponse, len(rows))
	for i, row := range rows {
		competitions[i] = model.NewGetCompetitionResponse(
			sqlc.GetCompetitionRow{
				Competition: row.Competition,
				User:        row.User,
			}, nil) // TODO:
	}

	var total int
	if len(rows) > 0 {
		total = int(rows[0].Total)
	}

	return &model.ListCompetitionsResponse{
		Count:        len(rows),
		Total:        total,
		Competitions: competitions,
	}, nil
}

func (s Service) UpdateCompetition(ctx context.Context, req model.UpdateCompetitionRequest) error {
	var closesAt *time.Time
	if req.ClosesAt != nil {
		date := parseDate(*req.ClosesAt)
		closesAt = &date

		// TODO: check if it's not after first match and not in past
	}

	return s.queries.UpdateCompetition(ctx,
		sqlc.UpdateCompetitionParams{
			ID:          req.ID,
			Name:        req.Name,
			Description: req.Description,
			ClosesAt:    closesAt,
		})
}

func (s Service) DeleteCompetition(ctx context.Context, id int64) error {
	return s.queries.DeleteCompetition(ctx, id)
}

func (s Service) GetScores(ctx context.Context, compID int64) ([]model.CompetitionScore, error) {
	players, err := s.queries.ListApprovedCompetitionPlayers(ctx, compID)
	if err != nil {
		return nil, err
	}

	scores := make([]model.CompetitionScore, len(players))
	for i, p := range players {
		score, err := s.queries.GetMatchPlayerLastScores(ctx, sqlc.GetMatchPlayerLastScoresParams{
			PlayerID:      p.User.ID,
			CompetitionID: compID,
		})
		if err != nil {
			scores[i] = model.CompetitionScore{
				WinScore:  0,
				LoseScore: 0,
				User:      model.NewGetPlayerResponse(p.User, p.Player),
			}
			continue
		}

		scores[i] = model.CompetitionScore{
			User:      model.NewGetPlayerResponse(p.User, p.Player),
			WinScore:  int(*score.WinScore),
			LoseScore: int(*score.LoseScore),
		}
	}
	return scores, nil
}

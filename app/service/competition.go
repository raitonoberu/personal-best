package service

import (
	"context"

	"github.com/raitonoberu/personal-best/app/model"
	"github.com/raitonoberu/personal-best/db/sqlc"
)

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

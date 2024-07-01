package service

import (
	"context"
	"database/sql"
	"log"
	"slices"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raitonoberu/personal-best/app/model"
	"github.com/raitonoberu/personal-best/db/sqlc"
)

const (
	matchDuration = time.Minute
	pauseDuration = time.Second * 30
)

var (
	ErrNotEnoughTime    = echo.NewHTTPError(400, "Выделено слишком мало времени")
	ErrNotEnoughPlayers = echo.NewHTTPError(400, "Недостаточно игроков")
	ErrCantChangeScore  = echo.NewHTTPError(400, "Нельзя изменить счёт")
)

// GenerateMatches creates matches based on competition days.
// Generated matches have null scored.
func (s Service) GenerateMatches(ctx context.Context, compID int64) error {
	days, err := s.queries.GetCompetitionDays(ctx, compID)
	if err != nil {
		return err
	}

	matches := []time.Time{}
	for _, d := range days {
		t := d.StartTime
		for t.Before(d.EndTime) {
			matches = append(matches, t)
			t = t.Add(matchDuration).Add(pauseDuration)
		}
	}

	if len(matches) == 0 {
		return ErrNotEnoughTime
	}

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	qtx := s.queries.WithTx(tx)

	for _, t := range matches {
		_, err := qtx.CreateMatch(ctx, sqlc.CreateMatchParams{
			CompetitionID: compID,
			StartTime:     t,
		})
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

// UpdateMatches generates teams for next (first that has not started) match
func (s Service) UpdateMatches(ctx context.Context, compID int64) error {
	comp, err := s.queries.GetCompetition(ctx, compID)
	if err != nil {
		return err
	}
	matches, err := s.queries.ListAllMatches(ctx, compID)
	if err != nil {
		return err
	}

	for i, m := range matches {
		if m.LeftScore == nil || m.RightScore == nil {
			return s.updateMatch(ctx, comp.Competition, matches, i)
		}
	}
	// competition finished, nothing to update
	return nil
}

func (s Service) updateMatch(ctx context.Context, comp sqlc.Competition, matches []sqlc.Match, index int) error {
	players, err := s.queries.ListCompetitionPlayers(ctx, comp.ID)
	if err != nil {
		return err
	}

	if len(players) < int(comp.Size)*2 {
		return ErrNotEnoughPlayers
	}

	// add some nice rng
	// rand.Shuffle(len(players), func(i, j int) { players[i], players[j] = players[j], players[i] })

	matchID := matches[index].ID

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	qtx := s.queries.WithTx(tx)

	// step 1. delete all existing match players
	if err := qtx.DeleteMatchPlayers(ctx, matchID); err != nil {
		return err
	}

	// step 2. get all previously played players
	prevMatchIDs := make([]int64, index)
	for i := 0; i < index; i++ {
		prevMatchIDs[i] = matches[i].ID
	}
	var prevMatchPlayers []sqlc.MatchPlayer
	if index != 0 {
		prevMatchPlayers, err = qtx.ListMatchPlayersBatch(ctx, prevMatchIDs)
		if err != nil {
			return err
		}
	}

	// step 3. find guys that are most likely to play next game
	playerScore := map[int64]int{} // players with smallest score will play
	// dont prefer those who plays a lot
	for _, p := range prevMatchPlayers {
		playerScore[p.PlayerID] += 1
	}
	// dont prefer those who played last match
	for _, p := range prevMatchPlayers {
		if p.MatchID == matches[index-1].ID {
			playerScore[p.PlayerID] += 1
		}
	}
	// more checks?

	scorePlayers := map[int][]int64{}
	for _, p := range players {
		scorePlayers[playerScore[p.User.ID]] = append(scorePlayers[playerScore[p.User.ID]], p.User.ID)
	}
	scores := make([]int, len(scorePlayers))
	i := 0
	for s := range scorePlayers {
		scores[i] = s
		i++
	}
	slices.Sort(scores)

	playerIDs := make([]int64, comp.Size*2)
	scoresIndex, playersIndex := 0, 0
	for i := 0; i < len(playerIDs); i++ {
		// WARN: this may explode :)
		if playersIndex == len(scorePlayers[scores[scoresIndex]]) {
			scoresIndex += 1
			playersIndex = 0
		}
		playerIDs[i] = scorePlayers[scores[scoresIndex]][playersIndex]
		playersIndex += 1
	}

	// step 4. save match players
	// TODO: sort em based on their rating
	for i := int64(0); i < comp.Size*2; i++ {
		team := false
		if i >= comp.Size {
			team = true
		}
		err := qtx.CreateMatchPlayer(ctx, sqlc.CreateMatchPlayerParams{
			MatchID:  matchID,
			PlayerID: playerIDs[i],
			Team:     team,
			Position: "Игрок", // TODO:
		})
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

// SetMatchScore saves the score and makes it possible to start next match
func (s Service) SetMatchScore(ctx context.Context, matchID, leftScore, rightScore int64) error {
	match, err := s.queries.GetMatch(ctx, matchID)
	if err != nil {
		return err
	}
	if match.RightScore != nil && match.LeftScore != nil {
		// score already set (past match)
		return ErrCantChangeScore
	}

	matchPlayers, err := s.queries.ListMatchPlayers(ctx, matchID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if len(matchPlayers) == 0 {
		// no players (future match)
		return ErrCantChangeScore
	}

	err = s.queries.UpdateMatchScore(ctx, sqlc.UpdateMatchScoreParams{
		ID:         matchID,
		LeftScore:  &leftScore,
		RightScore: &rightScore,
	})
	if err != nil {
		return err
	}

	if err := s.UpdatePlayerScores(ctx, match.CompetitionID); err != nil {
		return err
	}

	err = s.UpdateMatches(ctx, match.CompetitionID)
	if err != nil && err != ErrNotEnoughPlayers {
		return err
	}
	return nil
}

// UpdatePlayerScores updates win/lose scores for all players in last match
func (s Service) UpdatePlayerScores(ctx context.Context, compID int64) error {
	match, err := s.queries.GetLastMatch(ctx, compID)
	if err != nil {
		return err
	}
	rightScore, leftScore := *match.RightScore, *match.LeftScore

	players, err := s.queries.ListMatchPlayers(ctx, match.ID)
	if err != nil {
		return err
	}

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	qtx := s.queries.WithTx(tx)

	for _, p := range players {
		lastMatchPlayers, err := qtx.GetMatchPlayersToUpdateScore(ctx, sqlc.GetMatchPlayersToUpdateScoreParams{
			PlayerID:      p.PlayerID,
			CompetitionID: compID,
		})
		if err != nil {
			return err
		}

		if len(lastMatchPlayers) == 0 {
			log.Printf("couldn't find last matches for player %d, comp %d", p.PlayerID, compID)
			continue
		}

		// not sure about it
		var win, lose int64
		if p.Team && rightScore > leftScore {
			win += rightScore
			lose += leftScore
		} else {
			win += leftScore
			lose += rightScore
		}

		if len(lastMatchPlayers) == 2 && lastMatchPlayers[1].WinScore != nil {
			win += *lastMatchPlayers[1].WinScore
			lose += *lastMatchPlayers[1].LoseScore
		}

		qtx.SetMatchPlayerWinLoseScores(ctx, sqlc.SetMatchPlayerWinLoseScoresParams{
			PlayerID:  p.PlayerID,
			MatchID:   p.MatchID,
			WinScore:  &win,
			LoseScore: &lose,
		})
	}

	return tx.Commit()
}

func (s Service) ListMatches(ctx context.Context, req model.ListMatchesRequest) (*model.ListMatchesResponse, error) {
	matches, err := s.queries.ListMatches(ctx, sqlc.ListMatchesParams{
		CompetitionID: req.ID,
		Limit:         req.Limit,
		Offset:        req.Offset,
	})
	if err != nil {
		return nil, err
	}

	matchIDs := make([]int64, len(matches))
	for i := 0; i < len(matches); i++ {
		matchIDs[i] = matches[i].Match.ID
	}

	matchPlayerRows, err := s.queries.ListMatchPlayersWithPlayersBatch(ctx, matchIDs)
	if err != nil {
		return nil, err
	}

	matchesMap := map[int64]model.Match{}
	for _, r := range matches {
		m := r.Match
		matchesMap[m.ID] = model.Match{
			ID:         m.ID,
			StartTime:  m.StartTime,
			LeftScore:  m.LeftScore,
			RightScore: m.RightScore,
			LeftTeam:   []model.MatchPlayer{},
			RightTeam:  []model.MatchPlayer{},
		}
	}

	for _, r := range matchPlayerRows {
		m := matchesMap[r.MatchPlayer.MatchID]
		p := model.MatchPlayer{
			ID:        r.User.ID,
			Name:      r.User.FirstName + " " + r.User.LastName,
			WinScore:  r.MatchPlayer.WinScore,
			LoseScore: r.MatchPlayer.LoseScore,
		}
		if r.MatchPlayer.Team == false {
			m.LeftTeam = append(m.LeftTeam, p)
		} else {
			m.RightTeam = append(m.RightTeam, p)
		}
		matchesMap[r.MatchPlayer.MatchID] = m
	}

	items := make([]model.Match, len(matches))
	for i, m := range matches {
		items[i] = matchesMap[m.Match.ID]
	}

	var total int
	if len(items) != 0 {
		total = int(matches[0].Total)
	}

	return &model.ListMatchesResponse{
		Count:   len(items),
		Total:   total,
		Matches: items,
	}, nil
}

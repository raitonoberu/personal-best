package model

import "time"

type ListMatchesRequest struct {
	ID     int64 `json:"-" param:"id" validate:"required"`
	Limit  int64 `query:"limit" validate:"gte=1,lte=100" default:"10"`
	Offset int64 `query:"offset" validate:"gte=0"`
}

type ListMatchesResponse struct {
	Count   int     `json:"count"`
	Total   int     `json:"total"`
	Matches []Match `json:"matches"`
}

type Match struct {
	ID         int64         `json:"id"`
	StartTime  time.Time     `json:"start_time"`
	LeftScore  *int64        `json:"left_score"`
	RightScore *int64        `json:"right_score"`
	LeftTeam   []MatchPlayer `json:"left_team"`
	RightTeam  []MatchPlayer `json:"right_team"`
}

type MatchPlayer struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	WinScore  *int64 `json:"win_score"`
	LoseScore *int64 `json:"lose_score"`
}

type UpdateMatchRequest struct {
	ID            int64 `json:"-" param:"id" validate:"required"`
	CompetitionID int64 `json:"-" param:"comp_id" validate:"required"`
	LeftScore     int64 `json:"left_score"`
	RightScore    int64 `json:"right_score"`
}

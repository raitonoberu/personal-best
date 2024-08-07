// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package sqlc

import (
	"time"
)

type Competition struct {
	ID          int64
	TrainerID   int64
	Name        string
	Description string
	Tours       int64
	Age         int64
	Size        int64
	ClosesAt    time.Time
	CreatedAt   time.Time
}

type CompetitionDay struct {
	CompetitionID int64
	Date          time.Time
	StartTime     time.Time
	EndTime       time.Time
}

type Document struct {
	ID        int64
	PlayerID  int64
	Name      string
	Url       string
	ExpiresAt time.Time
	CreatedAt time.Time
}

type Match struct {
	ID            int64
	CompetitionID int64
	StartTime     time.Time
	LeftScore     *int64
	RightScore    *int64
}

type MatchPlayer struct {
	MatchID   int64
	PlayerID  int64
	Position  string
	Team      bool
	WinScore  *int64
	LoseScore *int64
}

type Player struct {
	UserID      int64
	BirthDate   time.Time
	IsMale      bool
	Phone       string
	Telegram    string
	Preparation string
	Position    string
}

type Registration struct {
	CompetitionID int64
	PlayerID      int64
	IsApproved    bool
	IsDropped     bool
	CreatedAt     time.Time
}

type Role struct {
	ID             int64
	Name           string
	CanView        bool
	CanParticipate bool
	CanCreate      bool
	IsFree         bool
	IsAdmin        bool
}

type User struct {
	ID         int64
	RoleID     int64
	Email      string
	Password   string
	FirstName  string
	LastName   string
	MiddleName string
	CreatedAt  time.Time
}

type UserPlayer struct {
	UserID      *int64
	BirthDate   *time.Time
	IsMale      *bool
	Phone       *string
	Telegram    *string
	Preparation *string
	Position    *string
}

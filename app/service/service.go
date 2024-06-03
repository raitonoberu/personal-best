package service

import (
	"database/sql"

	"github.com/raitonoberu/personal-best/db/sqlc"
)

type Service struct {
	db      *sql.DB
	queries *sqlc.Queries
}

func New(db *sql.DB) Service {
	queries := sqlc.New(db)

	return Service{
		db:      db,
		queries: queries,
	}
}

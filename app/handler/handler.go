package handler

import (
	"database/sql"

	"github.com/raitonoberu/personal-best/db/sqlc"
)

type Handler struct {
	db      *sql.DB
	queries *sqlc.Queries
}

func New(db *sql.DB) Handler {
	queries := sqlc.New(db)

	return Handler{
		db:      db,
		queries: queries,
	}
}

package handler

import "github.com/raitonoberu/personal-best/db/sqlc"

type Handler struct {
	db *sqlc.Queries
}

func New(db *sqlc.Queries) Handler {
	return Handler{
		db: db,
	}
}

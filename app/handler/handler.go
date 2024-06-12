package handler

import (
	"database/sql"

	"github.com/raitonoberu/personal-best/app/service"
	"github.com/raitonoberu/personal-best/db/sqlc"
)

type Handler struct {
	db      *sql.DB
	queries *sqlc.Queries
	service *service.Service
}

func New(db *sql.DB, service *service.Service) Handler {
	queries := sqlc.New(db)

	return Handler{
		db:      db,
		queries: queries,
		service: service,
	}
}

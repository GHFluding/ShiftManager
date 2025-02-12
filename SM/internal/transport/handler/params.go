package handler

import (
	"log/slog"
	"sm/internal/database/postgres"
	"time"
)

type Params struct {
	db  *postgres.Queries
	log *slog.Logger
}

func CreateParams(db *postgres.Queries, log *slog.Logger) Params {
	return Params{
		db:  db,
		log: log,
	}
}

type requestParams struct {
	startTime time.Time
	requestId string
}

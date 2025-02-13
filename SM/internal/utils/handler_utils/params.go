package handler_utils

import (
	"log/slog"
	"sm/internal/database/postgres"
	"time"
)

// Params stores parameters for handlers
type Params struct {
	DB  *postgres.Queries
	Log *slog.Logger
}

func CreateParams(db *postgres.Queries, log *slog.Logger) Params {
	return Params{
		DB:  db,
		Log: log,
	}
}

// Params for logger
type RequestParams struct {
	StartTime time.Time
	RequestId string
}

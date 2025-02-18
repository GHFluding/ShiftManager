package services

import (
	"log/slog"
	"sm/internal/database/postgres"
)

type ServicesParams struct {
	db  *postgres.Queries
	log *slog.Logger
}

func NewServicesParams(db *postgres.Queries, log *slog.Logger) *ServicesParams {
	return &ServicesParams{db: db, log: log}
}

package services

import (
	"log/slog"

	"github.com/GHFluding/ShiftManager/SM/internal/database/postgres"
)

type ServicesParams struct {
	db  *postgres.Queries
	log *slog.Logger
}

func NewServicesParams(db *postgres.Queries, log *slog.Logger) *ServicesParams {
	return &ServicesParams{db: db, log: log}
}

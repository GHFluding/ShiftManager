package grpc_storage

import (
	"log/slog"

	"github.com/GHFluding/ShiftManager/SM/internal/database/postgres"
)

type Services struct {
	db  *postgres.Queries
	log *slog.Logger
}

func New(db *postgres.Queries, log *slog.Logger) *Services {
	return &Services{
		db:  db,
		log: log,
	}
}

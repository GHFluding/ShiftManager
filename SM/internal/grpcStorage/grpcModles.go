package models_grpc

import (
	"log/slog"

	"github.com/GHFluding/ShiftManager/SM/internal/database/postgres"
)

type Services struct {
	db  *postgres.Queries
	log *slog.Logger
}

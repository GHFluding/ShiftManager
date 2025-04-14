package main

import (
	"log/slog"
	"smgrpc/internal/config"
	sl "smgrpc/internal/utils/logger"
)

func main() {
	cfg := config.MustLoad()
	log := sl.SetupLogger(cfg.Env)
	log.Debug("starting application",
		slog.String("Env", cfg.Env),
	)
}

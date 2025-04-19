package main

import (
	"log/slog"
	"smgrpc/internal/app"
	grpcapp "smgrpc/internal/app/grpc"
	"smgrpc/internal/config"
	sl "smgrpc/internal/utils/logger"
)

func main() {
	cfg := config.MustLoad()

	log := sl.SetupLogger(cfg.Env)

	log.Debug("starting application",
		slog.String("Env", cfg.Env),
	)

	machineApplication := app.New(log, cfg.GRPC.Port, grpcapp.MachineServer)
	machineApplication.GRPCServer.MustRun()
}

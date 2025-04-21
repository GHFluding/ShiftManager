package main

import (
	"log/slog"
	"os"
	"os/signal"
	"smgrpc/internal/app"
	grpcapp "smgrpc/internal/app/grpc"
	"smgrpc/internal/config"
	sl "smgrpc/internal/utils/logger"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	log := sl.SetupLogger(cfg.Env)

	log.Debug("starting application",
		slog.String("Env", cfg.Env),
	)

	machineApplication := app.New(log, cfg.GRPC.Port, grpcapp.MachineServer)
	taskApplication := app.New(log, cfg.GRPC.Port+1, grpcapp.TaskServer)
	userApplication := app.New(log, cfg.GRPC.Port+2, grpcapp.UserServer)
	shiftApplication := app.New(log, cfg.GRPC.Port+3, grpcapp.ShiftServer)

	go machineApplication.GRPCServer.MustRun()
	go shiftApplication.GRPCServer.MustRun()
	go userApplication.GRPCServer.MustRun()
	go taskApplication.GRPCServer.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	machineApplication.GRPCServer.Stop()
	taskApplication.GRPCServer.Stop()
	userApplication.GRPCServer.Stop()
	shiftApplication.GRPCServer.Stop()

	log.Info("Gracefully stopped")
}

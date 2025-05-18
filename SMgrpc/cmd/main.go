package main

import (
	"log/slog"
	"os"
	"os/signal"
	"smgrpc/internal/app"
	grpcapp "smgrpc/internal/app/grpc"
	"smgrpc/internal/config"
	sl "smgrpc/internal/utils/logger"
	"strconv"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	log := sl.SetupLogger(cfg.Env)

	log.Debug("starting application",
		slog.String("Env", cfg.Env),
	)
	modeInt, err := strconv.Atoi(cfg.EnvMode)
	if err != nil {
		panic(err)
	}
	mode := grpcapp.CommandCode(modeInt)

	//TODO: Run without DB
	application := app.New(log, cfg.GRPC.Port, mode)

	go application.GRPCServer.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCServer.Stop()

	log.Info("Gracefully stopped")
}

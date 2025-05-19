package main

import (
	"log/slog"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/GHFluding/ShiftManager/SMgrpc/internal/app"
	grpcapp "github.com/GHFluding/ShiftManager/SMgrpc/internal/app/grpc"
	"github.com/GHFluding/ShiftManager/SMgrpc/internal/config"
	"github.com/GHFluding/ShiftManager/SMgrpc/internal/domain/models"
	sl "github.com/GHFluding/ShiftManager/SMgrpc/internal/utils/logger"
)

// main is an unused function in this module
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

	DB_Mock := models.DBFunction{}
	application := app.New(log, cfg.GRPC.Port, DB_Mock, mode)

	go application.GRPCServer.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCServer.Stop()

	log.Info("Gracefully stopped")
}

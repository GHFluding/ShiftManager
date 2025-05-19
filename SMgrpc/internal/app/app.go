package app

import (
	"log/slog"
	"os"
	"os/signal"
	grpcapp "smgrpc/internal/app/grpc"
	"smgrpc/internal/config"
	"smgrpc/internal/domain/models"
	sl "smgrpc/internal/utils/logger"
	"strconv"
	"syscall"
)

type App struct {
	GRPCServer *grpcapp.App
}

// New creates a new variable of type *App
func New(log *slog.Logger, grpcPort int, db_app models.DBFunction, command grpcapp.CommandCode) *App {
	grpcApp := grpcapp.New(command, log, db_app, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}

func (a *App) RunExternal(db_app models.DBFunction) {
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
	application := New(log, cfg.GRPC.Port, db_app, mode)

	//maybe use run
	go application.GRPCServer.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCServer.Stop()

	log.Info("Gracefully stopped")
}

package app

import (
	"log/slog"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	grpcapp "github.com/GHFluding/ShiftManager/SMgrpc/internal/app/grpc"
	"github.com/GHFluding/ShiftManager/SMgrpc/internal/config"
	sl "github.com/GHFluding/ShiftManager/SMgrpc/internal/utils/logger"
	"github.com/GHFluding/ShiftManager/SMgrpc/pkg/domain/models"
)

type App struct {
	GRPCServer *grpcapp.App
}

// New creates a new variable of type *App
func New(log *slog.Logger, grpcPort int, db_app models.DBFunction, command models.CommandCode) *App {
	grpcApp := grpcapp.New(command, log, db_app, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}

// Run method of *App type variable, which starts grpc server,
// takes as argument a set of functions implementing a number of interfaces for working with data
func (a *App) Run(db_app models.DBFunction) {
	cfg := config.MustLoad()

	log := sl.SetupLogger(cfg.Env)

	log.Debug("starting application",
		slog.String("Env", cfg.Env),
	)
	modeInt, err := strconv.Atoi(cfg.EnvMode)
	if err != nil {
		panic(err)
	}
	mode := models.CommandCode(modeInt)
	application := New(log, cfg.GRPC.Port, db_app, mode)

	go application.GRPCServer.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCServer.Stop()

	log.Info("Gracefully stopped")
}

package shiftManager

import (
	"fmt"

	"github.com/GHFluding/ShiftManager/SM/internal/config"
	"github.com/GHFluding/ShiftManager/SM/internal/database"
	"github.com/GHFluding/ShiftManager/SM/internal/database/postgres"
	grpc_storage "github.com/GHFluding/ShiftManager/SM/internal/grpcStorage"
	"github.com/GHFluding/ShiftManager/SM/internal/services"
	grpchandler "github.com/GHFluding/ShiftManager/SM/internal/transport/gRPCHandler"
	"github.com/GHFluding/ShiftManager/SM/internal/utils/logger"
	"github.com/golang-migrate/migrate/v4"
)

func RunGRPC() {
	cfg := config.MustLoad()
	log := logger.Setup(cfg.Env)
	log.Debug("Successful load config")
	log.Debug("Config env: " + cfg.Env)
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Storage.User, cfg.Storage.Password, cfg.Storage.Host, cfg.Storage.Port, cfg.Storage.DBName)

	dbpool, err := database.Connect(*cfg, log)
	if err != nil {
		log.Info("No connection with database", logger.ErrToAttr(err))
	}
	defer dbpool.Close()

	//TODO: init migrations
	m, err := migrate.New(
		"file://migrations",
		dbURL,
	)
	if err != nil {
		log.Error("failed to initialize migrations", logger.ErrToAttr(err))
		return
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Error("failed to apply migrations", logger.ErrToAttr(err))
		return
	} else {
		log.Info("migrations applied successfully")
	}

	db := postgres.New(dbpool)
	services.NewServicesParams(db, log)

	handler := grpchandler.New(log, cfg.Storage.Port, grpc_storage.New(db, log))
	handler.RunMachine()
	handler.RunShift()
	handler.RunTask()
	handler.RunUser()
}

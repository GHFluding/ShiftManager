package shiftManager

import (
	"fmt"
	"sm/internal/config"
	"sm/internal/database"
	"sm/internal/database/postgres"
	"sm/internal/transport/handler"
	"sm/internal/transport/middleware"
	"sm/internal/utils/handler_utils"
	"sm/internal/utils/logger"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
)

func Run() {
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

	r := gin.Default()
	r.Use(middleware.RequestId())

	handlerParams := handler_utils.CreateParams(db, log)

	usersGroup := r.Group("/api/users")
	{
		usersGroup.GET("/", handler.GetUserList(handlerParams))
		usersGroup.GET("/:role", handler.GetUserListByRole(handlerParams))
		usersGroup.DELETE("/:id", handler.DeleteUser(handlerParams))
	}
	machineGroup := r.Group("/api/machine")
	{
		machineGroup.PUT("/:id", handler.ChangeMachineToRepair(handlerParams))
	}
	shiftGroup := r.Group("/api/shifts")
	{
		shiftGroup.GET("/", handler.GetShiftList(handlerParams))
		shiftGroup.GET("/active/", handler.GetActiveShiftList(handlerParams))
		shiftGroup.GET("/tasks/", handler.GetShiftTaskList(handlerParams))
		shiftGroup.GET("/workers/", handler.GetShiftWorkersList(handlerParams))
	}

}

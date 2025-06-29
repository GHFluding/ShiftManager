package shiftManager

import (
	"fmt"

	"github.com/GHFluding/ShiftManager/SM/internal/config"
	"github.com/GHFluding/ShiftManager/SM/internal/database"
	"github.com/GHFluding/ShiftManager/SM/internal/database/postgres"
	"github.com/GHFluding/ShiftManager/SM/internal/services"
	"github.com/GHFluding/ShiftManager/SM/internal/transport/handler"
	"github.com/GHFluding/ShiftManager/SM/internal/transport/middleware"
	"github.com/GHFluding/ShiftManager/SM/internal/utils/logger"

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
	services.NewServicesParams(db, log)

	r := gin.Default()
	r.Use(middleware.RequestId())

	handlerParams := services.NewServicesParams(db, log)

	usersGroup := r.Group("/api/users")
	{
		usersGroup.GET("/", handler.GetUserList(log, handlerParams))
		usersGroup.POST("/", handler.CreateUser(log, handlerParams))
		usersGroup.GET("/:role", handler.GetUserListByRole(log, handlerParams))
		usersGroup.DELETE("/:id", handler.DeleteUser(log, handlerParams))
		usersGroup.GET("/:id", handler.CheckUserRoleHandler(log, handlerParams))
	}
	machineGroup := r.Group("/api/machine")
	{
		machineGroup.PUT("/:id", handler.ChangeMachineToRepair(log, handlerParams))
		machineGroup.POST("/", handler.CreateMachine(log, handlerParams))
	}
	shiftGroup := r.Group("/api/shifts")
	{
		shiftGroup.GET("/", handler.GetShiftList(log, handlerParams))
		shiftGroup.GET("/active/", handler.GetActiveShiftList(log, handlerParams))
		shiftGroup.GET("/tasks/", handler.GetShiftTaskList(log, handlerParams))
		shiftGroup.DELETE("/:shiftid/taskid/:taskid", handler.DeleteShiftTask(log, handlerParams))
		shiftGroup.DELETE("/:shiftid/workerid/:userid", handler.DeleteShiftWorker(log, handlerParams))
		shiftGroup.GET("/workers/", handler.GetShiftWorkersList(log, handlerParams))
		shiftGroup.POST("/", handler.CreateShift(log, handlerParams))
	}
	taskGroup := r.Group("/api/task")
	{
		taskGroup.POST("/", handler.CreateShift(log, handlerParams))
		taskGroup.PATCH("/{id}", handler.UpdateTask(log, handlerParams))
	}
}

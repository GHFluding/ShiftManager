package shiftManager

import (
	"sm/internal/config"
	"sm/internal/database"
	"sm/internal/database/postgres"
	"sm/internal/transport/handler"
	"sm/internal/transport/middleware"
	"sm/internal/utils/handler_utils"
	"sm/internal/utils/logger"

	"github.com/gin-gonic/gin"
)

func Run() {
	cfg := config.MustLoad()
	log := logger.Setup(cfg.Env)
	log.Debug("Successful load config")
	log.Debug("Config env: " + cfg.Env)
	//TODO: init database
	dbpool, err := database.Connect(*cfg, log)
	if err != nil {
		log.Info("No connection with database", logger.ErrToAttr(err))
	}
	defer dbpool.Close()

	//TODO: init migrations

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

}

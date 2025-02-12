package shiftManager

import (
	"sm/internal/config"
	"sm/internal/transport/middleware"
	"sm/internal/utils/logger"

	"github.com/gin-gonic/gin"
)

func Run() {
	cfg := config.MustLoad()
	log := logger.Setup(cfg.Env)
	log.Debug("Successful load config")
	log.Debug("Config env: " + cfg.Env)
	//TODO: init database

	//TODO: init migrations
	
	r := gin.Default()
	r.Use(middleware.RequestId())
	
	//TODO: add handlers

}

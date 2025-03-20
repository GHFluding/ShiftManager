package main

import (
	"linkSM/internal/config"
	logger "linkSM/internal/utils"
)

func main() {
	cfg := config.MustLoad()
	log := logger.Setup(cfg.Env)
	log.Info("Logger is up")
	//TODO: init webhooks
	//webhooks to accept request and return result

	//TODO: init routers
	//routers to send request to db

}

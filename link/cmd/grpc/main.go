package main

import (
	"github.com/GHFluding/ShiftManager/SMgrpc/pkg/client"
	"github.com/GHFluding/ShiftManager/link/internal/config"
	logger "github.com/GHFluding/ShiftManager/link/internal/utils"
)

func main() {
	cfg := config.MustLoad()
	log := logger.Setup(cfg.Env)
	log.Info("Logger is up")

	c, err := client.New(cfg.Routing.GetBaseURL())
	if err != nil {
		panic("Server started with error: " + err.Error())
	}
	_ = c
	//TODO: init routers
	//routers to send request to db

}

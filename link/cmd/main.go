package main

import (
	"github.com/GHFluding/ShiftManager/link/internal/config"
	"github.com/GHFluding/ShiftManager/link/internal/transport/webhooks"
	logger "github.com/GHFluding/ShiftManager/link/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.MustLoad()
	log := logger.Setup(cfg.Env)
	log.Info("Logger is up")

	r := gin.Default()
	webhookTaskGroup := r.Group(cfg.Webhooks.Task)
	{
		webhookTaskGroup.POST("/create", webhooks.ProcessWebhookGRPC(log, cfg.Routing.GetTaskBaseURL()))
	}
	webhookUserGroup := r.Group(cfg.Webhooks.Users)
	{
		webhookUserGroup.POST("/create", webhooks.ProcessWebhookGRPC(log, cfg.Routing.GetUserBaseURL()))
	}
	webhookMachineGroup := r.Group(cfg.Webhooks.Machine)
	{
		webhookMachineGroup.POST("/create", webhooks.ProcessWebhookGRPC(log, cfg.Routing.GetMachineBaseURL()))
	}
	webhookShiftGroup := r.Group(cfg.Webhooks.Shift)
	{
		webhookShiftGroup.POST("/create", webhooks.ProcessWebhookGRPC(log, cfg.Routing.GetShiftBaseURL()))
	}
	//TODO: init routers
	//routers to send request to db

}

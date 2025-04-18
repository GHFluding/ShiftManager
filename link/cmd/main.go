package main

import (
	"linkSM/internal/config"
	"linkSM/internal/services"
	"linkSM/internal/transport/webhooks"
	logger "linkSM/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.MustLoad()
	log := logger.Setup(cfg.Env)
	log.Info("Logger is up")

	r := gin.Default()
	webhookTaskGroup := r.Group(cfg.Webhooks.Task)
	{
		webhookTaskGroup.POST("/", webhooks.WebhookHandler(log, services.CreateTask, cfg.Routing.GetTaskBaseURL()))
	}
	webhookUserGroup := r.Group(cfg.Webhooks.Users)
	{
		webhookUserGroup.POST("/", webhooks.WebhookHandler(log, services.CreateUser, cfg.Routing.GetUserBaseURL()))
	}
	webhookMachineGroup := r.Group(cfg.Webhooks.Machine)
	{
		webhookMachineGroup.POST("/", webhooks.WebhookHandler(log, services.CreateMachine, cfg.Routing.GetMachineBaseURL()))
	}
	webhookShiftGroup := r.Group(cfg.Webhooks.Shift)
	{
		webhookShiftGroup.POST("/", webhooks.WebhookHandler(log, services.CreateShift, cfg.Routing.GetShiftBaseURL()))
	}
	//TODO: init routers
	//routers to send request to db

}

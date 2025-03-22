package main

import (
	config "bsm/internal/config/loadconfig"

	apilogic "bsm/internal/services/apiLogic"
	"bsm/internal/utils/logger"
	"log"

	"github.com/gin-gonic/gin"
	goBX24 "github.com/whatcrm/go-bitrix24"
)

func main() {
	// TODO: init config loading and making webhook struct with data of .env file
	cfg := config.MustLoad()
	webhook := config.WebhookB24Init(
		"example-client-id",
		"example-client-secret",
		"example-bitrix24-domain",
		"example-auth-token",
		"example-url",
	)
	b24 := goBX24.NewAPI(webhook.GetID(), webhook.GetSecret())
	if err := b24.SetOptions(webhook.GetDomain(), webhook.GetAuthToken(), true); err != nil {
		log.Fatalf("Setting API error: %v", err)
	}
	log := logger.Setup(cfg.Env)
	r := gin.Default()
	webhookURL := webhook.GetURL()
	// Handling incoming message
	r.POST(webhookURL, apilogic.HandleMessage(cfg, log))

	if err := r.Run(":8080"); err != nil {
		log.Info("Run server error: ", logger.ErrToAttr(err))
	}
}

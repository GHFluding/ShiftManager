package main

import (
	"bsm/internal/config/structures/variables"
	apilogic "bsm/internal/services/apiLogic"
	"bsm/internal/services/logger"
	"log"

	"github.com/gin-gonic/gin"
	goBX24 "github.com/whatcrm/go-bitrix24"
)

func main() {
	// TODO: init config loading and making webhook struct with data of .env file
	webhook := variables.WebhookInit(
		"example-client-id",
		"example-client-secret",
		"example-bitrix24-domain",
		"example-auth-token",
	)
	b24 := goBX24.NewAPI(webhook.GetID(), webhook.GetSecret())
	if err := b24.SetOptions(webhook.GetDomain(), webhook.GetAuthToken(), true); err != nil {
		log.Fatalf("Setting API error: %v", err)
	}
	log := logger.Setup("local")
	r := gin.Default()
	// Handling incoming message
	r.POST("/webhook", apilogic.HandleMessage(log))

	if err := r.Run(":8080"); err != nil {
		log.Info("Run server error: ", logger.ErrToAttr(err))
	}
}

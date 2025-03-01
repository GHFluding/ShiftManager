package bsm

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	//TODO: init logger
	webhookURL := "https://your-bitrix-portal.bitrix24.ru/rest/1/WEBHOOK_TOKEN"
	secret := "your_webhook_secret"

	bitrixClient := bitrix.NewClient(webhookURL)
	botHandler := bot.NewBotHandler(bitrixClient, secret)

	// TODO: inti kafka

	//TODO: rework
	r := gin.Default()
	r.POST("/bitrix/webhook", botHandler.WebhookHandler)

	if err := r.Run(":8080"); err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}

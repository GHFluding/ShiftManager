package bsm

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// TODO: inti kafka
func main() {
	// Конфигурация
	webhookURL := "https://your-bitrix-portal.bitrix24.ru/rest/1/WEBHOOK_TOKEN"
	secret := "your_webhook_secret"

	// Инициализация клиента
	bitrixClient := bitrix.NewClient(webhookURL)
	botHandler := bot.NewBotHandler(bitrixClient, secret)

	// Настройка роутера
	r := gin.Default()
	r.POST("/bitrix/webhook", botHandler.WebhookHandler)

	// Запуск сервера
	if err := r.Run(":8080"); err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}

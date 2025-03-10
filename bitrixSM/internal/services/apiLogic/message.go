package apilogic

import (
	bot_command "bsm/internal/services/command"
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type IncomingMessage struct {
	UpdateID int `json:"update_id"`
	Message  struct {
		MessageID int    `json:"message_id"`
		Text      string `json:"text"`
		Chat      struct {
			ID int `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

func HandleMessage(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var msg IncomingMessage
		if err := c.ShouldBindJSON(&msg); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		spreadingMessage(msg, log)

		c.JSON(http.StatusOK, gin.H{"status": "received"})
	}
}

func spreadingMessage(msg IncomingMessage, log *slog.Logger) {
	parts := strings.Fields(msg.Message.Text)
	if len(parts) == 0 {
		return
	}

	command := parts[0]
	args := parts[1:]
	switch command {
	case "/help":
		resp := bot_command.Help(log)
		sendMessageToBitrix(int(msg.Message.Chat.ID), resp)
	case "/create-task":
		err := bot_command.CreateTask(args, log)
		if err != nil {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Ошибка выполнения команды по созданию задания")
		} else {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Задание успешно добавлено")
		}
	case "/shift-list":
		list, err := bot_command.ShiftList(args, log)
		if err != nil {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Ошибка выполения команды по получению списка смен")
		} else {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Список смен:\n"+list)
		}
	}

}

func sendMessageToBitrix(chatID int, message string) error {
	url := "https:/example-domain.bitrix24.ru/rest/1/example-webhook-token/im.message.add"
	data := map[string]interface{}{
		"DIALOG_ID": chatID,
		"MESSAGE":   message,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send message, status code: %d", resp.StatusCode)
	}

	return nil
}

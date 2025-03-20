package apilogic

import (
	config "bsm/internal/config/loadconfig"
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

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

func HandleMessage(cfg *config.Config, log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var msg IncomingMessage
		if err := c.ShouldBindJSON(&msg); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		spreadingMessage(cfg, msg, log)

		c.JSON(http.StatusOK, gin.H{"status": "received"})
	}
}

type sendingMessage struct {
	DialogId int
	Message  string
}

func sendMessageToBitrix(chatID int, message string, url string) error {

	data := sendingMessage{
		DialogId: chatID,
		Message:  message,
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

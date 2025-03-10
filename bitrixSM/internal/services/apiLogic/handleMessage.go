package apilogic

import (
	"bsm/internal/services/command"
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

var commands = map[string]func([]string) error{
	"/create-task": command.CreateTask,
	"/help":        command.Help,
}

func spreadingMessage(msg IncomingMessage, log *slog.Logger) {
	parts := strings.Fields(msg.Message.Text)
	if len(parts) == 0 {
		return
	}

	command := parts[0]
	args := parts[1:]

	if commandFunc, exists := commands[command]; exists {
		commandFunc(args)
	} else {
		log.Info("Unknown command: ", "\n", command)
	}

}

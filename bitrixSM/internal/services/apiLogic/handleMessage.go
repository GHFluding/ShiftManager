package apilogic

import (
	"fmt"
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

func HandleMessage(c *gin.Context) {
	var msg IncomingMessage
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	spreadingMessage(msg)

	c.JSON(http.StatusOK, gin.H{"status": "received"})
}

func spreadingMessage(msg IncomingMessage) {
	fmt.Printf("Incoming message: %s\n", msg.Message.Text)
	// TODO: init spreading
}

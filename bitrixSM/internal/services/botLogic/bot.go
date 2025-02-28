package bot

import (
	b24models "bsm/internal/models/b24"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BotHandler struct {
	BitrixClient *b24models.Client
	Secret       []byte
}

func NewBotHandler(client *b24models.Client, secret string) *BotHandler {
	return &BotHandler{
		BitrixClient: client,
		Secret:       []byte(secret),
	}
}

func (h *BotHandler) WebhookHandler(c *gin.Context) {

	if err := b24models.VerifyRequest(c.Request, h.Secret); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	event, err := b24models.ParseWebhook(c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	switch event.Event {
	case "ONIMBOTMESSAGEADD":
		var message b24models.ImMessage
		if err := json.Unmarshal(event.Data, &message); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid message data"})
			return
		}
		h.handleMessage(message)

	case "ONTASKADD":
		var task b24models.Task
		if err := json.Unmarshal(event.Data, &task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task data"})
			return
		}
		h.handleTask(task)

	default:
		c.JSON(http.StatusOK, gin.H{"status": "unhandled event type"})
	}

	c.JSON(http.StatusOK, gin.H{"status": "processed"})
}

func (h *BotHandler) handleMessage(msg b24models.ImMessage) {
	//TODO
}

func (h *BotHandler) handleTask(msg b24models.Task) {
	//TODO
}

func (h *BotHandler) createTaskFromMessage(msg b24models.ImMessage) {
	//TODO
}

func (h *BotHandler) sendMessage(dialogID, text string) error {
	params := map[string]interface{}{
		"DIALOG_ID": dialogID,
		"MESSAGE":   text,
	}

	var response interface{}
	return h.BitrixClient.CallMethod("im.message.add", params, &response)
}

package bot

import (
	b24models "bsm/internal/models/b24"
	handler "bsm/internal/transport/handlers"
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
		h.handleMessage(c)

	case "ONTASKADD":
		c.JSON(http.StatusBadRequest, gin.H{"error": "this bot can't make task on b24"})

	default:
		c.JSON(http.StatusOK, gin.H{"status": "unhandled event type"})
	}

	c.JSON(http.StatusOK, gin.H{"status": "processed"})
}

func (h *BotHandler) handleMessage(c *gin.Context) {
	var event struct {
		Data struct {
			DialogID string `json:"DIALOG_ID"`
			Message  string `json:"MESSAGE"`
			UserID   int    `json:"USER_ID"`
		}
	}

	if err := c.BindJSON(&event); err != nil {
		c.JSON(400, gin.H{"error": "invalid payload"})
		return
	}

	// maybe refactor
	switch event.Data.Message {
	case "/start":
		const helloString = "Напишите /help для получения списка команда"
		h.sendMessage(event.Data.DialogID, helloString)
	case "/help":
		const commandList = `/create_task - создает задание
		`
		h.sendMessage(event.Data.DialogID, commandList)
	case "/create_task":
		err := handler.CreateTask(c)
		if err != nil {
			h.sendMessage(event.Data.DialogID, err.Error())
		} else {
			//task created successfully
			h.sendMessage(event.Data.DialogID, "задание успешно созданно ")
		}
	default:
		//error
	}
}

func (h *BotHandler) sendMessage(dialogID, text string) error {
	params := map[string]interface{}{
		"DIALOG_ID": dialogID,
		"MESSAGE":   text,
	}

	var response interface{}
	return h.BitrixClient.CallMethod("im.message.add", params, &response)
}

package apilogic

import (
	config "bsm/internal/config/loadconfig"
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

func spreadingMessage(cfg *config.Config, msg IncomingMessage, log *slog.Logger) {
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
		err := bot_command.CreateTask(cfg.Webhooks.GetCreateTaskUrl(), args, log)
		if err != nil {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Ошибка выполнения команды по созданию задания")
		} else {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Задание успешно добавлено")
		}
	case "/shift-list":
		list, err := bot_command.ShiftList(cfg.Webhooks.GetShiftListUrl(), log)
		if err != nil {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Ошибка выполения команды для получения списка смен")
		} else {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Список смен:\n"+list)
		}
	case "/create-shift":
		resp, err := bot_command.CreateShift(cfg.Webhooks.GetCreateShiftUrl(), args, log)
		if err != nil {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Ошибка выполения команды для создания смены смен")
		} else {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Смена упешно создана:\n"+resp)
		}
	case "/add-worker":
		resp, err := bot_command.AddShiftWorker(cfg.Webhooks.GETAddShiftWorkerURL(), args, log)
		if err != nil {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Ошибка выполения команды добавления работника на смену")
		} else {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Работник успешно добавлен:\n"+resp)
		}
	case "/task-list":
		resp, err := bot_command.GetTaskList(cfg.Webhooks.GETTaskListURL(), log)
		if err != nil {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Ошибка выполения команды получения списка заданий")
		} else {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Список заданий:\n"+resp)
		}
	case "/add-machine":
		resp, err := bot_command.AddMachine(cfg.Webhooks.GETAddMachineURL(), args, log)
		if err != nil {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Ошибка выполения команды добавления работника на смену")
		} else {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Работник успешно добавлен:\n"+resp)
		}
	default:
		sendMessageToBitrix(int(msg.Message.Chat.ID), "Неверная команда, отправте /help для получения списка команд")
	}

}

type sendingMessage struct {
	dialogId int
	message  string
}

func sendMessageToBitrix(chatID int, message string) error {
	url := "https:/example-domain.bitrix24.ru/rest/1/example-webhook-token/im.message.add"
	data := sendingMessage{
		dialogId: chatID,
		message:  message,
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

package apilogic

import (
	"log/slog"
	"strings"

	config "github.com/GHFluding/ShiftManager/bitrixSM/internal/config/loadconfig"
	bot_command "github.com/GHFluding/ShiftManager/bitrixSM/internal/services/command"
)

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
		sendMessageToBitrix(int(msg.Message.Chat.ID), resp, cfg.WebhookB24.GetURL())
	case "/create-task":
		err := bot_command.CreateTask(cfg.Webhooks.GETBaseURL(), args, log)
		if err != nil {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Ошибка выполнения команды по созданию задания", cfg.WebhookB24.GetURL())
		} else {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Задание успешно добавлено", cfg.WebhookB24.GetURL())
		}
	case "/shift-list":
		list, err := bot_command.ShiftList(cfg.Webhooks.GETBaseURL(), log)
		if err != nil {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Ошибка выполения команды для получения списка смен", cfg.WebhookB24.GetURL())
		} else {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Список смен:\n"+list, cfg.WebhookB24.GetURL())
		}
	case "/create-shift":
		resp, err := bot_command.CreateShift(cfg.Webhooks.GETBaseURL(), args, log)
		if err != nil {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Ошибка выполения команды для создания смены смен", cfg.WebhookB24.GetURL())
		} else {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Смена упешно создана:\n"+resp, cfg.WebhookB24.GetURL())
		}
	case "/add-worker":
		resp, err := bot_command.AddShiftWorker(cfg.Webhooks.GETBaseURL(), args, log)
		if err != nil {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Ошибка выполения команды добавления работника на смену", cfg.WebhookB24.GetURL())
		} else {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Работник успешно добавлен:\n"+resp, cfg.WebhookB24.GetURL())
		}
	case "/task-list":
		resp, err := bot_command.GetTaskList(cfg.Webhooks.GETBaseURL(), log)
		if err != nil {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Ошибка выполения команды получения списка заданий", cfg.WebhookB24.GetURL())
		} else {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Список заданий:\n"+resp, cfg.WebhookB24.GetURL())
		}
	case "/add-machine":
		resp, err := bot_command.AddMachine(cfg.Webhooks.GETBaseURL(), args, log)
		if err != nil {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Ошибка выполения команды добавления работника на смену", cfg.WebhookB24.GetURL())
		} else {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Работник успешно добавлен:\n"+resp, cfg.WebhookB24.GetURL())
		}
	default:
		sendMessageToBitrix(int(msg.Message.Chat.ID), "Неверная команда, отправте /help для получения списка команд", cfg.WebhookB24.GetURL())
	}

}

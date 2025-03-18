package apilogic

import (
	config "bsm/internal/config/loadconfig"
	bot_command "bsm/internal/services/command"
	"log/slog"
	"strings"
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
		sendMessageToBitrix(int(msg.Message.Chat.ID), resp)
	case "/create-task":
		err := bot_command.CreateTask(cfg.Webhooks.GETBaseURL(), args, log)
		if err != nil {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Ошибка выполнения команды по созданию задания")
		} else {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Задание успешно добавлено")
		}
	case "/shift-list":
		list, err := bot_command.ShiftList(cfg.Webhooks.GETBaseURL(), log)
		if err != nil {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Ошибка выполения команды для получения списка смен")
		} else {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Список смен:\n"+list)
		}
	case "/create-shift":
		resp, err := bot_command.CreateShift(cfg.Webhooks.GETBaseURL(), args, log)
		if err != nil {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Ошибка выполения команды для создания смены смен")
		} else {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Смена упешно создана:\n"+resp)
		}
	case "/add-worker":
		resp, err := bot_command.AddShiftWorker(cfg.Webhooks.GETBaseURL(), args, log)
		if err != nil {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Ошибка выполения команды добавления работника на смену")
		} else {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Работник успешно добавлен:\n"+resp)
		}
	case "/task-list":
		resp, err := bot_command.GetTaskList(cfg.Webhooks.GETBaseURL(), log)
		if err != nil {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Ошибка выполения команды получения списка заданий")
		} else {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Список заданий:\n"+resp)
		}
	case "/add-machine":
		resp, err := bot_command.AddMachine(cfg.Webhooks.GETBaseURL(), args, log)
		if err != nil {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Ошибка выполения команды добавления работника на смену")
		} else {
			sendMessageToBitrix(int(msg.Message.Chat.ID), "Работник успешно добавлен:\n"+resp)
		}
	default:
		sendMessageToBitrix(int(msg.Message.Chat.ID), "Неверная команда, отправте /help для получения списка команд")
	}

}

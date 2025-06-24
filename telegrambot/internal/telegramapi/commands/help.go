package commands

import (
	"context"
	"fmt"
	"telegramSM/internal/telegramapi/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HelpHandler(userService UserService) model.ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		chatID := update.Message.Chat.ID
		userID := update.Message.From.ID

		user, err := userService.GetUser(ctx, userID)
		if err != nil {
			msg := tgbotapi.NewMessage(chatID, "Ошибка получения данных пользователя")
			_, _ = bot.Send(msg)
			return err
		}

		var commands []model.CommandMeta
		switch user.Role {
		case model.RoleAdmin:
			commands = model.AdminCommands
		case model.RoleManager:
			commands = model.ManagerCommands
		case model.RoleMaster:
			commands = model.MasterCommands
		case model.RoleWorker:
			commands = model.WorkerCommands
		default:
			msg := tgbotapi.NewMessage(chatID, "Вначале пройдите регистрацию: /start")
			_, err = bot.Send(msg)
		}
		messageText := "Доступные команды:\n\n"
		for _, cmd := range commands {
			messageText += fmt.Sprintf("/%s - %s\n", cmd.Type, cmd.Description)
		}

		msg := tgbotapi.NewMessage(chatID, messageText)
		_, err = bot.Send(msg)
		return err
	}
}

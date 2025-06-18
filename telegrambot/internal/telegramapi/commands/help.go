package commands

import (
	"context"
	"fmt"
	"telegramSM/internal/telegramapi/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HelpHandler(userService UserService) model.ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		userID := update.Message.From.ID
		chatID := update.Message.Chat.ID

		user, err := userService.GetUser(ctx, userID)
		if err != nil {
			msg := tgbotapi.NewMessage(chatID, fmt.Sprintf(
				"введите "+string(model.CmdStart)+" для регистрации.",
			))
			_, err = bot.Send(msg)
			return err
		}
		switch user.Role {
		case model.RoleAdmin:

		}

		return nil
	}
}

func adminHelp(userService UserService) model.ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		var messageText string
		for _, adminCmd := range model.AdminCommands {
			messageText += string(adminCmd) + " - " + model.GetDescription(adminCmd) + "\n"
		}
		chatID := update.Message.Chat.ID
		msg := tgbotapi.NewMessage(chatID, messageText)
		_, err := bot.Send(msg)
		return err
	}
}

package commands

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"telegramSM/internal/telegramapi/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// User, UserData и Role
type User struct {
	TelegramID int      `json:"telegram_id"`
	Data       UserData `json:"data"`
	Role       string   `json:"role"`
}

const emptyString = ""

type UserData struct {
	Name   string `json:"name"`
	BtrxID string `json:"btrx_id"`
}

// UserService is interface for handler function
type UserService interface {
	GetUser(ctx context.Context, telegramID int) (*User, error)
	SaveUser(ctx context.Context, user *User) error
}

type RestUserService struct {
	baseURL string
	client  *http.Client
}

// StartHandler
func StartHandler(userService UserService) model.ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		userID := update.Message.From.ID
		chatID := update.Message.Chat.ID

		user, err := userService.GetUser(ctx, userID)
		if err != nil {
			user = &User{TelegramID: userID}
		}
		if user.Data.Name == emptyString {
			handler := NameHandler(userService)
			if err := handler(ctx, bot, update); err != nil {
				return err
			}
		}
		if user.Data.BtrxID == emptyString {
			handler := BitrixHandler(userService)
			if err := handler(ctx, bot, update); err != nil {
				return err
			}
		}
		msg := tgbotapi.NewMessage(chatID, fmt.Sprintf(
			"👋 Добро пожаловать, %s!\n\n "+string(model.CmdHelp)+" для списка команд.",
			user.Data.Name,
		))
		_, err = bot.Send(msg)
		return err
	}
}

func NameHandler(userService UserService) model.ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		userID := update.Message.From.ID
		chatID := update.Message.Chat.ID
		name := strings.TrimSpace(update.Message.Text)

		user, err := userService.GetUser(ctx, userID)
		if err != nil {
			return err
		}

		if len(name) < 5 || !strings.Contains(name, " ") {
			msg := tgbotapi.NewMessage(chatID, "❌ Неверный формат ФИО. Введите полное имя:")
			_, _ = bot.Send(msg)
			return nil
		}

		user.Data.Name = name
		_ = userService.SaveUser(ctx, user)

		msg := tgbotapi.NewMessage(chatID, "✅ ФИО сохранено! Теперь введите ваш Bitrix24 ID (или 'нет'):")
		_, err = bot.Send(msg)
		return err
	}
}

func BitrixHandler(userService UserService) model.ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		userID := update.Message.From.ID
		chatID := update.Message.Chat.ID
		input := strings.TrimSpace(update.Message.Text)

		user, err := userService.GetUser(ctx, userID)
		if err != nil {
			return err
		}

		if strings.ToLower(input) != "нет" {
			user.Data.BtrxID = input
		}

		_ = userService.SaveUser(ctx, user)

		confirmation := fmt.Sprintf("✅ Регистрация завершена!\n\n👤 ФИО: %s", user.Data.Name)
		if user.Data.BtrxID != "" {
			confirmation += fmt.Sprintf("\n Bitrix24 ID: %s", user.Data.BtrxID)
		}
		confirmation += "\n\nТеперь вам доступны все функции бота."
		msg := tgbotapi.NewMessage(chatID, confirmation)
		_, err = bot.Send(msg)
		return err
	}
}

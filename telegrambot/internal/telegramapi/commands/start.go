package commands

import (
	"context"
	"fmt"
	"strings"
	"telegramSM/internal/telegramapi/model"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api"
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

func StartHandler(userService UserService) model.ViewFunc {
	return func(ctx context.Context, bot *tgBotAPI.BotAPI, update tgBotAPI.Update) error {
		userID := update.Message.From.ID
		chatID := update.Message.Chat.ID

		user, err := userService.GetUser(ctx, userID)
		if err != nil {
			user = &User{
				TelegramID: userID,
				Data: UserData{
					Name:   "",
					BtrxID: "",
				},
			}
			if err := userService.SaveUser(ctx, user); err != nil {
				return err
			}
		}

		if user.Data.Name == emptyString {
			msg := tgBotAPI.NewMessage(chatID, "👤 Введите ваше ФИО (полное имя с пробелом):")
			_, err = bot.Send(msg)
			return err
		}

		if user.Data.BtrxID == emptyString {
			return askBitrixID(ctx, bot, chatID)
		}

		msg := tgBotAPI.NewMessage(chatID, fmt.Sprintf(
			"👋 Добро пожаловать, %s!\n\nВведите /help для списка команд.",
			user.Data.Name,
		))
		_, err = bot.Send(msg)
		return err
	}
}

func NameHandler(userService UserService) model.ViewFunc {
	return func(ctx context.Context, bot *tgBotAPI.BotAPI, update tgBotAPI.Update) error {
		userID := update.Message.From.ID
		chatID := update.Message.Chat.ID
		name := strings.TrimSpace(update.Message.Text)

		user, err := userService.GetUser(ctx, userID)
		if err != nil {
			user = &User{
				TelegramID: userID,
				Data:       UserData{},
				Role:       model.RoleWorker,
			}
		}

		if len(name) < 5 || !strings.Contains(name, " ") {
			msg := tgBotAPI.NewMessage(chatID, "❌ Неверный формат ФИО. Введите полное имя:")
			_, _ = bot.Send(msg)
			return nil
		}

		user.Data.Name = name
		user.Data.BtrxID = emptyString
		if err := userService.SaveUser(ctx, user); err != nil {
			return err
		}

		return askBitrixID(ctx, bot, chatID)
	}
}

func askBitrixID(
	ctx context.Context,
	bot *tgBotAPI.BotAPI,
	chatID int64,
) error {
	msg := tgBotAPI.NewMessage(chatID, "Введите ваш Bitrix24 ID или нажмите 'Пропустить':")
	keyboard := createInlineKeyboard(
		[]string{"Пропустить"},
		func(item string) string { return item },
		func(item string) string { return "skip_bitrix" },
	)

	msg.ReplyMarkup = keyboard
	_, err := bot.Send(msg)
	return err
}

func BitrixHandler(userService UserService) model.ViewFunc {
	return func(ctx context.Context, bot *tgBotAPI.BotAPI, update tgBotAPI.Update) error {
		bitrixID := strings.TrimSpace(update.Message.Text)
		if bitrixID == "" {
			return nil
		}

		userID := update.Message.From.ID
		chatID := update.Message.Chat.ID

		user, err := userService.GetUser(ctx, userID)
		if err != nil {
			return err
		}
		if user.Data.Name != emptyString && user.Data.Name != bitrixID {
			user.Data.BtrxID = bitrixID
			if err := userService.SaveUser(ctx, user); err != nil {
				return err
			}
			userData, err := userService.GetUser(ctx, userID)
			if err != nil {
				return err
			}
			confirmation := fmt.Sprintf(
				"✅ Регистрация завершена!\n👤 ФИО: %s\n🔗 Bitrix24 ID: %s\n\nТеперь вам доступны все функции бота.",
				userData.Data.Name,
				userData.Data.BtrxID,
			)
			msg := tgBotAPI.NewMessage(chatID, confirmation)
			_, err = bot.Send(msg)
		}
		return err
	}
}

func SkipBitrixHandler(userService UserService) model.ViewFunc {
	return func(ctx context.Context, bot *tgBotAPI.BotAPI, update tgBotAPI.Update) error {
		callback := update.CallbackQuery
		if callback == nil || callback.Data != "skip_bitrix" {
			return nil
		}

		userID := callback.From.ID
		chatID := callback.Message.Chat.ID

		user, err := userService.GetUser(ctx, userID)
		if err != nil {
			return err
		}

		user.Data.BtrxID = ""
		if err := userService.SaveUser(ctx, user); err != nil {
			return err
		}

		edit := tgBotAPI.NewEditMessageReplyMarkup(chatID, callback.Message.MessageID, tgBotAPI.InlineKeyboardMarkup{})
		bot.Send(edit)

		editText := tgBotAPI.NewEditMessageText(chatID, callback.Message.MessageID, "✅ Ввод Bitrix24 ID пропущен")
		bot.Send(editText)

		confirmation := fmt.Sprintf(
			"✅ Регистрация завершена!\n👤 ФИО: %s\n🔗 Bitrix24 ID пропущен.\n\nТеперь вам доступны все функции бота.",
			user.Data.Name,
		)

		msg := tgBotAPI.NewMessage(chatID, confirmation)
		_, err = bot.Send(msg)
		return err
	}
}

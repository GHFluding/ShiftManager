package commands

import (
	"context"
	"fmt"
	"strings"
	"telegramSM/internal/telegramapi/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type state string

const (
	StateInitial    = state("initial")
	StateAwaitName  = state("await_name")
	StateAwaitBtrx  = state("await_btrx")
	StateRegistered = state("registered")
)

type User struct {
	TelegramID int
	Data       UserData
	State      state
	Role       string
}

type UserData struct {
	Name   string
	BtrxID string
}

type UserStorage interface {
	GetUser(ctx context.Context, telegramID int) (*User, error)
	SaveUser(ctx context.Context, user User) error
	CheckUserRole(ctx context.Context, telegramID int) (string, error)
}

// StartHandler - handler command /start with user role cheking
func StartHandler(userStorage UserStorage) model.ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		userID := update.Message.From.ID
		chatID := update.Message.Chat.ID

		// Check user data .if user is registered, check user role
		role, err := userStorage.CheckUserRole(ctx, userID)
		if err == nil && role != "" {
			msg := tgbotapi.NewMessage(chatID, fmt.Sprintf(
				"👋 С возвращением! Ваша роль: %s\nИспользуйте /help для списка команд",
				role,
			))
			_, err = bot.Send(msg)
			return err
		}

		user, err := userStorage.GetUser(ctx, userID)
		if err != nil {
			return fmt.Errorf("ошибка получения пользователя: %w", err)
		}

		if user.State != StateRegistered {
			switch user.State {
			case StateAwaitName:
				msg := tgbotapi.NewMessage(chatID, "✏️ Продолжим регистрацию. Введите ваше ФИО:")
				_, err = bot.Send(msg)
				return err
			case StateAwaitBtrx:
				msg := tgbotapi.NewMessage(chatID, "✏️ Пожалуйста, введите ваш Bitrix24 ID (или 'нет' если отсутствует):")
				_, err = bot.Send(msg)
				return err
			}
		}

		newUser := User{
			TelegramID: userID,
			State:      StateAwaitName,
		}

		if err := userStorage.SaveUser(ctx, newUser); err != nil {
			return fmt.Errorf("ошибка сохранения пользователя: %w", err)
		}

		msg := tgbotapi.NewMessage(chatID, "👋 Добро пожаловать! Для начала работы заполните данные.\n\n📝 Введите ваше ФИО:")
		_, err = bot.Send(msg)
		return err
	}
}

// NameHandler
func NameHandler(userStorage UserStorage) model.ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		if update.Message == nil || update.Message.Text == "" {
			return nil
		}

		userID := update.Message.From.ID
		chatID := update.Message.Chat.ID
		name := strings.TrimSpace(update.Message.Text)

		user, err := userStorage.GetUser(ctx, userID)
		if err != nil {
			return fmt.Errorf("ошибка получения пользователя: %w", err)
		}

		if user == nil || user.State != StateAwaitName {
			msg := tgbotapi.NewMessage(chatID, "ℹ️ Пожалуйста, начните регистрацию с команды /start")
			_, err = bot.Send(msg)
			return err
		}

		if len(name) < 5 || !strings.Contains(name, " ") {
			msg := tgbotapi.NewMessage(chatID, "❌ Неверный формат ФИО. Введите полное имя (минимум 2 слова):")
			_, err = bot.Send(msg)
			return err
		}

		user.Data.Name = name
		user.State = StateAwaitBtrx

		if err := userStorage.SaveUser(ctx, *user); err != nil {
			return err
		}

		msg := tgbotapi.NewMessage(chatID, "✅ ФИО сохранено!\n\nТеперь введите ваш Bitrix24 ID (или 'нет' если отсутствует):")
		_, err = bot.Send(msg)
		return err
	}
}

// BitrixHandler
func BitrixHandler(userStorage UserStorage) model.ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		if update.Message == nil || update.Message.Text == "" {
			return nil
		}

		userID := update.Message.From.ID
		chatID := update.Message.Chat.ID
		input := strings.TrimSpace(update.Message.Text)

		user, err := userStorage.GetUser(ctx, userID)
		if err != nil {
			return fmt.Errorf("ошибка получения пользователя: %w", err)
		}

		if user == nil || user.State != StateAwaitBtrx {
			msg := tgbotapi.NewMessage(chatID, "ℹ️ Пожалуйста, начните регистрацию с команды /start")
			_, err = bot.Send(msg)
			return err
		}

		btrxID := ""
		if strings.ToLower(input) != "нет" {
			btrxID = input
		}

		user.Data.BtrxID = btrxID
		user.State = StateRegistered

		if err := userStorage.SaveUser(ctx, *user); err != nil {
			return err
		}

		confirmation := fmt.Sprintf("✅ Регистрация завершена!\n\n👤 ФИО: %s", user.Data.Name)
		if btrxID != "" {
			confirmation += fmt.Sprintf("\n🆔 Bitrix24 ID: %s", btrxID)
		}
		confirmation += fmt.Sprintf("\n👤 Ваша роль: %s\n\nТеперь вам доступны все функции бота.", user.Role)

		msg := tgbotapi.NewMessage(chatID, confirmation)
		_, err = bot.Send(msg)
		return err
	}
}

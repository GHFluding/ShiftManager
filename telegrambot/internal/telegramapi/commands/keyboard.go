package commands

import (
	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api"
)

func createInlineKeyboard[T any](
	items []T,
	prefix string,
	textFunc func(item T) string,
	callBackFunc func(item T) string,
) tgBotAPI.InlineKeyboardMarkup {
	var rows []tgBotAPI.InlineKeyboardButton
	for _, item := range items {
		btn := tgBotAPI.NewInlineKeyboardButtonData(
			textFunc(item),
			callBackFunc(item),
		)
		rows = append(rows, btn)
	}
	return tgBotAPI.NewInlineKeyboardMarkup(tgBotAPI.NewInlineKeyboardRow(rows...))
}

package commands

import (
	"context"
	"telegramSM/internal/telegramapi/model"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api"
)

const emptyInt = 0

type Router struct {
	commandHandlers  map[model.CommandType]model.ViewFunc
	messageHandlers  []model.ViewFunc
	callbackHandlers []model.ViewFunc
}

func NewRouter() *Router {
	return &Router{
		commandHandlers: make(map[model.CommandType]model.ViewFunc),
		messageHandlers: make([]model.ViewFunc, 0),
	}
}

func (r *Router) RegisterCommandHandler(cmd model.CommandType, handler model.ViewFunc) {
	r.commandHandlers[cmd] = handler
}

func (r *Router) RegisterMessageHandler(handler model.ViewFunc) {
	r.messageHandlers = append(r.messageHandlers, handler)
}

func (r *Router) RegisterCallbackHandler(handler model.ViewFunc) {
	r.callbackHandlers = append(r.callbackHandlers, handler)
}

func (r *Router) HandleUpdate(ctx context.Context, bot *tgBotAPI.BotAPI, update tgBotAPI.Update) error {
	if update.CallbackQuery != nil {
		for _, handler := range r.callbackHandlers {
			if err := handler(ctx, bot, update); err != nil {
				return err
			}
		}
		return nil
	}

	if update.Message != nil && update.Message.IsCommand() {
		cmd := model.CommandType(update.Message.Command())
		if handler, exists := r.commandHandlers[cmd]; exists {
			return handler(ctx, bot, update)
		}
	}

	if update.Message != nil {
		for _, handler := range r.messageHandlers {
			if err := handler(ctx, bot, update); err != nil {
				return err
			}
		}
	}

	return nil
}

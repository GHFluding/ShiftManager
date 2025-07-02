package commands

import (
	"context"
	"telegramSM/internal/telegramapi/model"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Router struct {
	commandHandlers map[model.CommandType]model.ViewFunc
	messageHandlers []model.ViewFunc
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

func (r *Router) HandleUpdate(ctx context.Context, bot *tgBotAPI.BotAPI, update tgBotAPI.Update) error {
	if update.Message != nil && update.Message.IsCommand() {
		cmd := model.CommandType(update.Message.Command())
		if handler, exists := r.commandHandlers[cmd]; exists {
			return handler(ctx, bot, update)
		}
	}

	for _, handler := range r.messageHandlers {
		if err := handler(ctx, bot, update); err != nil {
			return err
		}
	}

	return nil
}

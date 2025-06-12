package model

import (
	"context"
	"log/slog"
	"runtime/debug"
	sl "telegramSM/internal/util/logger"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	api      *tgbotapi.BotAPI
	cmdViews map[string]ViewFunc
	log      *slog.Logger
}

type ViewFunc func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error

func New(api *tgbotapi.BotAPI, log *slog.Logger) *Bot {
	return &Bot{
		api: api,
		log: log,
	}
}

func (b *Bot) RegisterCmdView(cmd string, view ViewFunc) {
	if b.cmdViews == nil {
		b.cmdViews = make(map[string]ViewFunc)
	}
	b.cmdViews[cmd] = view
}

func (b *Bot) handleUpdate(ctx context.Context, update tgbotapi.Update) {
	defer func() {
		if p := recover(); p != nil {
			b.log.Info("panic recovered %v\n%s", p, slog.String("error: ", string(debug.Stack())))
		}
	}()

	if update.Message.IsCommand() {
		return
	}

	cmd := update.Message.Command()

	cmdView, ok := b.cmdViews[cmd]
	if !ok {
		return
	}

	if err := cmdView(ctx, b.api, update); err != nil {
		b.log.Info("bot command: %s, is failed", cmd, " \n with Error: ", sl.Err(err))

		if _, err := b.api.Send(
			tgbotapi.NewMessage(update.Message.Chat.ID, "internal error"),
		); err != nil {
			b.log.Info("Failed error info message", sl.Err(err))
		}

	}
}
func (b *Bot) Run(ctx context.Context) error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.api.GetUpdatesChan(u)
	if err != nil {
		b.log.Info("failed to run bot", sl.Err(err))
		return err
	}

	for {
		select {
		case update := <-updates:
			updateCtx, updateCancel := context.WithTimeout(context.Background(), 5*time.Minute)
			b.handleUpdate(updateCtx, update)
			updateCancel()
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

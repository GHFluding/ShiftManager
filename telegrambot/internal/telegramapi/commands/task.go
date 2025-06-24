package commands

import (
	"context"
	"fmt"
	"strings"
	"telegramSM/internal/telegramapi/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Task struct {
	Machineid    int64
	Shiftid      int64
	Frequency    string
	Taskpriority string
	Description  string
	Createdby    int64
}

type TaskService interface {
	GetTaskToday(ctx context.Context, telegramID int) (*Task, error)
	SaveTask(ctx context.Context, task *Task) error
}

func CreateTaskHandler(taskService TaskService) model.ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		text := update.Message.Text
		chatID := update.Message.Chat.ID
		userID := int64(update.Message.From.ID)

		task, err := parseTaskInput(text, userID)
		if err != nil {
			msg := tgbotapi.NewMessage(chatID,
				"❌ Неверный формат задачи. Используйте:\n"+
					"Машина: [ID]\n"+
					"Смена: [ID]\n"+
					"Частота: [частота]\n"+
					"Приоритет: [приоритет]\n"+
					"Описание: [текст]")
			_, err = bot.Send(msg)
			return err
		}

		if err := taskService.SaveTask(ctx, task); err != nil {
			errorMsg := fmt.Sprintf("❌ Ошибка сохранения задачи: %v", err)
			msg := tgbotapi.NewMessage(chatID, errorMsg)
			_, err = bot.Send(msg)
			return err
		}

		confirmation := fmt.Sprintf(
			"✅ Задача создана!\n\n"+
				"Машина: %d\n"+
				"Смена: %d\n"+
				"Частота: %s\n"+
				"Приоритет: %s\n"+
				"Описание: %s",
			task.Machineid,
			task.Shiftid,
			task.Frequency,
			task.Taskpriority,
			task.Description,
		)

		msg := tgbotapi.NewMessage(chatID, confirmation)
		_, err = bot.Send(msg)
		return err
	}
}

// parseTaskInput parsing message to Task
func parseTaskInput(input string, createdBy int64) (*Task, error) {
	lines := strings.Split(input, "\n")
	task := &Task{Createdby: createdBy}

	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) < 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch strings.ToLower(key) {
		case "машина":
			_, err := fmt.Sscanf(value, "%d", &task.Machineid)
			if err != nil {
				return nil, fmt.Errorf("неверный ID машины")
			}
		case "смена":
			_, err := fmt.Sscanf(value, "%d", &task.Shiftid)
			if err != nil {
				return nil, fmt.Errorf("неверный ID смены")
			}
		case "частота":
			task.Frequency = value
		case "приоритет":
			task.Taskpriority = value
		case "описание":
			task.Description = value
		}
	}

	if task.Machineid == emptyInt || task.Shiftid == emptyInt || task.Description == emptyString {
		return nil, fmt.Errorf("не заполнены обязательные поля")
	}

	return task, nil
}

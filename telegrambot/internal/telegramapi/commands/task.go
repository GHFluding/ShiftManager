package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"telegramSM/internal/telegramapi/model"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api"
)

type MachineService interface {
	ListMachines(ctx context.Context) ([]Machine, error)
}

type Machine struct {
	ID   int64
	Name string
}

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

const (
	CallbackTypeMachine = "machine"
	CallbackTypeShift   = "shift"
)

type CallbackData struct {
	Type      string `json:"type"`
	MachineID int64  `json:"machine_id,omitempty"`
	ShiftID   int64  `json:"shift_id,omitempty"`
}

func SerializeCallbackData(data CallbackData) (string, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func ParseCallbackData(input string) (CallbackData, error) {
	var data CallbackData
	if err := json.Unmarshal([]byte(input), &data); err != nil {
		return CallbackData{}, err
	}
	return data, nil
}

func CreateTaskHandler(taskService TaskService, machineService MachineService, shiftService ShiftService) model.ViewFunc {
	return func(ctx context.Context, bot *tgBotAPI.BotAPI, update tgBotAPI.Update) error {
		chatID := update.Message.Chat.ID

		return showTaskMachineSelection(ctx, bot, chatID, machineService)
	}
}

func showTaskMachineSelection(
	ctx context.Context,
	bot *tgBotAPI.BotAPI,
	chatID int64,
	machineService MachineService,
) error {
	machines, err := machineService.ListMachines(ctx)
	if err != nil {
		return err
	}
	keyboard := createInlineKeyboard(
		machines,
		func(m Machine) string { return m.Name },
		func(m Machine) string {
			data, _ := SerializeCallbackData(CallbackData{
				Type:      CallbackTypeMachine,
				MachineID: m.ID,
			})
			return data
		},
	)

	msg := tgBotAPI.NewMessage(chatID, "Выберите машину:")
	msg.ReplyMarkup = keyboard
	_, err = bot.Send(msg)
	return err
}

func TaskCallbackHandler(
	machineService MachineService,
	shiftService ShiftService,
	taskService TaskService,
) model.ViewFunc {
	return func(ctx context.Context, bot *tgBotAPI.BotAPI, update tgBotAPI.Update) error {
		callback := update.CallbackQuery
		if callback == nil {
			return nil
		}

		data, err := ParseCallbackData(callback.Data)
		if err != nil {
			return nil
		}

		chatID := callback.Message.Chat.ID
		messageID := callback.Message.MessageID

		switch data.Type {
		case CallbackTypeMachine:
			return showShiftSelection(ctx, bot, chatID, messageID, shiftService, taskService)

		case CallbackTypeShift:
			edit := tgBotAPI.NewEditMessageText(chatID, messageID, "✅ Машина и смена выбраны")
			bot.Send(edit)

			msg := tgBotAPI.NewMessage(chatID,
				"Введите остальные данные задачи в формате:\n"+
					"Частота: [частота]\n"+
					"Приоритет: [приоритет]\n"+
					"Описание: [текст]")
			_, err = bot.Send(msg)
			return err
		}

		return nil
	}
}

func showShiftSelection(
	ctx context.Context,
	bot *tgBotAPI.BotAPI,
	chatID int64,
	messageID int,
	shiftService ShiftService,
	taskService TaskService,
) error {
	shifts, err := shiftService.ListShifts(ctx)
	if err != nil {
		return err
	}

	keyboard := createInlineKeyboard(
		shifts,
		func(s Shift) string { return fmt.Sprintf("shift_master_%d", s.ShiftMasterID) },
		func(s Shift) string {
			return fmt.Sprintf("shift_machine_%d", s.Machineid)
		},
	)

	edit := tgBotAPI.NewEditMessageText(chatID, messageID, "Выберите смену:")
	edit.ReplyMarkup = &keyboard
	_, err = bot.Send(edit)
	return err
}

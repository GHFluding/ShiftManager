package commands

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"telegramSM/internal/telegramapi/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Shift struct {
	Machineid     int64
	ShiftMasterID int64
}

type MachineIconService interface {
	ListMachines(ctx context.Context) ([]MachineIcon, error)
}

type MasterIconService interface {
	ListMasters(ctx context.Context) ([]MasterIcon, error)
}

type MachineIcon struct {
	ID   int64
	Name string
}

type MasterIcon struct {
	ID   int64
	Name string
}

type ShiftService interface {
	SaveShift(ctx context.Context, shift *Shift) error
}

type ShiftCreationState struct {
	MachineID     int64
	ShiftMasterID int64
}
type createShiftState int

const (
	stateMachine  = 1
	stateMaster   = 2
	stateComplete = 3
)

func (s *ShiftCreationState) CurrentStep() createShiftState {
	if s.MachineID == emptyInt {
		return stateMachine
	} else {
		if s.ShiftMasterID == emptyInt {
			return stateMaster
		}
	}
	return stateComplete
}

var shiftStates = make(map[int64]*ShiftCreationState)

const (
	callbackMachinePrefix = "machine_"
	callbackMasterPrefix  = "master_"
)

func CreateShiftHandler(
	shiftSvc ShiftService,
	machineService MachineIconService,
	masterService MasterIconService,
) model.ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		chatID := update.Message.Chat.ID
		userID := update.Message.From.ID

		state, exists := shiftStates[int64(userID)]
		if !exists {
			state = &ShiftCreationState{}
			shiftStates[int64(userID)] = state
		}

		currentStep := state.CurrentStep()

		switch currentStep {
		case stateMachine:
			return handleMachineStep(ctx, bot, update, state, machineService, masterService, chatID)
		case stateMaster:
			return handleMasterStep(ctx, bot, update, state, masterService, chatID)
		case stateComplete:
			msg := tgbotapi.NewMessage(chatID,
				"✅ Смена уже создана!\n"+
					"Используйте /newshift для создания новой смены")
			_, err := bot.Send(msg)
			return err
		default:
			msg := tgbotapi.NewMessage(chatID,
				"🚫 Неизвестное состояние. Начните заново с /createshift")
			_, err := bot.Send(msg)
			return err
		}
	}
}

func handleMachineStep(
	ctx context.Context,
	bot *tgbotapi.BotAPI,
	update tgbotapi.Update,
	state *ShiftCreationState,
	machineService MachineIconService,
	masterService MasterIconService,
	chatID int64,
) error {
	if update.Message.Text == "/createshift" {
		return showMachineSelection(ctx, bot, chatID, machineService)
	}

	machineID, err := strconv.ParseInt(update.Message.Text, 10, 64)
	if err != nil {
		msg := tgbotapi.NewMessage(chatID,
			"❌ Неверный формат ID машины. Введите число или выберите из списка.")
		_, err := bot.Send(msg)
		return err
	}

	state.MachineID = machineID
	return showMasterSelection(ctx, bot, chatID, masterService)
}

func handleMasterStep(
	ctx context.Context,
	bot *tgbotapi.BotAPI,
	update tgbotapi.Update,
	state *ShiftCreationState,
	masterService MasterIconService,
	chatID int64,
) error {
	if update.Message.Text == "" {
		return showMasterSelection(ctx, bot, chatID, masterService)
	}

	masterID, err := strconv.ParseInt(update.Message.Text, 10, 64)
	if err != nil {
		msg := tgbotapi.NewMessage(chatID,
			"❌ Неверный формат ID мастера. Введите число или выберите из списка.")
		_, err := bot.Send(msg)
		return err
	}

	state.ShiftMasterID = masterID
	return nil
}

func showMachineSelection(
	ctx context.Context,
	bot *tgbotapi.BotAPI,
	chatID int64,
	machineService MachineIconService,
) error {
	machines, err := machineService.ListMachines(ctx)
	if err != nil {
		msg := tgbotapi.NewMessage(chatID, "❌ Ошибка получения списка машин")
		_, err := bot.Send(msg)
		return err
	}

	if len(machines) == 0 {
		msg := tgbotapi.NewMessage(chatID, "Нет доступных машин")
		_, err := bot.Send(msg)
		return err
	}

	var rows []tgbotapi.InlineKeyboardButton
	for _, machine := range machines {
		btn := tgbotapi.NewInlineKeyboardButtonData(
			fmt.Sprintf("%s (ID: %d)", machine.Name, machine.ID),
			callbackMachinePrefix+strconv.FormatInt(machine.ID, 10),
		)
		rows = append(rows, btn)
	}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(rows...))

	msg := tgbotapi.NewMessage(chatID, "Выберите машину:")
	msg.ReplyMarkup = keyboard
	_, err = bot.Send(msg)
	return err
}

func showMasterSelection(
	ctx context.Context,
	bot *tgbotapi.BotAPI,
	chatID int64,
	masterService MasterIconService,
) error {
	masters, err := masterService.ListMasters(ctx)
	if err != nil {
		msg := tgbotapi.NewMessage(chatID, "❌ Ошибка получения списка мастеров")
		_, err := bot.Send(msg)
		return err
	}

	if len(masters) == 0 {
		msg := tgbotapi.NewMessage(chatID, "Нет доступных мастеров")
		_, err := bot.Send(msg)
		return err
	}

	var rows []tgbotapi.InlineKeyboardButton
	for _, master := range masters {
		btn := tgbotapi.NewInlineKeyboardButtonData(
			fmt.Sprintf("%s (ID: %d)", master.Name, master.ID),
			callbackMasterPrefix+strconv.FormatInt(master.ID, 10),
		)
		rows = append(rows, btn)
	}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(rows...))

	msg := tgbotapi.NewMessage(chatID, "Выберите мастера смены:")
	msg.ReplyMarkup = keyboard
	_, err = bot.Send(msg)
	return err
}

func ShiftCallbackHandler(
	shiftSvc ShiftService,
	machineService MachineIconService,
	masterService MasterIconService,
) model.ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		callback := update.CallbackQuery
		if callback == nil {
			return nil
		}

		chatID := callback.Message.Chat.ID
		userID := callback.From.ID
		data := callback.Data

		state, exists := shiftStates[int64(userID)]
		if !exists {
			state = &ShiftCreationState{}
			shiftStates[int64(userID)] = state
		}

		currentStep := state.CurrentStep()

		switch {
		case strings.HasPrefix(data, callbackMachinePrefix) && currentStep == stateMachine:
			return handleMachineCallback(ctx, bot, callback, state, machineService, masterService, chatID)

		case strings.HasPrefix(data, callbackMasterPrefix) && currentStep == stateMaster:
			return handleMasterCallback(ctx, bot, callback, state, shiftSvc, chatID, int64(userID))
		}

		return nil
	}
}

func handleMachineCallback(
	ctx context.Context,
	bot *tgbotapi.BotAPI,
	callback *tgbotapi.CallbackQuery,
	state *ShiftCreationState,
	machineService MachineIconService,
	masterService MasterIconService,
	chatID int64,
) error {
	data := callback.Data
	machineID, err := strconv.ParseInt(strings.TrimPrefix(data, callbackMachinePrefix), 10, 64)
	if err != nil {
		return err
	}

	state.MachineID = machineID

	edit := tgbotapi.NewEditMessageText(
		chatID,
		callback.Message.MessageID,
		"✅ Машина выбрана: "+getButtonText(callback.Message),
	)
	bot.Send(edit)

	return showMasterSelection(ctx, bot, chatID, masterService)
}

func handleMasterCallback(
	ctx context.Context,
	bot *tgbotapi.BotAPI,
	callback *tgbotapi.CallbackQuery,
	state *ShiftCreationState,
	shiftSvc ShiftService,
	chatID int64,
	userID int64,
) error {
	data := callback.Data
	masterID, err := strconv.ParseInt(strings.TrimPrefix(data, callbackMasterPrefix), 10, 64)
	if err != nil {
		return err
	}

	state.ShiftMasterID = masterID

	edit := tgbotapi.NewEditMessageText(
		chatID,
		callback.Message.MessageID,
		"✅ Мастер выбран: "+getButtonText(callback.Message),
	)
	bot.Send(edit)

	shift := &Shift{
		Machineid:     state.MachineID,
		ShiftMasterID: state.ShiftMasterID,
	}

	if err := shiftSvc.SaveShift(ctx, shift); err != nil {
		errorMsg := fmt.Sprintf("❌ Ошибка сохранения смены: %v", err)
		msg := tgbotapi.NewMessage(chatID, errorMsg)
		bot.Send(msg)
		return err
	}

	confirmation := fmt.Sprintf(
		"✅ Смена создана!\n\n"+
			"Машина: %d\n"+
			"Мастер смены: %d",
		shift.Machineid,
		shift.ShiftMasterID,
	)

	msg := tgbotapi.NewMessage(chatID, confirmation)
	bot.Send(msg)

	delete(shiftStates, userID)
	return nil
}

func getButtonText(msg *tgbotapi.Message) string {
	if msg == nil {
		return ""
	}

	return msg.Text
}

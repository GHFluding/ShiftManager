package main

import (
	"context"
	"log"
	"os"
	service_mock "telegramSM/internal/services/mock"
	"telegramSM/internal/telegramapi/commands"
	"telegramSM/internal/telegramapi/model"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	//TODO: init config
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	bot, err := tgBotAPI.NewBotAPI(botToken)
	if err != nil {
		log.Panicf("Ошибка инициализации бота: %v", err)
	}
	bot.Debug = true

	router := commands.NewRouter()
	//TODO: init interfaces
	userService := service_mock.UserServiceMock{}
	taskService := service_mock.TaskServiceMock{}
	machineService := service_mock.MachineServiceMock{}
	shiftService := service_mock.ShiftServiceMock{}
	masterService := service_mock.MasterServiceMock{}
	router.RegisterCommandHandler(model.CmdStart, commands.StartHandler(userService))
	router.RegisterCommandHandler(model.CmdHelp, commands.HelpHandler(userService))
	router.RegisterCommandHandler(model.CmdCreateTask, commands.CreateTaskHandler(taskService, machineService, shiftService))
	router.RegisterCommandHandler(model.CmdCreateShift, commands.CreateShiftHandler(shiftService, machineService, masterService))
	//TODO: refactor callback register function
	router.RegisterMessageHandler(commands.NameHandler(userService))
	router.RegisterMessageHandler(commands.SkipBitrixHandler(userService))
	router.RegisterMessageHandler(commands.TaskCallbackHandler(machineService, shiftService, taskService))

	u := tgBotAPI.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panicf("Error getting update: %v", err)
	}

	ctx := context.Background()
	for update := range updates {
		if err := router.HandleUpdate(ctx, bot, update); err != nil {
			log.Printf("Error handle update: %v", err)
		}
	}
}

package bot_command

import "log/slog"

func Help(log *slog.Logger) string {
	//standard message to bitrix user in chat, have all command for this bot
	const commandList = `/create-task - создает задание. список аргументов: 
	`
	log.Info("Help request was successful.")
	return commandList
}

package bot_command

import "log/slog"

func Help(log *slog.Logger) string {
	const commandList = `/create-task - создает задание. список аргументов: 
	`
	log.Info("Help request was successful.")
	return commandList
}

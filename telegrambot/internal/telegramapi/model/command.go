package model

type Command string

const (
	CmdStart        Command = "/start"
	CmdHelp         Command = "/help"
	CmdTaskComplete Command = "/task-complete"
)

func GetDescription(command Command) string {
	switch command {
	case CmdStart:
		return "команда для авторизации, либо регестрации"
	case CmdTaskComplete:
		return "Команда отмечает задание как сделанное"
	}
	return "неизвестная команда"
}

var AdminCommands = [3]Command{CmdStart,
	CmdHelp,
	CmdTaskComplete}

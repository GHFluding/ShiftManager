package model

type CommandType string

const (
	CmdStart      CommandType = "start"
	CmdHelp       CommandType = "help"
	CmdCreateTask CommandType = "createtask"
)

type CommandMeta struct {
	Type        CommandType
	Description string
	MinRole     string
}

var (
	AdminCommands = []CommandMeta{
		{CmdStart, "Начало работы", "admin"},
		{CmdHelp, "Помощь", "admin"},
		{CmdCreateTask, "Создать задачу", "admin"},
	}

	MasterCommands = []CommandMeta{
		{CmdStart, "Начало работы", "admin"},
		{CmdHelp, "Помощь", "admin"},
		{CmdCreateTask, "Создать задачу", "admin"},
	}
	ManagerCommands = []CommandMeta{
		{CmdStart, "Начало работы", "admin"},
		{CmdHelp, "Помощь", "admin"},
		{CmdCreateTask, "Создать задачу", "admin"},
	}
	WorkerCommands = []CommandMeta{
		{CmdStart, "Начало работы", "admin"},
		{CmdHelp, "Помощь", "admin"},
		{CmdCreateTask, "Создать задачу", "admin"},
	}
)

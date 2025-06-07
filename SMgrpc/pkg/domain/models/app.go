package models

type App struct {
	ID   int
	Name string
}

type DBFunction struct {
	Machine MachineDB
	User    UserDB
	Task    TaskDB
	Shift   ShiftDB
}

type CommandCode int

const (
	MachineServer CommandCode = iota
	UserServer
	TaskServer
	ShiftServer
)

var commandName = map[CommandCode]string{
	MachineServer: "machine",
	UserServer:    "user",
	TaskServer:    "task",
	ShiftServer:   "shift",
}

func (cs CommandCode) String() string {
	return commandName[cs]
}

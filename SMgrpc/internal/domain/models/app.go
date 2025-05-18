package models

import (
	"smgrpc/internal/grpc/sm/machine"
	"smgrpc/internal/grpc/sm/shift"
	"smgrpc/internal/grpc/sm/task"
	"smgrpc/internal/grpc/sm/user"
)

type App struct {
	ID   int
	Name string
}

type DBFunction struct {
	Machine machine.MachineInterface
	User    user.UserInterface
	Task    task.TaskInterface
	Shift   shift.ShiftInterface
}

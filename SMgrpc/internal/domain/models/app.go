package models

import (
	"github.com/GHFluding/ShiftManager/SMgrpc/internal/grpc/sm/shift"
	"github.com/GHFluding/ShiftManager/SMgrpc/internal/grpc/sm/task"
	"github.com/GHFluding/ShiftManager/SMgrpc/internal/grpc/sm/user"
)

type App struct {
	ID   int
	Name string
}

type DBFunction struct {
	Machine MachineDB
	User    user.UserInterface
	Task    task.TaskInterface
	Shift   shift.ShiftInterface
}

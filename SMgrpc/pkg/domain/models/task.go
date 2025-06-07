package models

import "context"

type Task struct {
	MachineId    int64
	ShiftId      int64
	Frequency    string
	TaskPriority string
	Description  string
}

type TaskSaver interface {
	SaveTask(
		ctx context.Context,
		machineId int64,
		shiftId int64,
		frequency string,
		taskPriority string,
		description string,
	) (
		id int64,
		err error,
	)
}
type TaskProvider interface {
	GetTask(ctx context.Context, id int64) (Task, error)
}

type TaskDB struct {
	Saver    TaskSaver
	Provider TaskProvider
}

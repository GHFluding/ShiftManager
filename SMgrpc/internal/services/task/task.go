package shift_service

import (
	"context"
	"log/slog"
	"smgrpc/internal/domain/models"
	sl "smgrpc/internal/utils/logger"
)

type TaskApp struct {
	log      *slog.Logger
	saver    TaskSaver
	provider TaskProvider
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
	//maybe id, not machine id and shift id
	Task(ctx context.Context, machineId int64, shiftId int64) (models.Task, error)
}

func New(log *slog.Logger, taskSaver TaskSaver, taskProvider TaskProvider) *TaskApp {
	return &TaskApp{
		log:      log,
		saver:    taskSaver,
		provider: taskProvider,
	}
}

func (t *TaskApp) Create(ctx context.Context,
	machineId int64,
	shiftId int64,
	frequency string,
	taskPriority string,
	description string,
) (
	int64,
	int64,
	string,
	string,
	string,
	error,
) {
	const op = "task.Create"
	log := t.log.With(
		slog.String("op", op),
		slog.Int64("shift id", shiftId),
	)
	log.Info("creating task")
	id, err := t.saver.SaveTask(ctx, machineId, shiftId, frequency, taskPriority, description)
	if err != nil {
		log.Error("failed to create task", sl.Err(err))
		return machineId, shiftId, frequency, taskPriority, description, err
	}
	log.Info("task is created", slog.Int64("id", id))

	task, err := t.provider.Task(ctx, machineId, shiftId)
	if err != nil {
		log.Error("failed to check task", sl.Err(err))
		return machineId, shiftId, frequency, taskPriority, description, err
	}

	return task.MachineId, task.ShiftId, task.Frequency, task.TaskPriority, task.Description, nil
}

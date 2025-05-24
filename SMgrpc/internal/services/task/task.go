package shift_service

import (
	"context"
	"log/slog"

	task_grpc "github.com/GHFluding/ShiftManager/SMgrpc/internal/grpc/sm/task"
	sl "github.com/GHFluding/ShiftManager/SMgrpc/internal/utils/logger"
	"github.com/GHFluding/ShiftManager/SMgrpc/pkg/domain/models"
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
	Task(ctx context.Context, id int64) (models.Task, error)
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
	task_grpc.TaskResponse,
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
		return task_grpc.TaskResponse{}, err
	}
	log.Info("task is created", slog.Int64("id", id))

	task, err := t.provider.Task(ctx, id)
	if err != nil {
		log.Error("failed to check task", sl.Err(err))
		return task_grpc.TaskResponse{}, err
	}
	taskResponse := task_grpc.TaskResponse{
		MachineId:    task.MachineId,
		ShiftId:      task.ShiftId,
		Frequency:    task.Frequency,
		TaskPriority: task.TaskPriority,
		Description:  task.Description,
	}
	return taskResponse, nil
}

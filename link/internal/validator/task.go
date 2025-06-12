package validator

import (
	"encoding/json"
	"fmt"
	"log/slog"

	entities "github.com/GHFluding/ShiftManager/SMgrpc/pkg/gen"
	logger "github.com/GHFluding/ShiftManager/link/internal/utils"
)

type taskDefault struct {
	Machineid    int64
	Shiftid      int64
	Frequency    string
	Taskpriority string
	Description  string
	Createdby    int64
}

func (t taskDefault) ToGRPCCreateParams() *entities.CreateTaskParams {
	return &entities.CreateTaskParams{
		MachineId:    t.Machineid,
		ShiftId:      t.Shiftid,
		Frequency:    t.Frequency,
		TaskPriority: t.Taskpriority,
		Description:  t.Description,
		CreatedBy:    t.Createdby,
	}
}

func Task(data []byte, log *slog.Logger) (taskDefault, error) {
	task, err := marshalTask(data, log)
	if err != nil {
		return taskDefault{}, err
	}
	return task, err
}

func marshalTask(data []byte, log *slog.Logger) (taskDefault, error) {
	var task taskDefault
	if err := json.Unmarshal(data, &task); err != nil {
		log.Error("JSON unmarshal error", logger.ErrToAttr(err))
		return task, fmt.Errorf("invalid request format: %w", err)
	}

	log.Info("Parsed task data",
		slog.Int64("machineid", task.Machineid),
		slog.Int64("shiftid", task.Shiftid),
		slog.String("frequency", task.Frequency),
		slog.String("taskpriority", task.Taskpriority),
		slog.String("description", task.Description),
		slog.Int64("createdby", task.Createdby))

	return task, nil
}

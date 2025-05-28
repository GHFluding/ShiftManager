package grpc_storage

import (
	"context"
	"time"

	"github.com/GHFluding/ShiftManager/SM/internal/database/postgres"
	"github.com/GHFluding/ShiftManager/SM/internal/utils/convertor"
	"github.com/GHFluding/ShiftManager/SMgrpc/pkg/domain/models"
)

func (sp *Services) SaveTask(
	ctx context.Context,
	machineId int64,
	shiftId int64,
	frequency string,
	taskPriority string,
	description string,
) (
	id int64,
	err error,
) {
	taskDB, err := convertTaskParams(machineId, shiftId, frequency, taskPriority, description)
	if err != nil {
		return 0, err
	}
	taskResp, err := sp.db.CreateTask(context.Background(), taskDB)
	return taskResp.ID, err
}

func (sp *Services) GetTask(ctx context.Context, id int64) (models.Task, error) {
	taskDB, err := sp.db.GetTask(context.Background(), id)
	if err != nil {
		return models.Task{}, err
	}
	task := models.Task{
		MachineId:    taskDB.Machineid,
		ShiftId:      taskDB.Shiftid.Int64,
		Frequency:    string(taskDB.Frequency),
		Description:  taskDB.Description,
		TaskPriority: string(taskDB.Taskpriority),
	}
	return task, nil
}

func convertTaskParams(machineId int64,
	shiftId int64,
	frequency string,
	taskPriority string,
	description string) (postgres.CreateTaskParams, error) {
	var taskDB postgres.CreateTaskParams
	taskDB.Machineid = machineId
	taskDB.Shiftid = convertor.PGInt64(&shiftId)
	err := taskDB.Frequency.Scan(frequency)
	if err != nil {
		return taskDB, err
	}
	err = taskDB.Taskpriority.Scan(taskPriority)
	if err != nil {
		return taskDB, err
	}
	taskDB.Description = description
	taskDB.Createdby = time.Now().UnixMicro()
	return taskDB, nil
}

package services

import (
	"context"
	"sm/internal/database/postgres"
	"sm/internal/services/logic"
	"sm/internal/utils/logger"

	"github.com/jackc/pgx/v5/pgtype"
)

func CreateTask(sp *ServicesParams, req Task) (Task, error) {
	taskParams := convertCreateTaskParams(req)
	taskDB, err := sp.db.CreateTask(context.Background(), taskParams)
	if err != nil {
		sp.log.Info("Failed to crete task: ", logger.ErrToAttr(err))
		return Task{}, err
	}
	if taskDB.Shiftid.Valid {
		logic.AddTaskToShift(sp.log, sp.db, taskDB.ID, taskDB.Shiftid.Int64)
	} else {
		logic.AddTaskToShiftWithMachine(sp.log, sp.db, taskDB.Machineid, taskDB.ID)
	}
	task := convertTaskDB(taskDB)
	return task, nil
}

func convertCreateTaskParams(req Task) postgres.CreateTaskParams {
	var shiftid pgtype.Int8
	if req.Shiftid == 0 {
		shiftid = pgtype.Int8{
			Valid: false,
			Int64: 0,
		}
	} else {
		shiftid = pgtype.Int8{
			Valid: true,
			Int64: req.Shiftid,
		}
	}
	return postgres.CreateTaskParams{
		ID:           req.ID,
		Machineid:    req.Machineid,
		Shiftid:      shiftid,
		Frequency:    postgres.Taskfrequency(req.Frequency),
		Taskpriority: postgres.Taskpriority(req.Taskpriority),
		Description:  req.Description,
		Createdby:    req.Createdby,
	}
}

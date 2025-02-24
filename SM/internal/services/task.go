package services

import (
	"context"
	"sm/internal/database/postgres"
	"sm/internal/services/logic"
	"sm/internal/utils/logger"

	"github.com/jackc/pgx/v5/pgtype"
)

type UpdateTaskParams struct {
	UserID  int64
	Comment string
	Status  string
}

func UpdateTask(sp *ServicesParams, reqId int64, reqParams UpdateTaskParams) error {
	userValid := true
	if reqParams.UserID == 0 {
		userValid = false
	}
	commentValid := true
	if reqParams.Comment == "" {
		commentValid = false
	}
	updateParams := postgres.UpdateTaskStatusParams{
		Taskid: reqId,
		Status: postgres.Taskstatus(reqParams.Status),
		Comment: pgtype.Text{
			String: reqParams.Status,
			Valid:  commentValid,
		},
		Userid: pgtype.Int8{
			Int64: reqParams.UserID,
			Valid: userValid,
		},
	}
	err := sp.db.UpdateTaskStatus(context.Background(), updateParams)
	return err
}

func DeleteTask(sp *ServicesParams, id int64) error {
	err := sp.db.DeleteTask(context.Background(), id)
	return err
}

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

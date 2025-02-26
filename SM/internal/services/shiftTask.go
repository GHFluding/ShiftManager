package services

import (
	"context"
	"sm/internal/database/postgres"
)

func DeleteShiftTask(sp *ServicesParams, reqTaskID int64, reqShiftId int64) error {
	req := postgres.DeleteShiftTaskParams{
		Shiftid: reqShiftId,
		Taskid:  reqTaskID,
	}
	err := sp.db.DeleteShiftTask(context.Background(), req)
	return err
}

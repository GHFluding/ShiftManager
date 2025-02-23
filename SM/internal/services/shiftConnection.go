package services

import (
	"context"
	"sm/internal/utils/logger"
)

func ShiftWorkersList(sp *ServicesParams, shiftid int64) ([]ShiftWorker, error) {
	var ShiftsWorkers []ShiftWorker
	shiftWorkersDB, err := sp.db.ShiftWorkersList(context.Background(), shiftid)
	if err != nil {
		sp.log.Info("Failed to convert shift workers from db", logger.ErrToAttr(err))
		return ShiftsWorkers, err
	}

	for _, i := range shiftWorkersDB {
		ShiftsWorkers = append(ShiftsWorkers, convertShiftWorkerDB(i))
	}
	return ShiftsWorkers, nil
}

func ShiftTasksList(sp *ServicesParams, shiftid int64) ([]ShiftTask, error) {
	var shiftTask []ShiftTask
	shiftTasksDB, err := sp.db.ShiftTasksList(context.Background(), shiftid)
	if err != nil {
		sp.log.Info("Failed to convert shift tasks from db", logger.ErrToAttr(err))
		return shiftTask, err
	}
	for _, i := range shiftTasksDB {
		shiftTask = append(shiftTask, convertShiftTaskDB(i))
	}
	return shiftTask, nil
}

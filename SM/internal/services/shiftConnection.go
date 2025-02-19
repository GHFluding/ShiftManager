package services

import (
	"context"
	"sm/internal/utils/logger"
)

func ShiftWorkersList(sp *ServicesParams, shiftid int64) ([]ShiftWorker, error) {
	var ShiftsWorkersToOut []ShiftWorker
	shiftWorkers, err := sp.db.ShiftWorkersList(context.Background(), shiftid)
	if err != nil {
		sp.log.Info("Failed to convert shift workers from db", logger.ErrToAttr(err))
		return ShiftsWorkersToOut, err
	}

	for _, i := range shiftWorkers {
		ShiftsWorkersToOut = append(ShiftsWorkersToOut, convertShiftWorker(i))
	}
	return ShiftsWorkersToOut, nil
}

func ShiftTasksList(sp *ServicesParams, shiftid int64) ([]ShiftTask, error) {
	var ShiftTaskToOut []ShiftTask
	shiftTasks, err := sp.db.ShiftTasksList(context.Background(), shiftid)
	if err != nil {
		sp.log.Info("Failed to convert shift tasks from db", logger.ErrToAttr(err))
		return ShiftTaskToOut, err
	}
	for _, i := range shiftTasks {
		ShiftTaskToOut = append(ShiftTaskToOut, convertShiftTask(i))
	}
	return ShiftTaskToOut, nil
}

package services

import (
	"context"
	"sm/internal/database/postgres"
	"sm/internal/utils/logger"
)

type ShiftWorkersListToTransfer struct {
	Valid          bool
	WorkersListDTO []ShiftWorkerDTO
}

type ShiftTasksListToTransfer struct {
	Valid        bool
	TasksListDTO []ShiftTaskDTO
}

func ShiftWorkersList(sp *ServicesParams, shiftid int64) ShiftWorkersListToTransfer {
	shiftWorkers, err := sp.db.ShiftWorkersList(context.Background(), shiftid)
	if err != nil {
		sp.log.Info("Failed to convert shift workers from db", logger.ErrToAttr(err))
		return ShiftWorkersListToTransfer{Valid: false}
	}
	shiftsWorkersDTO, err := convertListToTransport[postgres.ShiftWorker, ShiftWorkerDTO](shiftWorkers)
	if err != nil {
		sp.log.Info("Failed to convert shift workers from db", logger.ErrToAttr(err))
		return ShiftWorkersListToTransfer{Valid: false}
	}
	return ShiftWorkersListToTransfer{Valid: true, WorkersListDTO: shiftsWorkersDTO}
}

func ShiftTasksList(sp *ServicesParams, shiftid int64) ShiftTasksListToTransfer {
	shiftTasks, err := sp.db.ShiftTasksList(context.Background(), shiftid)
	if err != nil {
		sp.log.Info("Failed to convert shift tasks from db", logger.ErrToAttr(err))
		return ShiftTasksListToTransfer{Valid: false}
	}
	shiftTasksDTO, err := convertListToTransport[postgres.ShiftTask, ShiftTaskDTO](shiftTasks)
	if err != nil {
		sp.log.Info("Failed to convert shift tasks from db", logger.ErrToAttr(err))
		return ShiftTasksListToTransfer{Valid: false}
	}
	return ShiftTasksListToTransfer{Valid: true, TasksListDTO: shiftTasksDTO}
}

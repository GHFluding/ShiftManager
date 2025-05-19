package logic

import (
	"context"
	"log/slog"

	"github.com/GHFluding/ShiftManager/SM/internal/database/postgres"
	"github.com/GHFluding/ShiftManager/SM/internal/utils/logger"
)

func AddTaskToShift(log *slog.Logger, db *postgres.Queries, reqTaskId int64, reqShiftID int64) {
	//TODO: out to transport
	_, err := db.AddShiftTask(context.Background(), postgres.AddShiftTaskParams{Shiftid: reqShiftID, Taskid: reqTaskId})
	if err != nil {
		log.Info("Auto add task to shift failed: ", logger.ErrToAttr(err))
	}
	log.Info("Auto add task to shift completed")
}

func AddTaskToShiftWithMachine(log *slog.Logger, db *postgres.Queries, reqMachineID int64, reqTaskId int64) {
	//TODO: out to transport
	shifts, err := db.ActiveShiftList(context.Background())
	if err != nil {
		log.Info("Auto add task to shift with machine failed: ", logger.ErrToAttr(err))
	}
	for _, i := range shifts {
		if i.Machineid == reqMachineID {
			_, err := db.AddShiftTask(context.Background(), postgres.AddShiftTaskParams{Shiftid: i.ID, Taskid: reqTaskId})
			if err != nil {
				log.Info("Auto add task to shift with machine failed: ", logger.ErrToAttr(err))
			}
		}
	}
	log.Info("Auto add task to shift completed")
}

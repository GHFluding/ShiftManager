package services

import (
	"context"
	"sm/internal/database/postgres"
	"sm/internal/utils/logger"
)

func AddShiftWorker(sp *ServicesParams, req ShiftWorker) (ShiftWorker, error) {
	shiftWorkerParams := postgres.AddShiftWorkerParams{
		Shiftid: req.Shiftid,
		Userid:  req.Userid,
	}
	shiftWorkerDB, err := sp.db.AddShiftWorker(context.Background(), shiftWorkerParams)
	if err != nil {
		sp.log.Info("Failed to add shift worker: ", logger.ErrToAttr(err))
		return ShiftWorker{}, err
	}
	shiftWorker := convertShiftWorkerDB(shiftWorkerDB)
	return shiftWorker, nil
}

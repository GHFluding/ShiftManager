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

func DeleteShiftWorker(sp *ServicesParams, reqUserID int64, reqShiftId int64) error {
	req := postgres.DeleteShiftWorkerParams{
		Shiftid: reqShiftId,
		Userid:  reqUserID,
	}
	err := sp.db.DeleteShiftWorker(context.Background(), req)
	return err
}

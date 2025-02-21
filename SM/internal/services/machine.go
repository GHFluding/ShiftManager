package services

import (
	"context"
	"sm/internal/utils/logger"
)

func MachineNeedRepair(sp *ServicesParams, machineid int64) bool {
	err := sp.db.MachineNeedRepair(context.Background(), machineid)
	if err != nil {
		sp.log.Info("Failed to change status to need repair from db", logger.ErrToAttr(err))
		return false
	}
	return true
}

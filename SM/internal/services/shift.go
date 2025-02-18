package services

import (
	"context"
	"sm/internal/database/postgres"
	"sm/internal/utils/logger"
)

type ShiftListToTransfer struct {
	Valid        bool
	ShiftListDTO []ShiftDTO
}

func ShiftList(sp *ServicesParams) ShiftListToTransfer {
	shifts, err := sp.db.ShiftList(context.Background())
	if err != nil {
		sp.log.Info("Failed to convert shifts from db", logger.ErrToAttr(err))
		return ShiftListToTransfer{Valid: false}
	}
	shiftsDTO, err := convertListToTransport[postgres.Shift, ShiftDTO](shifts)
	if err != nil {
		sp.log.Info("Failed to convert shifts from db", logger.ErrToAttr(err))
		return ShiftListToTransfer{Valid: false}
	}
	return ShiftListToTransfer{Valid: true, ShiftListDTO: shiftsDTO}
}

func ActiveShiftList(sp *ServicesParams) ShiftListToTransfer {
	shifts, err := sp.db.ActiveShiftList(context.Background())
	if err != nil {
		sp.log.Info("Failed to convert shifts from db", logger.ErrToAttr(err))
		return ShiftListToTransfer{Valid: false}
	}
	shiftsDTO, err := convertListToTransport[postgres.Shift, ShiftDTO](shifts)
	if err != nil {
		sp.log.Info("Failed to convert shifts from db", logger.ErrToAttr(err))
		return ShiftListToTransfer{Valid: false}
	}
	return ShiftListToTransfer{Valid: true, ShiftListDTO: shiftsDTO}
}

package services

import (
	"context"
	"sm/internal/utils/logger"
)

func ShiftList(sp *ServicesParams) ([]Shift, error) {
	//here for return blank struct if error
	var ShiftsToOut []Shift
	shifts, err := sp.db.ShiftList(context.Background())
	if err != nil {
		sp.log.Info("Failed to convert shifts from db", logger.ErrToAttr(err))
		return ShiftsToOut, err
	}
	for _, i := range shifts {
		ShiftsToOut = append(ShiftsToOut, convertShift(i))
	}
	return ShiftsToOut, nil
}

func ActiveShiftList(sp *ServicesParams) ([]Shift, error) {
	//here for return blank struct if error
	var ShiftsToOut []Shift
	shifts, err := sp.db.ActiveShiftList(context.Background())
	if err != nil {
		sp.log.Info("Failed to convert shifts from db", logger.ErrToAttr(err))
		return ShiftsToOut, err
	}
	for _, i := range shifts {
		ShiftsToOut = append(ShiftsToOut, convertShift(i))
	}

	return ShiftsToOut, nil
}

package services

import (
	"context"
	"sm/internal/utils/logger"
)

func ShiftList(sp *ServicesParams) ([]Shift, error) {
	//here for return blank struct if error
	var shifts []Shift
	shiftsDB, err := sp.db.ShiftList(context.Background())
	if err != nil {
		sp.log.Info("Failed to convert shifts from db", logger.ErrToAttr(err))
		return shifts, err
	}
	for _, i := range shiftsDB {
		shifts = append(shifts, convertShift(i))
	}
	return shifts, nil
}

func ActiveShiftList(sp *ServicesParams) ([]Shift, error) {
	//here for return blank struct if error
	var shiftsToOut []Shift
	shifts, err := sp.db.ActiveShiftList(context.Background())
	if err != nil {
		sp.log.Info("Failed to convert shifts from db", logger.ErrToAttr(err))
		return shiftsToOut, err
	}
	for _, i := range shifts {
		shiftsToOut = append(shiftsToOut, convertShift(i))
	}

	return shiftsToOut, nil
}

package services

import (
	"context"
	"sm/internal/database/postgres"
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
		shifts = append(shifts, convertShiftDB(i))
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
		shiftsToOut = append(shiftsToOut, convertShiftDB(i))
	}

	return shiftsToOut, nil
}

func CreateShift(sp *ServicesParams, req Shift) (Shift, error) {
	shiftParams := convertCreateShiftParams(req)
	shiftDB, err := sp.db.CreateShift(context.Background(), shiftParams)
	if err != nil {
		sp.log.Info("Failed to create shift: ", logger.ErrToAttr(err))
		return Shift{}, err
	}
	shift := convertShiftDB(shiftDB)
	return shift, nil
}

func convertCreateShiftParams(req Shift) postgres.CreateShiftParams {
	return postgres.CreateShiftParams{
		ID:          req.ID,
		Machineid:   req.Machineid,
		ShiftMaster: req.ShiftMaster,
	}
}

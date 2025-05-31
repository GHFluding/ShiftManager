package models

import "context"

type Shift struct {
	MachineId     int64
	ShiftMasterID int64
}

type ShiftSaver interface {
	SaveShift(
		ctx context.Context,
		machineId int64,
		shiftMasterID int64,
	) (
		id int64,
		err error,
	)
}
type ShiftProvider interface {
	GETShift(ctx context.Context, id int64) (Shift, error)
}

type ShiftDB struct {
	Saver    ShiftSaver
	Provider ShiftProvider
}

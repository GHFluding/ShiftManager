package grpc_storage

import (
	"context"

	"github.com/GHFluding/ShiftManager/SM/internal/database/postgres"
	"github.com/GHFluding/ShiftManager/SM/internal/utils/logger"
	"github.com/GHFluding/ShiftManager/SMgrpc/pkg/domain/models"
)

func (sp *Services) SaveShift(
	ctx context.Context,
	machineId int64,
	shiftMasterID int64,
) (
	id int64,
	err error,
) {
	shiftParams := postgres.CreateShiftParams{
		Machineid:   machineId,
		ShiftMaster: shiftMasterID,
	}
	shiftDB, err := sp.db.CreateShift(context.Background(), shiftParams)
	if err != nil {
		sp.log.Info("Failed to create shift: ", logger.ErrToAttr(err))
		return 0, err
	}
	id = shiftDB.ID
	return id, nil
}

func (sp *Services) GETShift(ctx context.Context, id int64) (models.Shift, error) {
	shift, err := sp.db.GetShift(context.Background(), id)
	if err != nil {
		return models.Shift{}, err
	}
	return models.Shift{
		MachineId:     shift.Machineid,
		ShiftMasterID: shift.ShiftMaster,
	}, nil
}

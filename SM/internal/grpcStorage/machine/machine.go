package machine_grpc

import (
	"context"
	"log/slog"

	"github.com/GHFluding/ShiftManager/SM/internal/database/postgres"
	"github.com/GHFluding/ShiftManager/SM/internal/utils/convertor"
	"github.com/GHFluding/ShiftManager/SM/internal/utils/logger"
	"github.com/GHFluding/ShiftManager/SMgrpc/pkg/domain/models"
)

type Services struct {
	db  *postgres.Queries
	log *slog.Logger
}

func (sp *Services) SaveMachine(
	ctx context.Context,
	name string,
	isRepairRequired *bool,
	isActive *bool,
) (
	id int64,
	err error,
) {
	machineParams := convertMachineParams(name, isRepairRequired, isActive)
	machineDB, err := sp.db.CreateMachine(context.Background(), machineParams)
	if err != nil {
		sp.log.Info("Failed to create machine: ", logger.ErrToAttr(err))
		return 0, err
	}
	id = machineDB.ID
	return id, nil
}

// Machine release interface Provide for grpc service machine
func (sp *Services) GetMachine(ctx context.Context, id int64) (models.Machine, error) {
	machine, err := sp.db.GetMachine(context.Background(), id)
	if err != nil {
		return models.Machine{}, err
	}
	return models.Machine{
		Name:             machine.Name,
		IsRepairRequired: &machine.Isrepairrequired.Bool,
		IsActive:         &machine.Isactive.Bool,
	}, nil
}

func convertMachineParams(name string,
	isRepairRequired *bool,
	isActive *bool) postgres.CreateMachineParams {
	machineParams := postgres.CreateMachineParams{
		Name:             name,
		Isrepairrequired: convertor.PGBool(isRepairRequired),
		Isactive:         convertor.PGBool(isActive),
	}
	return machineParams
}

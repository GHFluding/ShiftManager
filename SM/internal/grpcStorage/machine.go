package machine

import (
	"context"
	"log/slog"

	"github.com/GHFluding/ShiftManager/SM/internal/database/postgres"
	"github.com/GHFluding/ShiftManager/SM/internal/utils/logger"
	"github.com/jackc/pgx/v5/pgtype"
)

type Services struct {
	db  *postgres.Queries
	log *slog.Logger
}

func (sp *Services) ParamsSaveMachine(
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

func convertMachineParams(name string,
	isRepairRequired *bool,
	isActive *bool) postgres.CreateMachineParams {
	var isRepairRequiredPG pgtype.Bool
	if isRepairRequired == nil {
		isRepairRequiredPG = pgtype.Bool{
			Valid: false,
		}
	} else {
		isRepairRequiredPG = pgtype.Bool{
			Bool:  true,
			Valid: false,
		}
	}
	var isActivePG pgtype.Bool
	if isActive == nil {
		isActivePG = pgtype.Bool{
			Valid: false,
		}
	} else {
		isActivePG = pgtype.Bool{
			Bool:  true,
			Valid: false,
		}
	}
	machineParams := postgres.CreateMachineParams{
		Name:             name,
		Isrepairrequired: isRepairRequiredPG,
		Isactive:         isActivePG,
	}
	return machineParams
}

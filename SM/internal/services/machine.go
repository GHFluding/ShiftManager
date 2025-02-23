package services

import (
	"context"
	"sm/internal/database/postgres"
	"sm/internal/utils/logger"

	"github.com/jackc/pgx/v5/pgtype"
)

func MachineNeedRepair(sp *ServicesParams, machineid int64) error {
	err := sp.db.MachineNeedRepair(context.Background(), machineid)
	if err != nil {
		sp.log.Info("Failed to change status to need repair from db", logger.ErrToAttr(err))
		return err
	}
	return nil
}

func CreateMachine(sp *ServicesParams, req Machine) (Machine, error) {
	machineParams := convertCreateMachineParams(req)
	machineDB, err := sp.db.CreateMachine(context.Background(), machineParams)
	if err != nil {
		sp.log.Info("Failed to create machine: ", logger.ErrToAttr(err))
		return Machine{}, err
	}
	machine := convertMachineDB(machineDB)
	return machine, nil
}
func convertCreateMachineParams(req Machine) postgres.CreateMachineParams {
	return postgres.CreateMachineParams{
		ID:   req.ID,
		Name: req.Name,
		Isrepairrequired: pgtype.Bool{
			Valid: true,
			Bool:  req.Isrepairrequired,
		},
		Isactive: pgtype.Bool{
			Valid: true,
			Bool:  req.Isactive,
		},
	}
}

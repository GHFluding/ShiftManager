package services

import (
	"context"
	"sm/internal/database/postgres"
	"sm/internal/utils/logger"
)

func MachineNeedRepair(sp *ServicesParams, machineid int64) error {
	err := sp.db.MachineNeedRepair(context.Background(), machineid)
	if err != nil {
		sp.log.Info("Failed to change status to need repair from db", logger.ErrToAttr(err))
		return err
	}
	return nil
}

func CreateMachine(sp *ServicesParams, req Machine) error {
	machineParams := convertCreateMachineParams(req)
	_, err := sp.db.CreateMachine(context.Background(), machineParams)
	if err != nil {
		return err
	}
	return nil
}
func convertCreateMachineParams(req Machine) postgres.CreateMachineParams {
	var machineParams postgres.CreateMachineParams
	machineParams.ID = req.ID
	machineParams.Name = req.Name
	machineParams.Isrepairrequired.Valid = true
	machineParams.Isrepairrequired.Bool = req.Isrepairrequired
	machineParams.Isactive.Valid = true
	machineParams.Isactive.Bool = req.Isactive
	return machineParams
}

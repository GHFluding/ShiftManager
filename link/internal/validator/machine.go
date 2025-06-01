package validator

import (
	"encoding/json"
	"fmt"
	"log/slog"

	entities "github.com/GHFluding/ShiftManager/SMgrpc/pkg/gen"
	logger "github.com/GHFluding/ShiftManager/link/internal/utils"
)

type MachineCode int

const (
	Create = iota
)

type machineDefault struct {
	Name             string
	IsRepairRequired *bool
	IsActive         *bool
}

func Machine(command MachineCode, data []byte, log *slog.Logger) (any, error) {
	machine, err := marshalCreateMachine(data, log)
	if err != nil {
		return nil, err
	}
	switch command {
	case Create:
		return &entities.CreateMachine{
			Name:             machine.Name,
			IsRepairRequired: machine.IsRepairRequired,
			IsActive:         machine.IsActive,
		}, err
	default:
		return nil, err
	}

}

func marshalCreateMachine(data []byte, log *slog.Logger) (machineDefault, error) {

	var machine machineDefault
	if err := json.Unmarshal(data, &machine); err != nil {
		log.Error("JSON unmarshal error", logger.ErrToAttr(err))
		return machine, fmt.Errorf("invalid request format: %w", err)
	}

	log.Info("Parsed machine data",
		slog.String("name", machine.Name),
		slog.Any("isrepairrequired", machine.IsRepairRequired),
		slog.Any("isactive", machine.IsActive))
	return machine, nil

}

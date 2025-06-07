package validator

import (
	"encoding/json"
	"fmt"
	"log/slog"

	entities "github.com/GHFluding/ShiftManager/SMgrpc/pkg/gen"
	logger "github.com/GHFluding/ShiftManager/link/internal/utils"
)

type machineDefault struct {
	Name             string
	IsRepairRequired *bool
	IsActive         *bool
}

func (m machineDefault) ToGRPCCreateParams() *entities.CreateMachine {
	return &entities.CreateMachine{
		Name:             m.Name,
		IsRepairRequired: m.IsRepairRequired,
		IsActive:         m.IsActive,
	}
}

func Machine(data []byte, log *slog.Logger) (machineDefault, error) {
	machine, err := marshalMachine(data, log)
	if err != nil {
		return machineDefault{}, err
	}
	return machine, err
}

func marshalMachine(data []byte, log *slog.Logger) (machineDefault, error) {

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

package machine_service

import (
	"context"
	"log/slog"

	machine_grpc "github.com/GHFluding/ShiftManager/SMgrpc/internal/grpc/sm/machine"
	sl "github.com/GHFluding/ShiftManager/SMgrpc/internal/utils/logger"
	"github.com/GHFluding/ShiftManager/SMgrpc/pkg/domain/models"
)

type MachineApp struct {
	log        *slog.Logger
	DBFunction models.MachineDB
}

func New(log *slog.Logger, machineDB models.MachineDB) *MachineApp {
	return &MachineApp{
		log:        log,
		DBFunction: machineDB,
	}
}

func (m *MachineApp) Create(ctx context.Context,
	name string,
	isRepairRequired *bool,
	isActive *bool,
) (machine_grpc.MachineResponse, error) {
	const op = "machine.Create"
	log := m.log.With(
		slog.String("op", op),
		slog.String("machine name", name),
	)
	log.Info("creating machine")
	id, err := m.DBFunction.Saver.SaveMachine(ctx, name, isRepairRequired, isActive)
	if err != nil {
		log.Error("failed to create machine", sl.Err(err))
		return machine_grpc.MachineResponse{}, err
	}
	log.Info("machine is created", slog.Int64("id", id))

	machine, err := m.DBFunction.Provider.GetMachine(ctx, id)
	if err != nil {
		log.Error("failed to check machine", sl.Err(err))
		return machine_grpc.MachineResponse{}, err
	}
	machineResponse := machine_grpc.MachineResponse{
		Name:             machine.Name,
		IsRepairRequired: machine.IsRepairRequired,
		IsActive:         machine.IsActive,
	}
	return machineResponse, nil
}

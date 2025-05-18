package machine_service

import (
	"context"
	"log/slog"
	"smgrpc/internal/domain/models"
	machine_grpc "smgrpc/internal/grpc/sm/machine"
	sl "smgrpc/internal/utils/logger"
)

type MachineApp struct {
	log      *slog.Logger
	saver    MachineSaver
	provider MachineProvider
}

type MachineSaver interface {
	SaveMachine(
		ctx context.Context,
		name string,
		isRepairRequired *bool,
		isActive *bool,
	) (
		id int64,
		err error,
	)
}
type MachineProvider interface {
	Machine(ctx context.Context, id int64) (models.Machine, error)
	// TODO: IsRepairRequired and isActive
}

func New(log *slog.Logger, machineSaver MachineSaver, machineProvider MachineProvider) *MachineApp {
	return &MachineApp{
		log:      log,
		saver:    machineSaver,
		provider: machineProvider,
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
	id, err := m.saver.SaveMachine(ctx, name, isRepairRequired, isActive)
	if err != nil {
		log.Error("failed to create machine", sl.Err(err))
		return machine_grpc.MachineResponse{}, err
	}
	log.Info("machine is created", slog.Int64("id", id))

	machine, err := m.provider.Machine(ctx, id)
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

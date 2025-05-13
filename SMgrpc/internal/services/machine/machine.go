package machine_service

import (
	"context"
	"log/slog"
	"smgrpc/internal/domain/models"
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
	Machine(ctx context.Context, name string) (models.Machine, error)
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
) (string, *bool, *bool, error) {
	const op = "machine.Create"
	log := m.log.With(
		slog.String("op", op),
		slog.String("machine name", name),
	)
	log.Info("creating machine")
	id, err := m.saver.SaveMachine(ctx, name, isRepairRequired, isActive)
	if err != nil {
		log.Error("failed to create machine", sl.Err(err))
		return "", nil, nil, err
	}
	log.Info("machine is created", slog.Int64("id", id))
	//TODO: check machine in db with id and return db's value
	return name, isRepairRequired, isActive, nil
}

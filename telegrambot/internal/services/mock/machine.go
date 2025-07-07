package service_mock

import (
	"context"
	"telegramSM/internal/telegramapi/commands"
)

type MachineServiceMock struct {
	machineList []commands.Machine
}

func (ms MachineServiceMock) ListMachines(ctx context.Context) ([]commands.Machine, error) {
	return ms.machineList, nil
}

package client

import (
	"context"

	entities "github.com/GHFluding/ShiftManager/SMgrpc/internal/gen"
)

func (c *Client) CreateTask(ctx context.Context, machineID, shiftID, createdBy int64, frequency, priority, description string) (*entities.TaskResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req := &entities.CreateTaskParams{
		MachineId:    machineID,
		ShiftId:      shiftID,
		CreatedBy:    createdBy,
		Frequency:    frequency,
		TaskPriority: priority,
		Description:  description,
	}

	return c.Clients.Task.Create(ctx, req)
}

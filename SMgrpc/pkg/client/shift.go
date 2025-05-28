package client

import (
	"context"
	"time"

	entities "github.com/GHFluding/ShiftManager/SMgrpc/internal/gen"
)

func (c *Client) CreateShift(ctx context.Context, machineID, shiftMaster int64) (*entities.ShiftResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	req := &entities.CreateShiftParams{
		MachineId:   machineID,
		ShiftMaster: shiftMaster,
	}

	return c.Clients.Shift.Create(ctx, req)
}

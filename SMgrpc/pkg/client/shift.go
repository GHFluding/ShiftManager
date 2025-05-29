package client

import (
	"context"

	entities "github.com/GHFluding/ShiftManager/SMgrpc/internal/gen"
)

func (c *Client) CreateShift(ctx context.Context, machineID, shiftMasterID int64) (*entities.ShiftResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req := &entities.CreateShiftParams{
		MachineId:   machineID,
		ShiftMaster: shiftMasterID,
	}

	return c.Clients.Shift.Create(ctx, req)
}

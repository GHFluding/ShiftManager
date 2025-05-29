package client

import (
	"context"

	entities "github.com/GHFluding/ShiftManager/SMgrpc/internal/gen"
)

func (c *Client) CreateMachine(ctx context.Context, name string, isRepairRequired, isActive *bool) (*entities.MachineResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req := &entities.CreateMachine{
		Name:             name,
		IsRepairRequired: isRepairRequired,
		IsActive:         isActive,
	}

	return c.Clients.Machine.Create(ctx, req)
}

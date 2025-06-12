package client

import (
	"context"

	entities "github.com/GHFluding/ShiftManager/SMgrpc/pkg/gen"
)

func (c *Client) CreateMachine(ctx context.Context, req *entities.CreateMachine) (*entities.MachineResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return c.Clients.Machine.Create(ctx, req)
}

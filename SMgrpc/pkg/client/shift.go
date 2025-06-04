package client

import (
	"context"

	entities "github.com/GHFluding/ShiftManager/SMgrpc/pkg/gen"
)

func (c *Client) CreateShift(ctx context.Context, req *entities.CreateShiftParams) (*entities.ShiftResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return c.Clients.Shift.Create(ctx, req)
}

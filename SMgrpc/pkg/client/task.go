package client

import (
	"context"

	entities "github.com/GHFluding/ShiftManager/SMgrpc/pkg/gen"
)

func (c *Client) CreateTask(ctx context.Context, req *entities.CreateTaskParams) (*entities.TaskResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return c.Clients.Task.Create(ctx, req)
}

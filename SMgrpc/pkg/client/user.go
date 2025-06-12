package client

import (
	"context"

	entities "github.com/GHFluding/ShiftManager/SMgrpc/pkg/gen"
)

func (c *Client) CreateUser(ctx context.Context, req *entities.CreateUserParams) (*entities.UserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return c.Clients.User.Create(ctx, req)
}

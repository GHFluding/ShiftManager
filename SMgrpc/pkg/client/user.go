package client

import (
	"context"

	entities "github.com/GHFluding/ShiftManager/SMgrpc/pkg/gen"
)

func (c *Client) CreateUser(ctx context.Context, name, role, telegramID string, bitrixID *int64) (*entities.UserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req := &entities.CreateUserParams{
		Name:       name,
		Role:       role,
		TelegramId: telegramID,
		BitrixId:   bitrixID,
	}

	return c.Clients.User.Create(ctx, req)
}

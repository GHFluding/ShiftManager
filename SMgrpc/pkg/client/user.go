package client

import (
	"context"
	"time"

	entities "github.com/GHFluding/ShiftManager/SMgrpc/internal/gen"
)

func (c *Client) CreateUser(ctx context.Context, name, role, telegramID string, bitrixID *int64) (*entities.UserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	req := &entities.CreateUserParams{
		Name:       name,
		Role:       role,
		TelegramId: telegramID,
		BitrixId:   bitrixID,
	}

	return c.Clients.User.Create(ctx, req)
}

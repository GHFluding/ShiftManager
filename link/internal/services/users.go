package services

import (
	"context"
	"encoding/json"
	"fmt"

	"log/slog"

	"github.com/GHFluding/ShiftManager/SMgrpc/pkg/client"
	entities "github.com/GHFluding/ShiftManager/SMgrpc/pkg/gen"
	logger "github.com/GHFluding/ShiftManager/link/internal/utils"
)

type createUserParams struct {
	Bitrixid   *int64 `json:"bitrixid,omitempty"`
	TelegramID string `json:"telegramid"`
	Name       string `json:"name"`
	Role       string `json:"role"`
}

func CreateUserGRPC(c *client.Client, data *entities.CreateUserParams, log *slog.Logger) (*entities.UserResponse, error) {
	log.Info("Start processing user creation request")

	if data.BitrixId == nil {
		log.Info("BitrixID is not set. This user use only telegram")
	}

	resp, err := c.CreateUser(context.Background(), data)
	if err != nil {
		log.Error("GRPC request failed", logger.ErrToAttr(err))
		return nil, fmt.Errorf("service unavailable: %w", err)
	}

	return resp, nil
}

func marshalCreateUser(data []byte, log *slog.Logger) (createUserParams, error) {
	var user createUserParams
	if err := json.Unmarshal(data, &user); err != nil {
		log.Error("JSON unmarshal error", logger.ErrToAttr(err))
		return user, fmt.Errorf("invalid request format: %w", err)
	}

	log.Info("Parsed user data",
		slog.String("name", user.Name),
		slog.String("role", user.Role),
		slog.String("telegramid", user.TelegramID),
		slog.Any("bitrixid", user.Bitrixid))
	return user, nil
}

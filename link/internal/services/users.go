package services

import (
	"context"
	"encoding/json"
	"fmt"

	"log/slog"

	"github.com/GHFluding/ShiftManager/SMgrpc/pkg/client"
	logger "github.com/GHFluding/ShiftManager/link/internal/utils"
)

type createUserParams struct {
	Bitrixid   *int64 `json:"bitrixid,omitempty"`
	TelegramID string `json:"telegramid"`
	Name       string `json:"name"`
	Role       string `json:"role"`
}

func CreateUser(c *client.Client, data []byte, log *slog.Logger, url string) ([]byte, error) {
	log.Info("Start processing user creation request")
	user, err := marshalCreateUser(data, log)
	if err != nil {
		return nil, err
	}

	if user.Bitrixid == nil {
		log.Info("BitrixID is not set. This user use only telegram")
	}

	resp, err := c.CreateUser(context.Background(), user.Name, user.Role, user.TelegramID, user.Bitrixid)
	if err != nil {
		log.Error("GRPC request failed", logger.ErrToAttr(err))
		return nil, fmt.Errorf("service unavailable: %w", err)
	}

	//using only response data for marshaling
	responseData, err := json.Marshal(resp.Data)
	if err != nil {
		log.Error("Failed to marshal response", logger.ErrToAttr(err))
		return nil, fmt.Errorf("response marshal failed: %w", err)
	}

	return responseData, nil
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

package validator

import (
	"encoding/json"
	"fmt"
	"log/slog"

	entities "github.com/GHFluding/ShiftManager/SMgrpc/pkg/gen"
	logger "github.com/GHFluding/ShiftManager/link/internal/utils"
)

type userDefault struct {
	Bitrixid   *int64
	TelegramID string
	Name       string
	Role       string
}

func (u userDefault) ToGRPCCreateParams() *entities.CreateUserParams {
	return &entities.CreateUserParams{
		Name:       u.Name,
		BitrixId:   u.Bitrixid,
		TelegramId: u.TelegramID,
		Role:       u.Role,
	}
}

func User(data []byte, log *slog.Logger) (userDefault, error) {
	user, err := marshalUser(data, log)
	if err != nil {
		return userDefault{}, err
	}
	return user, err
}

func marshalUser(data []byte, log *slog.Logger) (userDefault, error) {
	var user userDefault
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

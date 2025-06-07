package user_service

import (
	"context"
	"log/slog"

	user_grpc "github.com/GHFluding/ShiftManager/SMgrpc/internal/grpc/sm/user"
	sl "github.com/GHFluding/ShiftManager/SMgrpc/internal/utils/logger"
	"github.com/GHFluding/ShiftManager/SMgrpc/pkg/domain/models"
)

type UserApp struct {
	log      *slog.Logger
	saver    models.UserSaver
	provider models.UserProvider
}

func New(log *slog.Logger, userSaver models.UserSaver, userProvider models.UserProvider) *UserApp {
	return &UserApp{
		log:      log,
		saver:    userSaver,
		provider: userProvider,
	}
}

func (s *UserApp) Create(ctx context.Context,
	bitrixId *int64,
	telegramId string,
	name string,
	role string,
) (
	user_grpc.UserResponse,
	error,
) {
	const op = "user.Create"
	log := s.log.With(
		slog.String("op", op),
		slog.String("Name ", name),
		slog.String("Role ", role),
	)
	log.Info("creating user")
	id, err := s.saver.SaveUser(ctx, bitrixId, telegramId, name, role)
	if err != nil {
		log.Error("failed to create user", sl.Err(err))
		return user_grpc.UserResponse{}, err
	}
	log.Info("user is created", slog.Int64("id", id))

	user, err := s.provider.GetUser(ctx, id)
	if err != nil {
		log.Error("failed to check user", sl.Err(err))
		return user_grpc.UserResponse{}, err
	}
	userResponse := user_grpc.UserResponse{
		BitrixId:   user.BitrixId,
		TelegramId: user.TelegramId,
		Name:       user.Name,
		Role:       user.Role,
	}
	return userResponse, nil
}

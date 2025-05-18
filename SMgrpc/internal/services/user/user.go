package user_service

import (
	"context"
	"log/slog"
	"smgrpc/internal/domain/models"
	user_grpc "smgrpc/internal/grpc/sm/user"
	sl "smgrpc/internal/utils/logger"
)

type UserApp struct {
	log      *slog.Logger
	saver    UserSaver
	provider UserProvider
}

type UserSaver interface {
	SaveUser(
		ctx context.Context,
		bitrixId *int64,
		telegramId string,
		name string,
		role string,
	) (
		id int64,
		err error,
	)
}
type UserProvider interface {
	User(ctx context.Context, id int64) (models.User, error)
}

func New(log *slog.Logger, userSaver UserSaver, userProvider UserProvider) *UserApp {
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

	user, err := s.provider.User(ctx, id)
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

package user_grpc

import (
	"context"
	"log/slog"

	"github.com/GHFluding/ShiftManager/SM/internal/database/postgres"
	"github.com/GHFluding/ShiftManager/SM/internal/utils/convertor"
	"github.com/GHFluding/ShiftManager/SM/internal/utils/logger"
	"github.com/GHFluding/ShiftManager/SMgrpc/pkg/domain/models"
)

type Services struct {
	db  *postgres.Queries
	log *slog.Logger
}

func (sp *Services) SaveUser(
	ctx context.Context,
	bitrixId *int64,
	telegramId string,
	name string,
	role string,
) (
	id int64,
	err error,
) {
	userParams, err := convertUserParams(bitrixId, telegramId, name, role)
	if err != nil {
		return 0, err
	}
	userDB, err := sp.db.CreateUser(context.Background(), userParams)
	if err != nil {
		sp.log.Info("Failed to create machine: ", logger.ErrToAttr(err))
		return 0, err
	}
	id = userDB.ID
	return id, nil
}

// Machine release interface Provide for grpc service machine
func (sp *Services) GetUser(ctx context.Context, id int64) (models.User, error) {
	user, err := sp.db.GetUser(context.Background(), id)
	if err != nil {
		return models.User{}, err
	}
	var BitrixIdInt64 *int64
	if user.Bitrixid.Valid {
		BitrixIdInt64 = &user.Bitrixid.Int64
	} else {
		BitrixIdInt64 = nil
	}
	return models.User{
		Name:       user.Name,
		BitrixId:   BitrixIdInt64,
		Role:       string(user.Role),
		TelegramId: user.Telegramid,
	}, nil
}

func convertUserParams(bitrixId *int64,
	telegramId string,
	name string,
	role string) (postgres.CreateUserParams, error) {
	var roleEnum postgres.Userrole
	err := roleEnum.Scan(role)
	if err != nil {
		return postgres.CreateUserParams{}, err
	}
	userParams := postgres.CreateUserParams{
		Bitrixid:   convertor.PGInt64(bitrixId),
		Name:       name,
		Role:       roleEnum,
		Telegramid: telegramId,
	}
	return userParams, nil
}

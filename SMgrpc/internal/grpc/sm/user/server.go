package user

import (
	"context"

	entities "github.com/GHFluding/ShiftManager/SMgrpc/internal/gen"
	"github.com/GHFluding/ShiftManager/SMgrpc/pkg/domain/models"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserResponse struct {
	BitrixId   *int64
	TelegramId string
	Name       string
	Role       string
}

type serverAPI struct {
	entities.UnimplementedUserServiceServer
	user models.UserDB
}

func RegisterServerAPI(gRPC *grpc.Server, user models.UserDB) {
	entities.RegisterUserServiceServer(gRPC, &serverAPI{user: user})
}
func (s *serverAPI) Create(ctx context.Context, req *entities.CreateUserParams) (*entities.UserResponse, error) {
	if req.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "name is empty")
	}
	if req.GetRole() == "" {
		return nil, status.Error(codes.InvalidArgument, "name is empty")
	}
	userID, err := s.user.Saver.SaveUser(ctx, req.BitrixId, req.TelegramId, req.Name, req.Role)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	user, err := s.user.Provider.GetUser(ctx, userID)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &entities.UserResponse{
		Data: &entities.CreateUserParams{
			BitrixId:   user.BitrixId,
			TelegramId: user.TelegramId,
			Name:       user.Name,
			Role:       user.Role,
		},
	}, nil
}

package user

import (
	"context"
	entities "smgrpc/internal/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserInterface interface {
	Create(ctx context.Context,
		BitrixId *int64,
		TelegramId string,
		Name string,
		Role string,
	) (
		*int64,
		string,
		string,
		string,
		error,
	)
}
type serverAPI struct {
	entities.UnimplementedUserServiceServer
	user UserInterface
}

func RegisterServerAPI(gRPC *grpc.Server, user UserInterface) {
	entities.RegisterUserServiceServer(gRPC, &serverAPI{user: user})
}
func (s *serverAPI) Create(ctx context.Context, req *entities.CreateUserParams) (*entities.UserResponse, error) {
	if req.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "name is empty")
	}
	if req.GetRole() == "" {
		return nil, status.Error(codes.InvalidArgument, "name is empty")
	}
	bitrixId, telegramId, name, role, err := s.user.Create(ctx, req.BitrixId, req.TelegramId, req.Name, req.Role)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &entities.UserResponse{
		Data: &entities.CreateUserParams{
			BitrixId:   bitrixId,
			TelegramId: telegramId,
			Name:       name,
			Role:       role,
		},
	}, nil
}

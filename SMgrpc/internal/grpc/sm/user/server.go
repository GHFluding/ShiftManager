package user

import (
	"context"

	entities "github.com/GHFluding/ShiftManager/SMgrpc/internal/gen"

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
type UserInterface interface {
	Create(ctx context.Context,
		BitrixId *int64,
		TelegramId string,
		Name string,
		Role string,
	) (
		UserResponse,
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
	user, err := s.user.Create(ctx, req.BitrixId, req.TelegramId, req.Name, req.Role)
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

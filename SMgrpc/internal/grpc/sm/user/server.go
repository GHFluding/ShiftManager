package user

import (
	"context"
	entities "smgrpc/internal/gen"

	"google.golang.org/grpc"
)

type serverAPI struct {
	entities.UnimplementedUserServiceServer
}

func RegisterServerAPI(gRPC *grpc.Server) {
	entities.RegisterUserServiceServer(gRPC, &serverAPI{})
}
func (s *serverAPI) Create(ctx context.Context, req *entities.CreateUserParams) (*entities.UserResponse, error) {
	panic("implement me")
}

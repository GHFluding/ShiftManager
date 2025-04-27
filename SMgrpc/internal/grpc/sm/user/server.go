package user

import (
	"context"
	entities "smgrpc/internal/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverAPI struct {
	entities.UnimplementedUserServiceServer
}

func RegisterServerAPI(gRPC *grpc.Server) {
	entities.RegisterUserServiceServer(gRPC, &serverAPI{})
}
func (s *serverAPI) Create(ctx context.Context, req *entities.CreateUserParams) (*entities.UserResponse, error) {
	if req.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "name is empty")
	}
	if req.GetRole() == "" {
		return nil, status.Error(codes.InvalidArgument, "name is empty")
	}
	var resp entities.UserResponse
	return &resp, nil
}

package task

import (
	"context"
	entities "smgrpc/internal/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverAPI struct {
	entities.UnimplementedTaskServiceServer
}

func RegisterServerAPI(gRPC *grpc.Server) {
	entities.RegisterTaskServiceServer(gRPC, &serverAPI{})
}
func (s *serverAPI) Create(ctx context.Context, req *entities.CreateTaskParams) (*entities.TaskResponse, error) {
	if req.GetDescription() == "" {
		return nil, status.Error(codes.InvalidArgument, "name is empty")
	}
	if req.GetFrequency() == "" {
		return nil, status.Error(codes.InvalidArgument, "name is empty")
	}
	if req.GetTaskPriority() == "" {
		return nil, status.Error(codes.InvalidArgument, "name is empty")
	}
	var resp entities.TaskResponse
	return &resp, nil
}

package task

import (
	"context"
	entities "smgrpc/internal/gen"

	"google.golang.org/grpc"
)

type serverAPI struct {
	entities.UnimplementedTaskServiceServer
}

func RegisterServerAPI(gRPC *grpc.Server) {
	entities.RegisterTaskServiceServer(gRPC, &serverAPI{})
}
func (s *serverAPI) Create(ctx context.Context, req *entities.CreateTaskParams) (*entities.TaskResponse, error) {
	panic("implement me")
}

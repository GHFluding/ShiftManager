package machine

import (
	"context"
	entities "smgrpc/internal/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverAPI struct {
	entities.UnimplementedMachineServiceServer
}

func RegisterServerAPI(gRPC *grpc.Server) {
	entities.RegisterMachineServiceServer(gRPC, &serverAPI{})
}
func (s *serverAPI) Create(ctx context.Context, req *entities.CreateMachine) (*entities.MachineResponse, error) {
	if req.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "name is empty")
	}
	var resp entities.MachineResponse
	return &resp, nil
}

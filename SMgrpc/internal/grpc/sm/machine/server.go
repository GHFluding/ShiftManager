package machine

import (
	"context"
	entities "smgrpc/internal/gen"

	"google.golang.org/grpc"
)

type serverAPI struct {
	entities.UnimplementedMachineServiceServer
}

func RegisterServerAPI(gRPC *grpc.Server) {
	entities.RegisterMachineServiceServer(gRPC, &serverAPI{})
}
func (s *serverAPI) Create(ctx context.Context, req *entities.CreateMachine) (*entities.MachineResponse, error) {
	panic("implement me")
}

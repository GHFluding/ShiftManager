package shift

import (
	"context"
	entities "smgrpc/internal/gen"

	"google.golang.org/grpc"
)

type serverAPI struct {
	entities.UnimplementedShiftServiceServer
}

func RegisterServerAPI(gRPC *grpc.Server) {
	entities.RegisterShiftServiceServer(gRPC, &serverAPI{})
}
func (s *serverAPI) Create(ctx context.Context, req *entities.CreateShiftParams) (*entities.ShiftResponse, error) {
	panic("implement me")
}

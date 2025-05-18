package machine

import (
	"context"
	entities "smgrpc/internal/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MachineResponse struct {
	Name             string
	IsRepairRequired *bool
	IsActive         *bool
}

type MachineInterface interface {
	Create(ctx context.Context,
		name string,
		isRepairRequired *bool,
		isActive *bool,
	) (
		MachineResponse,
		error,
	)
}
type serverAPI struct {
	entities.UnimplementedMachineServiceServer
	machine MachineInterface
}

func RegisterServerAPI(gRPC *grpc.Server, machine MachineInterface) {
	entities.RegisterMachineServiceServer(gRPC, &serverAPI{machine: machine})
}
func (s *serverAPI) Create(ctx context.Context, req *entities.CreateMachine) (*entities.MachineResponse, error) {
	if req.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "name is empty")
	}
	machine, err := s.machine.Create(ctx, req.Name, req.IsRepairRequired, req.IsActive)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &entities.MachineResponse{
		Data: &entities.CreateMachine{
			Name:             machine.Name,
			IsRepairRequired: machine.IsRepairRequired,
			IsActive:         machine.IsActive,
		},
	}, nil
}

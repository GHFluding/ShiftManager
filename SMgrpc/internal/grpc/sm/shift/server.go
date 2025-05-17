package shift

import (
	"context"
	entities "smgrpc/internal/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ShiftInterface interface {
	Create(ctx context.Context,
		machineId int64,
		shiftMaster int64,
	) (
		int64,
		int64,
		error,
	)
}

type serverAPI struct {
	entities.UnimplementedShiftServiceServer
	shift ShiftInterface
}

func RegisterServerAPI(gRPC *grpc.Server, shift ShiftInterface) {
	entities.RegisterShiftServiceServer(gRPC, &serverAPI{shift: shift})
}
func (s *serverAPI) Create(ctx context.Context, req *entities.CreateShiftParams) (*entities.ShiftResponse, error) {

	machine, master, err := s.shift.Create(ctx, req.MachineId, req.ShiftMaster)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &entities.ShiftResponse{
		Data: &entities.CreateShiftParams{
			MachineId:   machine,
			ShiftMaster: master,
		},
	}, nil
}

package shift

import (
	"context"
	entities "smgrpc/internal/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ShiftResponse struct {
	MachineId     int64
	ShiftMasterID int64
}
type ShiftInterface interface {
	Create(ctx context.Context,
		machineId int64,
		shiftMasterID int64,
	) (
		ShiftResponse,
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

	shift, err := s.shift.Create(ctx, req.MachineId, req.ShiftMaster)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &entities.ShiftResponse{
		Data: &entities.CreateShiftParams{
			MachineId:   shift.MachineId,
			ShiftMaster: shift.ShiftMasterID,
		},
	}, nil
}

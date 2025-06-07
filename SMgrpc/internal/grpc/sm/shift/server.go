package shift

import (
	"context"

	entities "github.com/GHFluding/ShiftManager/SMgrpc/internal/gen"
	"github.com/GHFluding/ShiftManager/SMgrpc/pkg/domain/models"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ShiftResponse struct {
	MachineId     int64
	ShiftMasterID int64
}

type serverAPI struct {
	entities.UnimplementedShiftServiceServer
	shift models.ShiftDB
}

func RegisterServerAPI(gRPC *grpc.Server, shift models.ShiftDB) {
	entities.RegisterShiftServiceServer(gRPC, &serverAPI{shift: shift})
}
func (s *serverAPI) Create(ctx context.Context, req *entities.CreateShiftParams) (*entities.ShiftResponse, error) {

	shiftID, err := s.shift.Saver.SaveShift(ctx, req.MachineId, req.ShiftMaster)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	shift, err := s.shift.Provider.GETShift(ctx, shiftID)
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

package shift_service

import (
	"context"
	"log/slog"

	shift_grpc "github.com/GHFluding/ShiftManager/SMgrpc/internal/grpc/sm/shift"
	sl "github.com/GHFluding/ShiftManager/SMgrpc/internal/utils/logger"
	"github.com/GHFluding/ShiftManager/SMgrpc/pkg/domain/models"
)

type ShiftApp struct {
	log      *slog.Logger
	saver    models.ShiftSaver
	provider models.ShiftProvider
}

func New(log *slog.Logger, shiftSaver models.ShiftSaver, shiftProvider models.ShiftProvider) *ShiftApp {
	return &ShiftApp{
		log:      log,
		saver:    shiftSaver,
		provider: shiftProvider,
	}
}

func (s *ShiftApp) Create(ctx context.Context,
	machineId int64,
	shiftMasterID int64,
) (
	shift_grpc.ShiftResponse,
	error,
) {
	const op = "shift.Create"
	log := s.log.With(
		slog.String("op", op),
		slog.Int64("master id", shiftMasterID),
		slog.Int64("machine id", machineId),
	)
	log.Info("creating shift")
	id, err := s.saver.SaveShift(ctx, machineId, shiftMasterID)
	if err != nil {
		log.Error("failed to create shift", sl.Err(err))
		return shift_grpc.ShiftResponse{}, err
	}
	log.Info("shift is created", slog.Int64("id", id))

	shift, err := s.provider.GETShift(ctx, id)
	if err != nil {
		log.Error("failed to check shift", sl.Err(err))
		return shift_grpc.ShiftResponse{}, err
	}
	shiftResponse := shift_grpc.ShiftResponse{
		MachineId:     shift.MachineId,
		ShiftMasterID: shift.ShiftMasterID,
	}
	return shiftResponse, nil
}

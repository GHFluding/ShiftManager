package shift_service

import (
	"context"
	"log/slog"
	"smgrpc/internal/domain/models"
	shift_grpc "smgrpc/internal/grpc/sm/shift"
	sl "smgrpc/internal/utils/logger"
)

type ShiftApp struct {
	log      *slog.Logger
	saver    ShiftSaver
	provider ShiftProvider
}

type ShiftSaver interface {
	SaveShift(
		ctx context.Context,
		machineId int64,
		masterID int64,
	) (
		id int64,
		err error,
	)
}
type ShiftProvider interface {
	Shift(ctx context.Context, id int64) (models.Shift, error)
}

func New(log *slog.Logger, shiftSaver ShiftSaver, shiftProvider ShiftProvider) *ShiftApp {
	return &ShiftApp{
		log:      log,
		saver:    shiftSaver,
		provider: shiftProvider,
	}
}

func (s *ShiftApp) Create(ctx context.Context,
	machineId int64,
	masterId int64,
) (
	shift_grpc.ShiftResponse,
	error,
) {
	const op = "shift.Create"
	log := s.log.With(
		slog.String("op", op),
		slog.Int64("master id", masterId),
		slog.Int64("machine id", machineId),
	)
	log.Info("creating shift")
	id, err := s.saver.SaveShift(ctx, machineId, masterId)
	if err != nil {
		log.Error("failed to create shift", sl.Err(err))
		return shift_grpc.ShiftResponse{}, err
	}
	log.Info("shift is created", slog.Int64("id", id))

	shift, err := s.provider.Shift(ctx, id)
	if err != nil {
		log.Error("failed to check shift", sl.Err(err))
		return shift_grpc.ShiftResponse{}, err
	}
	shiftResponse := shift_grpc.ShiftResponse{
		MachineId:   shift.MachineId,
		ShiftMaster: shift.ShiftMaster,
	}
	return shiftResponse, nil
}

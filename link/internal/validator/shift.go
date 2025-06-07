package validator

import (
	"encoding/json"
	"fmt"
	"log/slog"

	entities "github.com/GHFluding/ShiftManager/SMgrpc/pkg/gen"
	logger "github.com/GHFluding/ShiftManager/link/internal/utils"
)

type shiftDefault struct {
	Machineid     int64
	ShiftMasterID int64
}

func (s shiftDefault) ToGRPCCreateParams() *entities.CreateShiftParams {
	return &entities.CreateShiftParams{
		MachineId:   s.Machineid,
		ShiftMaster: s.ShiftMasterID,
	}
}

func Shift(data []byte, log *slog.Logger) (shiftDefault, error) {
	shift, err := marshalShift(data, log)
	if err != nil {
		return shiftDefault{}, err
	}
	return shift, err
}

func marshalShift(data []byte, log *slog.Logger) (shiftDefault, error) {
	var shift shiftDefault
	if err := json.Unmarshal(data, &shift); err != nil {
		log.Error("JSON unmarshal error", logger.ErrToAttr(err))
		return shift, fmt.Errorf("invalid request format: %w", err)
	}

	log.Info("Parsed shift data",
		slog.Int64("machineid", shift.Machineid),
		slog.Int64("shiftmaster", shift.ShiftMasterID))
	return shift, nil

}

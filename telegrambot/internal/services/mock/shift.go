package service_mock

import (
	"context"
	"telegramSM/internal/telegramapi/commands"
)

type ShiftServiceMock struct {
	shiftList []commands.Shift
}

func (ss ShiftServiceMock) SaveShift(ctx context.Context, shift *commands.Shift) error {
	ss.shiftList = append(ss.shiftList, *shift)
	return nil
}
func (ss ShiftServiceMock) ListShifts(ctx context.Context) ([]commands.Shift, error) {
	return ss.shiftList, nil
}

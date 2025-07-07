package service_mock

import (
	"context"
	"telegramSM/internal/telegramapi/commands"
)

type TaskServiceMock struct {
	taskList []commands.Task
}

func (ts TaskServiceMock) GetTaskToday(ctx context.Context, telegramID int) (*commands.Task, error) {
	return &ts.taskList[0], nil
}
func (ts TaskServiceMock) SaveTask(ctx context.Context, task *commands.Task) error {
	ts.taskList = append(ts.taskList, *task)
	return nil
}

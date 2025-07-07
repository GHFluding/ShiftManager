package service_mock

import (
	"context"
	"fmt"
	"telegramSM/internal/telegramapi/commands"
)

type UserServiceMock struct {
	userList []commands.User
}

func (us UserServiceMock) GetUser(ctx context.Context, telegramID int) (*commands.User, error) {
	for _, user := range us.userList {
		if user.TelegramID == telegramID {
			return &user, nil
		}
	}
	err := fmt.Errorf("undefined user")
	return nil, err
}
func (us UserServiceMock) SaveUser(ctx context.Context, user *commands.User) error {
	us.userList = append(us.userList, *user)
	return nil
}

type MasterServiceMock struct {
	masterList []commands.MasterIcon
}

func (ms MasterServiceMock) ListMasters(ctx context.Context) ([]commands.MasterIcon, error) {
	return ms.masterList, nil
}

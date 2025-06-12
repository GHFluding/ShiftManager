package models

import "context"

type User struct {
	BitrixId   *int64
	TelegramId string
	Name       string
	Role       string
}

type UserSaver interface {
	SaveUser(
		ctx context.Context,
		bitrixId *int64,
		telegramId string,
		name string,
		role string,
	) (
		id int64,
		err error,
	)
}
type UserProvider interface {
	GetUser(ctx context.Context, id int64) (User, error)
}

type UserDB struct {
	Saver    UserSaver
	Provider UserProvider
}

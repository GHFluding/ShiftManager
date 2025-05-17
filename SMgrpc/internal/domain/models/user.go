package models

type User struct {
	BitrixId   *int64
	TelegramId string
	Name       string
	Role       string
}

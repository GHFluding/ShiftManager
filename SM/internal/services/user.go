package services

import (
	"context"
	"sm/internal/database/postgres"
	"sm/internal/utils/logger"
)

type UserListToTransfer struct {
	Valid       bool
	UserListDTO []UserDTO
}

func DeleteUser(sp *ServicesParams, userid int64) bool {
	err := sp.db.DeleteUser(context.Background(), userid)
	if err != nil {
		sp.log.Info("Failed to delete user from db")
		return false
	}
	return true
}

func UsersList(sp *ServicesParams) UserListToTransfer {
	users, err := sp.db.UsersList(context.Background())
	if err != nil {
		sp.log.Info("Failed to retrieve users from db", logger.ErrToAttr(err))
		return UserListToTransfer{Valid: false}
	}
	usersDTO, err := convertListToTransport[postgres.User, UserDTO](users)
	if err != nil {
		sp.log.Info("Failed to convert users from db", logger.ErrToAttr(err))
		return UserListToTransfer{Valid: false}
	}
	return UserListToTransfer{Valid: true, UserListDTO: usersDTO}
}

func UsersListByRole(sp *ServicesParams, role postgres.Userrole) UserListToTransfer {
	users, err := sp.db.UsersListByRole(context.Background(), role)
	if err != nil {
		sp.log.Info("Failed to retrieve users from db", logger.ErrToAttr(err))
		return UserListToTransfer{Valid: false}
	}
	usersDTO, err := convertListToTransport[postgres.User, UserDTO](users)
	if err != nil {
		sp.log.Info("Failed to convert users from db", logger.ErrToAttr(err))
		return UserListToTransfer{Valid: false}
	}
	return UserListToTransfer{Valid: true, UserListDTO: usersDTO}
}

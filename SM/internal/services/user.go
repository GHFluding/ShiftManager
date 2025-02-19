package services

import (
	"context"
	"sm/internal/database/postgres"
	"sm/internal/utils/logger"
)

func DeleteUser(sp *ServicesParams, userid int64) bool {
	err := sp.db.DeleteUser(context.Background(), userid)
	if err != nil {
		sp.log.Info("Failed to delete user from db")
		return false
	}
	return true
}

func UsersList(sp *ServicesParams) ([]User, error) {
	var UsersToOUT []User
	users, err := sp.db.UsersList(context.Background())
	if err != nil {
		sp.log.Info("Failed to retrieve users from db", logger.ErrToAttr(err))
		return UsersToOUT, err
	}
	for _, i := range users {
		UsersToOUT = append(UsersToOUT, convertUser(i))
	}
	return UsersToOUT, nil
}

func UsersListByRole(sp *ServicesParams, role postgres.Userrole) ([]User, error) {
	var UsersToOUT []User
	users, err := sp.db.UsersListByRole(context.Background(), role)
	if err != nil {
		sp.log.Info("Failed to retrieve users from db", logger.ErrToAttr(err))
		return UsersToOUT, err
	}
	for _, i := range users {
		UsersToOUT = append(UsersToOUT, convertUser(i))
	}
	return UsersToOUT, nil
}

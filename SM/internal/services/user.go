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
	var users []User
	usersDB, err := sp.db.UsersList(context.Background())
	if err != nil {
		sp.log.Info("Failed to retrieve users from db", logger.ErrToAttr(err))
		return users, err
	}
	for _, i := range usersDB {
		users = append(users, convertUser(i))
	}
	return users, nil
}

func UsersListByRole(sp *ServicesParams, role postgres.Userrole) ([]User, error) {
	var users []User
	usersDB, err := sp.db.UsersListByRole(context.Background(), role)
	if err != nil {
		sp.log.Info("Failed to retrieve users from db", logger.ErrToAttr(err))
		return users, err
	}
	for _, i := range usersDB {
		users = append(users, convertUser(i))
	}
	return users, nil
}

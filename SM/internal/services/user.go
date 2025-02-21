package services

import (
	"context"
	"errors"
	"sm/internal/database/postgres"
	"sm/internal/utils/logger"
)

func DeleteUser(sp *ServicesParams, userid int64) error {
	err := sp.db.DeleteUser(context.Background(), userid)
	if err != nil {
		sp.log.Info("Failed to delete user from db")
		return err
	}
	return nil
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

func UsersListByRole(sp *ServicesParams, reqRole string) ([]User, error) {
	var users []User
	role, ok := detectUserRole(reqRole)
	if !ok {
		return users, errors.New("invalid role")
	}
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

func CreateUser(sp *ServicesParams, req User) error {
	userParams, err := convertCreateUserParams(req)
	if err != nil {
		return err
	}
	user, err := sp.db.CreateUser(context.Background(), userParams)
	_ = user
	if err != nil {
		return err
	}
	return nil
}

func convertCreateUserParams(req User) (postgres.CreateUserParams, error) {
	var userParams postgres.CreateUserParams
	userParams.ID = req.ID
	userParams.Bitrixid = req.Bitrixid
	userParams.Name = req.Name
	var ok bool
	userParams.Role, ok = detectUserRole(req.Role)
	if !ok {
		return userParams, errors.New("Invalid role")
	}
	return userParams, nil
}

func detectUserRole(sRole string) (postgres.Userrole, bool) {
	switch sRole {
	case "engineer":
		return postgres.UserroleEngineer, true
	case "worker":
		return postgres.UserroleWorker, true
	case "master":
		return postgres.UserroleMaster, true
	case "manager":
		return postgres.UserroleManager, true
	case "admin":
		return postgres.UserroleAdmin, true
	default:
		return postgres.Userrole(""), false
	}
}

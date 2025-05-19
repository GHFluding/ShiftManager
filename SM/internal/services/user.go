package services

import (
	"context"
	"errors"

	"github.com/GHFluding/ShiftManager/SM/internal/database/postgres"
	"github.com/GHFluding/ShiftManager/SM/internal/utils/logger"
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
		users = append(users, convertUserDB(i))
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
		users = append(users, convertUserDB(i))
	}
	return users, nil
}

func CreateUser(sp *ServicesParams, req User) (User, error) {
	userParams, err := convertCreateUserParams(req)
	if err != nil {
		sp.log.Info("Failed to convert params for creating user: ", logger.ErrToAttr(err))
		return User{}, err
	}
	userDB, err := sp.db.CreateUser(context.Background(), userParams)
	if err != nil {
		sp.log.Info("Failed to create user: ", logger.ErrToAttr(err))
		return User{}, err
	}
	user := convertUserDB(userDB)
	return user, nil
}

func convertCreateUserParams(req User) (postgres.CreateUserParams, error) {
	role, ok := detectUserRole(req.Role)
	if !ok {
		return postgres.CreateUserParams{}, errors.New("Invalid role")
	}
	return postgres.CreateUserParams{
		ID:       req.ID,
		Bitrixid: req.Bitrixid,
		Name:     req.Name,
		Role:     role,
	}, nil
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

func CheckUserRole(sp *ServicesParams, userId int64) (string, error) {
	user, err := sp.db.CheckUserRole(context.Background(), userId)
	if err != nil {
		return "", err
	}
	return string(user.Role), err
}

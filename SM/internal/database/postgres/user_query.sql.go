// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user_query.sql

package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const changeUserRole = `-- name: ChangeUserRole :exec
UPDATE Users
SET 
    role = $1
WHERE id = $2
`

type ChangeUserRoleParams struct {
	Role Userrole
	ID   int64
}

func (q *Queries) ChangeUserRole(ctx context.Context, arg ChangeUserRoleParams) error {
	_, err := q.db.Exec(ctx, changeUserRole, arg.Role, arg.ID)
	return err
}

const createUser = `-- name: CreateUser :one
INSERT INTO Users(
    id, bitrixid,telegramid, name, role 
) VALUES (
    $1, $2,$3, $4, $5
)
RETURNING id, bitrixid, telegramid, name, role
`

type CreateUserParams struct {
	ID         int64
	Bitrixid   pgtype.Int8
	Telegramid string
	Name       string
	Role       Userrole
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.ID,
		arg.Bitrixid,
		arg.Telegramid,
		arg.Name,
		arg.Role,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Bitrixid,
		&i.Telegramid,
		&i.Name,
		&i.Role,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM Users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, bitrixid, telegramid, name, role FROM Users
WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Bitrixid,
		&i.Telegramid,
		&i.Name,
		&i.Role,
	)
	return i, err
}

const usersList = `-- name: UsersList :many
Select id, bitrixid, telegramid, name, role FROM Users
ORDER BY id
`

func (q *Queries) UsersList(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, usersList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Bitrixid,
			&i.Telegramid,
			&i.Name,
			&i.Role,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const usersListByRole = `-- name: UsersListByRole :many
Select id, bitrixid, telegramid, name, role FROM Users
WHERE role = $1
ORDER BY id
`

func (q *Queries) UsersListByRole(ctx context.Context, role Userrole) ([]User, error) {
	rows, err := q.db.Query(ctx, usersListByRole, role)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Bitrixid,
			&i.Telegramid,
			&i.Name,
			&i.Role,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

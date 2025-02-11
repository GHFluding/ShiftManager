// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const activeShiftList = `-- name: ActiveShiftList :many
Select id, machineid, shift_master, createdat, isactive, deactivatedat FROM Shifts
WHERE isActive IS TRUE
ORDER BY id
`

func (q *Queries) ActiveShiftList(ctx context.Context) ([]Shift, error) {
	rows, err := q.db.Query(ctx, activeShiftList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Shift
	for rows.Next() {
		var i Shift
		if err := rows.Scan(
			&i.ID,
			&i.Machineid,
			&i.ShiftMaster,
			&i.Createdat,
			&i.Isactive,
			&i.Deactivatedat,
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

const addShiftTask = `-- name: AddShiftTask :one
INSERT INTO Shift_tasks(
    shiftId, taskId
) VALUES (
    $1, $2
)
RETURNING shiftid, taskid
`

type AddShiftTaskParams struct {
	Shiftid int64
	Taskid  int64
}

func (q *Queries) AddShiftTask(ctx context.Context, arg AddShiftTaskParams) (ShiftTask, error) {
	row := q.db.QueryRow(ctx, addShiftTask, arg.Shiftid, arg.Taskid)
	var i ShiftTask
	err := row.Scan(&i.Shiftid, &i.Taskid)
	return i, err
}

const addShiftWorker = `-- name: AddShiftWorker :one
INSERT INTO Shift_workers(
    shiftId, userId
) VALUES (
    $1, $2
)
RETURNING shiftid, userid
`

type AddShiftWorkerParams struct {
	Shiftid int64
	Userid  int64
}

func (q *Queries) AddShiftWorker(ctx context.Context, arg AddShiftWorkerParams) (ShiftWorker, error) {
	row := q.db.QueryRow(ctx, addShiftWorker, arg.Shiftid, arg.Userid)
	var i ShiftWorker
	err := row.Scan(&i.Shiftid, &i.Userid)
	return i, err
}

const changeMachineActivity = `-- name: ChangeMachineActivity :exec
UPDATE Machine
SET 
    isActive = FALSE
WHERE id = $1
`

func (q *Queries) ChangeMachineActivity(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, changeMachineActivity, id)
	return err
}

const changeUserRole = `-- name: ChangeUserRole :exec
UPDATE Users
SET 
    role = $2
WHERE id = $1
`

type ChangeUserRoleParams struct {
	ID   int64
	Role Userrole
}

func (q *Queries) ChangeUserRole(ctx context.Context, arg ChangeUserRoleParams) error {
	_, err := q.db.Exec(ctx, changeUserRole, arg.ID, arg.Role)
	return err
}

const createMachine = `-- name: CreateMachine :one
INSERT INTO Machine(
    id, name, isRepairRequired, isActive 
) VALUES (
    $1, $2, $3, $4
)
RETURNING id, name, isrepairrequired, isactive
`

type CreateMachineParams struct {
	ID               int64
	Name             string
	Isrepairrequired pgtype.Bool
	Isactive         pgtype.Bool
}

func (q *Queries) CreateMachine(ctx context.Context, arg CreateMachineParams) (Machine, error) {
	row := q.db.QueryRow(ctx, createMachine,
		arg.ID,
		arg.Name,
		arg.Isrepairrequired,
		arg.Isactive,
	)
	var i Machine
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Isrepairrequired,
		&i.Isactive,
	)
	return i, err
}

const createShift = `-- name: CreateShift :one
INSERT INTO Shifts(
    id, machineId, shift_master, createdAt, isActive, deactivatedAt
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING id, machineid, shift_master, createdat, isactive, deactivatedat
`

type CreateShiftParams struct {
	ID            int64
	Machineid     int64
	ShiftMaster   int64
	Createdat     pgtype.Date
	Isactive      pgtype.Bool
	Deactivatedat pgtype.Date
}

func (q *Queries) CreateShift(ctx context.Context, arg CreateShiftParams) (Shift, error) {
	row := q.db.QueryRow(ctx, createShift,
		arg.ID,
		arg.Machineid,
		arg.ShiftMaster,
		arg.Createdat,
		arg.Isactive,
		arg.Deactivatedat,
	)
	var i Shift
	err := row.Scan(
		&i.ID,
		&i.Machineid,
		&i.ShiftMaster,
		&i.Createdat,
		&i.Isactive,
		&i.Deactivatedat,
	)
	return i, err
}

const createTask = `-- name: CreateTask :one
INSERT INTO Tasks(
    id, machineId, shiftId, frequency, taskPriority, description, createdBy, createdAt
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING id, machineid, shiftid, frequency, taskpriority, description, createdby, createdat, verifiedby, verifiedat, completedby, completedat, status, comment, movedinprogressby, movedinprogressat
`

type CreateTaskParams struct {
	ID           int64
	Machineid    int64
	Shiftid      pgtype.Int8
	Frequency    Taskfrequency
	Taskpriority Taskpriority
	Description  string
	Createdby    int64
	Createdat    pgtype.Date
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error) {
	row := q.db.QueryRow(ctx, createTask,
		arg.ID,
		arg.Machineid,
		arg.Shiftid,
		arg.Frequency,
		arg.Taskpriority,
		arg.Description,
		arg.Createdby,
		arg.Createdat,
	)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Machineid,
		&i.Shiftid,
		&i.Frequency,
		&i.Taskpriority,
		&i.Description,
		&i.Createdby,
		&i.Createdat,
		&i.Verifiedby,
		&i.Verifiedat,
		&i.Completedby,
		&i.Completedat,
		&i.Status,
		&i.Comment,
		&i.Movedinprogressby,
		&i.Movedinprogressat,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO Users(
    id, bitrixid, name, role 
) VALUES (
    $1, $2, $3, $4
)
RETURNING id, bitrixid, name, role
`

type CreateUserParams struct {
	ID       int64
	Bitrixid int64
	Name     string
	Role     Userrole
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.ID,
		arg.Bitrixid,
		arg.Name,
		arg.Role,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Bitrixid,
		&i.Name,
		&i.Role,
	)
	return i, err
}

const deleteMachine = `-- name: DeleteMachine :exec
DELETE FROM Machine
WHERE id = $1
`

func (q *Queries) DeleteMachine(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteMachine, id)
	return err
}

const deleteShift = `-- name: DeleteShift :exec
DELETE FROM Shifts
WHERE id = $1
`

func (q *Queries) DeleteShift(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteShift, id)
	return err
}

const deleteShiftTask = `-- name: DeleteShiftTask :exec
DELETE FROM Shift_tasks
WHERE shiftId = $1 AND taskId = $2
`

type DeleteShiftTaskParams struct {
	Shiftid int64
	Taskid  int64
}

func (q *Queries) DeleteShiftTask(ctx context.Context, arg DeleteShiftTaskParams) error {
	_, err := q.db.Exec(ctx, deleteShiftTask, arg.Shiftid, arg.Taskid)
	return err
}

const deleteShiftWorker = `-- name: DeleteShiftWorker :exec
DELETE FROM Shift_workers
WHERE shiftId = $1 AND userId = $2
`

type DeleteShiftWorkerParams struct {
	Shiftid int64
	Userid  int64
}

func (q *Queries) DeleteShiftWorker(ctx context.Context, arg DeleteShiftWorkerParams) error {
	_, err := q.db.Exec(ctx, deleteShiftWorker, arg.Shiftid, arg.Userid)
	return err
}

const deleteTask = `-- name: DeleteTask :exec
DELETE FROM Tasks
WHERE id = $1
`

func (q *Queries) DeleteTask(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteTask, id)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM Users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const machineNeedRepair = `-- name: MachineNeedRepair :exec
UPDATE Machine
SET 
    isRepairRequired = TRUE
WHERE id = $1
`

func (q *Queries) MachineNeedRepair(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, machineNeedRepair, id)
	return err
}

const setTaskStatusCompleted = `-- name: SetTaskStatusCompleted :exec
UPDATE Tasks
SET 
    status = 'completed',
    completedAt = CURRENT_DATE,
    movedInProgressBy = $2
WHERE id = $1
`

type SetTaskStatusCompletedParams struct {
	ID                int64
	Movedinprogressby pgtype.Int8
}

func (q *Queries) SetTaskStatusCompleted(ctx context.Context, arg SetTaskStatusCompletedParams) error {
	_, err := q.db.Exec(ctx, setTaskStatusCompleted, arg.ID, arg.Movedinprogressby)
	return err
}

const setTaskStatusFailed = `-- name: SetTaskStatusFailed :exec
UPDATE Tasks
SET 
    status = 'failed',
    comment = $2
WHERE id = $1
`

type SetTaskStatusFailedParams struct {
	ID      int64
	Comment pgtype.Text
}

func (q *Queries) SetTaskStatusFailed(ctx context.Context, arg SetTaskStatusFailedParams) error {
	_, err := q.db.Exec(ctx, setTaskStatusFailed, arg.ID, arg.Comment)
	return err
}

const setTaskStatusInProgress = `-- name: SetTaskStatusInProgress :exec
UPDATE Tasks
SET 
    status = 'inProgress',
    movedInProgressAt = CURRENT_DATE,
    movedInProgressBy = $2
WHERE id = $1
`

type SetTaskStatusInProgressParams struct {
	ID                int64
	Movedinprogressby pgtype.Int8
}

func (q *Queries) SetTaskStatusInProgress(ctx context.Context, arg SetTaskStatusInProgressParams) error {
	_, err := q.db.Exec(ctx, setTaskStatusInProgress, arg.ID, arg.Movedinprogressby)
	return err
}

const setTaskStatusVerified = `-- name: SetTaskStatusVerified :exec
UPDATE Tasks
SET 
    status = 'verified',
    verifiedAt = CURRENT_DATE,
    movedInProgressBy = $2
WHERE id = $1
`

type SetTaskStatusVerifiedParams struct {
	ID                int64
	Movedinprogressby pgtype.Int8
}

func (q *Queries) SetTaskStatusVerified(ctx context.Context, arg SetTaskStatusVerifiedParams) error {
	_, err := q.db.Exec(ctx, setTaskStatusVerified, arg.ID, arg.Movedinprogressby)
	return err
}

const shiftList = `-- name: ShiftList :many
Select id, machineid, shift_master, createdat, isactive, deactivatedat FROM Shifts
ORDER BY id
`

func (q *Queries) ShiftList(ctx context.Context) ([]Shift, error) {
	rows, err := q.db.Query(ctx, shiftList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Shift
	for rows.Next() {
		var i Shift
		if err := rows.Scan(
			&i.ID,
			&i.Machineid,
			&i.ShiftMaster,
			&i.Createdat,
			&i.Isactive,
			&i.Deactivatedat,
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

const shiftTasksList = `-- name: ShiftTasksList :many
Select shiftid, taskid FROM Shift_tasks
WHERE shiftId = $1
ORDER BY taskId
`

func (q *Queries) ShiftTasksList(ctx context.Context, shiftid int64) ([]ShiftTask, error) {
	rows, err := q.db.Query(ctx, shiftTasksList, shiftid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ShiftTask
	for rows.Next() {
		var i ShiftTask
		if err := rows.Scan(&i.Shiftid, &i.Taskid); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const shiftWorkersList = `-- name: ShiftWorkersList :many
Select shiftid, userid FROM Shift_workers
WHERE shiftId = $1  
ORDER BY userId
`

func (q *Queries) ShiftWorkersList(ctx context.Context, shiftid int64) ([]ShiftWorker, error) {
	rows, err := q.db.Query(ctx, shiftWorkersList, shiftid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ShiftWorker
	for rows.Next() {
		var i ShiftWorker
		if err := rows.Scan(&i.Shiftid, &i.Userid); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const usersList = `-- name: UsersList :many
Select id, bitrixid, name, role FROM Users
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
Select id, bitrixid, name, role FROM Users
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

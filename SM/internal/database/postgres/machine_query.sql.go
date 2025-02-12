// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: machine_query.sql

package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const changeMachineActivity = `-- name: ChangeMachineActivity :exec
UPDATE Machine
SET 
    isActive = $1
WHERE id = $2
`

type ChangeMachineActivityParams struct {
	Isactive pgtype.Bool
	ID       int64
}

func (q *Queries) ChangeMachineActivity(ctx context.Context, arg ChangeMachineActivityParams) error {
	_, err := q.db.Exec(ctx, changeMachineActivity, arg.Isactive, arg.ID)
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

const deleteMachine = `-- name: DeleteMachine :exec
DELETE FROM Machine
WHERE id = $1
`

func (q *Queries) DeleteMachine(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteMachine, id)
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

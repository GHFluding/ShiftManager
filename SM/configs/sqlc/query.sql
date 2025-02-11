-- name: CreateUser :one
INSERT INTO Users(
    id, bitrixid, name, role 
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;
-- name: DeleteUser :exec
DELETE FROM Users
WHERE id = $1;
-- name: ChangeUserRole :exec
UPDATE Users
SET 
    role = $2
WHERE id = $1;
-- name: UsersList :many
Select * FROM Users
ORDER BY id;
-- name: UsersListByRole :many
Select * FROM Users
WHERE role = $1
ORDER BY id;


-- name: CreateShift :one
INSERT INTO Shifts(
    id, machineId, shift_master, createdAt, isActive, deactivatedAt
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;
-- name: ShiftList :many
Select * FROM Shifts
ORDER BY id;
-- name: ActiveShiftList :many
Select * FROM Shifts
WHERE isActive IS TRUE
ORDER BY id;
-- name: DeleteShift :exec
DELETE FROM Shifts
WHERE id = $1;

-- name: AddShiftWorker :one
INSERT INTO Shift_workers(
    shiftId, userId
) VALUES (
    $1, $2
)
RETURNING *;
-- name: ShiftWorkersList :many
Select * FROM Shift_workers
WHERE shiftId = $1  
ORDER BY userId;
-- name: DeleteShiftWorker :exec
DELETE FROM Shift_workers
WHERE shiftId = $1 AND userId = $2;

-- name: AddShiftTask :one
INSERT INTO Shift_tasks(
    shiftId, taskId
) VALUES (
    $1, $2
)
RETURNING *;
-- name: ShiftTasksList :many
Select * FROM Shift_tasks
WHERE shiftId = $1
ORDER BY taskId;
-- name: DeleteShiftTask :exec
DELETE FROM Shift_tasks
WHERE shiftId = $1 AND taskId = $2;


-- name: CreateTask :one
INSERT INTO Tasks(
    id, machineId, shiftId, frequency, taskPriority, description, createdBy, createdAt
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;
-- name: DeleteTask :exec
DELETE FROM Tasks
WHERE id = $1;
-- name: SetTaskStatusInProgress :exec
UPDATE Tasks
SET 
    status = 'inProgress',
    movedInProgressAt = CURRENT_DATE,
    movedInProgressBy = $2
WHERE id = $1;
-- name: SetTaskStatusFailed :exec
UPDATE Tasks
SET 
    status = 'failed',
    comment = $2
WHERE id = $1;
-- name: SetTaskStatusCompleted :exec
UPDATE Tasks
SET 
    status = 'completed',
    completedAt = CURRENT_DATE,
    movedInProgressBy = $2
WHERE id = $1;
-- name: SetTaskStatusVerified :exec
UPDATE Tasks
SET 
    status = 'verified',
    verifiedAt = CURRENT_DATE,
    movedInProgressBy = $2
WHERE id = $1;

-- name: CreateMachine :one
INSERT INTO Machine(
    id, name, isRepairRequired, isActive 
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;
-- name: ChangeMachineActivity :exec
UPDATE Machine
SET 
    isActive = FALSE
WHERE id = $1;
-- name: MachineNeedRepair :exec
UPDATE Machine
SET 
    isRepairRequired = TRUE
WHERE id = $1;
-- name: DeleteMachine :exec
DELETE FROM Machine
WHERE id = $1;


--name: CreateUser: one
INSERT INTO Users(
    id, bitrixid, name, role 
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;
-- name: DeleteUser: exec
DELETE FROM Users
WHERE id = $1;
--name: ChangeUserRole: exec
UPDATE Users
SET 
    role = $2,
WHERE id = $1;
--name: UsersList: many
Select * FROM Users
ORDER BY id;
--name: UsersListByRole: many
Select * FROM Users
WHERE role = $1
ORDER BY id;


--name: CreateShift: one
INSERT INTO Shifts(
    id, machineId, shift_master, createdAt, isActive, deactivatedAt
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;
--name: ShiftList: many
Select * FROM Shifts
ORDER BY id;
--name: ActiveShiftList: many
Select * FROM Shifts
WHERE isActive IS TRUE
ORDER BY id;
-- name: DeleteShift :exec
DELETE FROM Shifts
WHERE id = $1;

--name: AddShiftWorker: one
INSERT INTO Shift_workers(
    shiftId, userId
) VALUES (
    $1, $2
)
RETURNING *;
--name: ShiftWorkersList: many
Select * FROM Shift_workers
WHERE shiftId = $1  
ORDER BY id;
-- name: DeleteShiftWorker :exec
DELETE FROM Shift_workers
WHERE shiftId = $1 AND userId = $2;

--name: AddShiftTask: one
INSERT INTO Shift_workers(
    shiftId, taskId
) VALUES (
    $1, $2
)
RETURNING *;
---name: ShiftTasksList: many
Select * FROM Shift_tasks
WHERE shiftId = $2
ORDER BY id;
-- name: DeleteShiftTask :exec
DELETE FROM Shift_tasks
WHERE shiftId = $1 AND taskId = $2;


--name: CreateTask: one
INSERT INTO Tasks(
    id, machineId, shiftId, taskPriority, description, createdBy, createdAt
) VALUES (
    $1, $2, $3, $4, $5, $6, $7,
)
RETURNING *;
-- name: DeleteTask :exec
DELETE FROM Tasks
WHERE id = $1;
--name: SetTaskStatusInProgress: one
UPDATE Tasks
SET 
    status = 'inProgress',
    NEW.movedInProgressAt = CURRENT_DATE
    movedInProgressBy = $2
WHERE id = $1;
--name: SetTaskStatusFailed: one
UPDATE Tasks
SET 
    status = 'failed',
    comment = $2
WHERE id = $1;
--name: SetTaskStatusCompleted: one
UPDATE Tasks
SET 
    status = 'completed',
    NEW.completedAt = CURRENT_DATE
    movedInProgressBy = $2
WHERE id = $1;
--name: SetTaskStatusVerified: one
UPDATE Tasks
SET 
    status = 'verified',
    NEW.verifiedAt = CURRENT_DATE
    movedInProgressBy = $2
WHERE id = $1;

--name: CreateMachine: one
INSERT INTO Machine(
    id, name, isRepairRequired, isActive 
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;
--name: ChangeMachineActivity: exec
UPDATE Machine
SET 
    isActive = FALSE,
WHERE id = $1;
--name: MachineNeedRepair: exec
UPDATE Machine
SET 
    isRepairRequired = TRUE,
WHERE id = $1;
-- name: DeleteMachine :exec
DELETE FROM Machine
WHERE id = $1;


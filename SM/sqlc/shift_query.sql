-- name: CreateShift :one
INSERT INTO Shifts(
    id, machineId, shift_master, createdAt, isActive
) VALUES (
    @id, @machineId, @shift_master, CURRENT_DATE, TRUE
)
RETURNING *;
-- name: ChangeShiftActivity :exec
UPDATE Shifts
SET 
    isActive = @isActive
WHERE id = @id;

-- name: ShiftList :many
Select * FROM Shifts
ORDER BY id;
-- name: ActiveShiftList :many
Select * FROM Shifts
WHERE isActive IS TRUE
ORDER BY id;
-- name: DeleteShift :exec
DELETE FROM Shifts
WHERE id = @id;

-- name: AddShiftWorker :one
INSERT INTO Shift_workers(
    shiftId, userId
) VALUES (
    @shiftId, @userId
)
RETURNING *;
-- name: ShiftWorkersList :many
Select * FROM Shift_workers
WHERE shiftId = @shiftId 
ORDER BY userId;
-- name: DeleteShiftWorker :exec
DELETE FROM Shift_workers
WHERE shiftId = @shiftId AND userId = @userId;

-- name: AddShiftTask :one
INSERT INTO Shift_tasks(
    shiftId, taskId
) VALUES (
    @shiftId, @taskId
)
RETURNING *;
-- name: ShiftTasksList :many
Select * FROM Shift_tasks
WHERE shiftId = @shiftId
ORDER BY taskId;
-- name: DeleteShiftTask :exec
DELETE FROM Shift_tasks
WHERE shiftId = @shiftId AND taskId = @taskId;
-- name: GetShift :one
SELECT * FROM Shifts
WHERE id = @id;
-- name: CreateTask :one
INSERT INTO Tasks(
    id, machineId, shiftId, frequency, taskPriority, description, createdBy, createdAt
) VALUES (
    @id, @machineId, @shiftId, @frequency, @taskPriority, @description, @createdBy, CURRENT_DATE
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
    movedInProgressBy = @movedInProgressBy
WHERE id = @id;
-- name: SetTaskStatusFailed :exec
UPDATE Tasks
SET 
    status = 'failed',
    comment = @comment
WHERE id = @id;
-- name: SetTaskStatusCompleted :exec
UPDATE Tasks
SET 
    status = 'completed',
    completedAt = CURRENT_DATE,
    completedBy = @completedBy
WHERE id = @id;
-- name: SetTaskStatusVerified :exec
UPDATE Tasks
SET 
    status = 'verified',
    verifiedAt = CURRENT_DATE,
    verifiedBy = @verifiedBy
WHERE id = @id;

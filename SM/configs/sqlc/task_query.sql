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

-- name: UpdateTaskStatus :exec
UPDATE Tasks
SET 
    status = @status,
    comment = COALESCE(@comment, comment),
    verifiedBy = CASE WHEN @status = 'verified' THEN @userId ELSE verifiedBy END,
    verifiedAt = CASE WHEN @status = 'verified' THEN NOW() ELSE verifiedAt END,
    completedBy = CASE WHEN @status = 'completed' THEN @userId ELSE completedBy END,
    completedAt = CASE WHEN @status = 'completed' THEN NOW() ELSE completedAt END,
    movedInProgressBy = CASE WHEN @status = 'inProgress' THEN @userId ELSE movedInProgressBy END,
    movedInProgressAt = CASE WHEN @status = 'inProgress' THEN NOW() ELSE movedInProgressAt END,
    failedBy = CASE WHEN @status = 'failed' THEN @userId ELSE failedBy END,
    failedAt = CASE WHEN @status = 'failed' THEN NOW() ELSE failedAt END
WHERE id = @taskId;

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

-- name: CreateMachine :one
INSERT INTO Machine(
    id, name, isRepairRequired, isActive 
) VALUES (
    @id, @name, @isRepairRequired, @isActive
)
RETURNING *;
-- name: ChangeMachineActivity :exec
UPDATE Machine
SET 
    isActive = @isActive
WHERE id = @id;
-- name: MachineNeedRepair :exec
UPDATE Machine
SET 
    isRepairRequired = TRUE
WHERE id = @id;
-- name: DeleteMachine :exec
DELETE FROM Machine
WHERE id = @id;
-- name: GetMachine :one
SELECT * FROM Machine
WHERE id = $1;

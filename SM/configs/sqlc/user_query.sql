-- name: CreateUser :one
INSERT INTO Users(
    id, bitrixid, name, role 
) VALUES (
    @id, @bitrixid, @name, @role
)
RETURNING *;
-- name: DeleteUser :exec
DELETE FROM Users
WHERE id = @id;
-- name: ChangeUserRole :exec
UPDATE Users
SET 
    role = @role
WHERE id = @id;
-- name: UsersList :many
Select * FROM Users
ORDER BY id;
-- name: UsersListByRole :many
Select * FROM Users
WHERE role = @role
ORDER BY id;
-- name: CheckUserRole :one
SELECT * FROM Users
WHERE id = @id;

-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;
-- name: InsertUser :exec
INSERT INTO users (id, name, password, email, phone, user_type)
VALUES ($1, $2, $3, $4, $5, $6);
-- name: UpdateUser :one
UPDATE Users
SET name = COALESCE($2, name),
    password = COALESCE($3, password),
    phone = COALESCE($4, phone)
WHERE ID = $1
RETURNING *;
-- name: DeleteUser :one
DELETE from Users
where id = $1
RETURNING *;
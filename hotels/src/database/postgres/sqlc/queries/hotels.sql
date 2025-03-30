-- name: GetHotel :one
SELECT *
FROM hotels
WHERE id = $1
LIMIT 1;
-- name: InsertHotel :exec
INSERT INTO hotels (id, name, location, address, manager_id)
VALUES ($1, $2, $3, $4, $5);
-- name: UpdateHotel :one
UPDATE Hotels
SET name = COALESCE($2, name),
    location = COALESCE($3, location),
    address = COALESCE($3, address),
    manager_id = COALESCE($4, manager_id)
WHERE ID = $1
RETURNING *;
-- name: DeleteHotel :one
DELETE from Hotels
where id = $1
RETURNING *;
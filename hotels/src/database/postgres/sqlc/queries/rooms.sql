-- name: GetRoomByID :one
SELECT *
FROM Rooms
WHERE id = $1
LIMIT 1;
-- name: CreateRoom :exec
INSERT INTO Rooms (id, hotel_id, floor, price, images, description)
VALUES ($1, $2, $3, $4, $5, $6);
-- name: UpdateRoom :one
UPDATE Rooms
SET price = COALESCE($2, price),
    description = COALESCE($3, description),
    images = CASE
                WHEN $4 IS NOT NULL THEN ARRAY_CAT(images, $4::TEXT[])
                ELSE images
            END
WHERE ID = $1
RETURNING *;
-- name: DeleteRoomByID :one
DELETE from Rooms
where id = $1
RETURNING *;
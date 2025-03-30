-- name: CreateReview :one
INSERT INTO reviews (id, user_id, booking_id, room_id, rating, review)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id;

-- name: GetReview :one
SELECT * FROM reviews
WHERE id = $1;

-- name: GetReviewsByFilters :many
SELECT * FROM reviews
WHERE
    (user_id = $1 OR $1 IS NULL) AND
    (room_id = $2 OR $2 IS NULL) AND
    (rating > $3 OR $3 IS NULL);

-- name: GetReviewsByUser :many
SELECT * FROM reviews
WHERE user_id = $1;

-- name: UpdateReview :one
UPDATE reviews
SET rating = $2, review = $3
WHERE id = $1
RETURNING id;

-- name: DeleteReview :exec
DELETE FROM reviews
WHERE id = $1;

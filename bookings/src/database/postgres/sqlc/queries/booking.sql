-- name: CreateBooking :one
INSERT INTO bookings (user_id, room_id, hotel_id, booking_time, duration)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;

-- name: GetBookingByID :one
SELECT * FROM bookings WHERE id = $1;

-- name: UpdateBooking :one
UPDATE bookings
SET room_id = COALESCE($2, room_id),
    booking_time = COALESCE($3, booking_time),
    duration = COALESCE($4, duration)
WHERE id = $1
RETURNING id;

-- name: DeleteBooking :one
DELETE FROM bookings WHERE id = $1
RETURNING id;

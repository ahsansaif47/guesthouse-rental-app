package utils

import (
	"github.com/google/uuid"
)

type Review struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	BookingID uuid.UUID `json:"booking_id"`
	RoomID    uuid.UUID `json:"room_id"`
	Rating    int32     `json:"rating"`
	Comment   string    `json:"comment"`
}

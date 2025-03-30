package services

import (
	"bookings/src/database"
	"bookings/src/database/postgres/sqlc/generated"
	"bookings/src/utils"
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func CreateReview(review utils.Review) (string, error) {

	id := uuid.New()

	reviewParams := generated.CreateReviewParams{
		ID:        pgtype.UUID{Bytes: id, Valid: true},
		UserID:    pgtype.UUID{Bytes: review.UserID, Valid: true},
		BookingID: pgtype.UUID{Bytes: review.BookingID, Valid: true},
		RoomID:    pgtype.UUID{Bytes: review.RoomID, Valid: true},
		Rating:    review.Rating,
		Review:    pgtype.Text{String: review.Comment, Valid: true},
	}

	revID, err := database.PgHandler.CreateReview(context.Background(), reviewParams)
	if err != nil {
		log.Println("Failed to create review object")
		return "", err
	}
	return revID.String(), nil
}

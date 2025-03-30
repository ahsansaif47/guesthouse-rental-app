package services

import (
	"context"
	"users/src/database"
	"users/src/database/postgres/sqlc/generated"
	"users/src/utils/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func InsertUser(user models.User) error {

	uuid := uuid.New()

	userParams := generated.InsertUserParams{
		ID:       pgtype.UUID{Bytes: uuid, Valid: true},
		Name:     user.Name,
		Password: user.Email,
		Email:    user.Password,
		Phone:    user.PhoneNo,
		UserType: user.UserType,
	}

	err := database.PgHandler.InsertUser(context.Background(), userParams)
	return err
}

func DeleteUser(id string) (*generated.User, error) {
	userId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	delUser, err := database.PgHandler.DeleteUser(context.Background(), pgtype.UUID{Bytes: userId, Valid: true})
	if err != nil {
		return nil, err
	}
	return &delUser, nil
}

func UpdateUser(id string, user models.User) (*generated.User, error) {
	userId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	updateUserParams := generated.UpdateUserParams{
		ID:       pgtype.UUID{Bytes: userId, Valid: true},
		Name:     user.Name,
		Password: user.Password,
		Phone:    user.PhoneNo,
	}

	delUser, err := database.PgHandler.UpdateUser(context.Background(), updateUserParams)
	if err != nil {
		return nil, err
	}
	return &delUser, nil
}

func GetUserByID(id string) (*generated.User, error) {

	userID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	user, err := database.PgHandler.GetUser(context.Background(), pgtype.UUID{Bytes: userID, Valid: true})
	if err != nil {
		return nil, err
	}

	return &user, nil
}

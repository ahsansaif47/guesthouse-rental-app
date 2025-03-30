package main

import (
	"fmt"
	"os"
	"users/src/database"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("Failed to load env! %s", err.Error()))
	}

	dbUrl := os.Getenv("POSTGRES_DATABASE_URL")

	// Connect to database.
	pgPool, err := database.ConnectPostgres(dbUrl)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database! %s", err.Error()))
	} else {
		database.PgPool = pgPool
		database.PgHandler = database.GetPgHandler()
	}
}

func main() {
	defer database.PgPool.Close()

}

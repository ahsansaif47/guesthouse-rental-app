package database

import (
	"context"
	"fmt"
	"time"
	"users/src/database/postgres/sqlc/generated"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	PgPool    *pgxpool.Pool      = &pgxpool.Pool{}
	PgHandler *generated.Queries = &generated.Queries{}
)

func setupPgxPool(config *pgxpool.Config) (*pgxpool.Pool, error) {

	// Custom pool settings
	config.MaxConns = 10                      // Maximum connections
	config.MinConns = 2                       // Minimum connections
	config.MaxConnLifetime = time.Hour        // Max connection lifetime
	config.MaxConnIdleTime = 30 * time.Minute // Max idle time

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}
	return pool, nil
}

func ConnectPostgres(dbUrl string) (*pgxpool.Pool, error) {
	poolConfig, err := pgxpool.ParseConfig(dbUrl)

	if err != nil {
		return nil, err
	}

	dbPool, err := setupPgxPool(poolConfig)
	if err != nil {
		return nil, err
	}

	// Its better to do this in main!
	// PgPool = dbPool

	return dbPool, nil
}

func GetPgHandler() *generated.Queries {
	return generated.New(PgPool)
}

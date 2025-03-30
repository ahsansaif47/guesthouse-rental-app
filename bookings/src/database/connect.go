package database

import (
	"bookings/src/database/postgres/sqlc/generated"
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	PgPool    *pgxpool.Pool      = &pgxpool.Pool{}
	PgHandler *generated.Queries = &generated.Queries{}
)

func SetupPgxPool(config *pgxpool.Config) (*pgxpool.Pool, error) {
	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnIdleTime = 30 * time.Minute
	config.MaxConnLifetime = time.Hour

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

func ConnectPostgres(dbUrl string) (*pgxpool.Pool, error) {
	poolConfig, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		return nil, err
	}
	pool, err := SetupPgxPool(poolConfig)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

func GetPgHandler() *generated.Queries {
	return generated.New(PgPool)
}

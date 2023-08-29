package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

func NewPostgresPool(dbUrl string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		return nil, errors.Wrap(err, "unable create pgxpool")
	}

	err = pool.Ping(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "unable ping database")
	}

	return pool, nil
}

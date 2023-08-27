package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

func NewPostgresConnection(dbUrl string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		fmt.Println(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("DB connected!")
	}
	return conn
}

func NewPostgresPool(dbUrl string) *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		fmt.Println(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("DB connected!")
	}
	return pool
}

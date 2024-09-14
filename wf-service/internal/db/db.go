package db

import (
	"context"
	"github.com/jackc/pgx/v4"
	"log"
)

func Connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://user:password@localhost:5432/wfdb")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		return nil, err
	}
	return conn, nil
}

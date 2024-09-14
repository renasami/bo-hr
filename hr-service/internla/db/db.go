package db

import (
	"context"
	"github.com/jackc/pgx/v4"
	"log"
	"os"
)

func Connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		return nil, err
	}
	return conn, nil
}

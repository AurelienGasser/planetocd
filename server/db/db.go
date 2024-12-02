package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

const PLANETOCD_DATABASE_URL = "PLANETOCD_DATABASE_URL"

func GetDbConnection() (*pgx.Conn, error) {
	connStr := os.Getenv(PLANETOCD_DATABASE_URL)
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

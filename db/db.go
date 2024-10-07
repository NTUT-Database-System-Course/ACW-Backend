package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	sqlc "github.com/NTUT-Database-System-Course/ACW-Backend/output/sqlcgen"
)

func NewDB() (*pgxpool.Pool, *sqlc.Queries) {
	connStr := "postgres://root:1234@localhost:5432/db"
	// connStr := "postgres://postgres:calcal321@localhost:5432/postgres"
	pool, err := pgxpool.New(context.Background(), connStr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	CreateUserTable(pool)

	q := sqlc.New(pool) // Initialize sqlc with the connection pool

	return pool, q
}

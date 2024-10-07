package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateUserTable(p *pgxpool.Pool) {
	p.Exec(context.Background(),
		`CREATE TABLE "user"(
			id SERIAL 	PRIMARY KEY,
			"username" 	TEXT NOT NULL,
			"email" 	TEXT NOT NULL;`,
	)
	fmt.Println("test")
}

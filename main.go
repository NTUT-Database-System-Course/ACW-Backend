package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/NTUT-Database-System-Course/ACW-Backend/pkg/config"
	"github.com/NTUT-Database-System-Course/ACW-Backend/pkg/router"
)

// @contact.name   API Support
// @contact.url    https://github.com/NTUT-Database-System-Course/ACW-Backend/issues
// @contact.email  chenshiang@onon1101.org
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	config.NewSwaggerInfo()

	e := echo.New()
	// ---------------

	connStr := "postgres://root:1234@localhost:5432/db"
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	// queries := tutorial.New(pool)
	// authors, err := queries.ListAuthors(context.Background())
	// // if err != nil {
	// // 	return err
	// // }
	// log.Println(authors)

	_, err = pool.Exec(context.Background(), `
	        CREATE TABLE authors (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL,
  bio  text
);
	    `)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create table: %v\n", err)
		os.Exit(1)
	}

	username := "new_user"
	_, err = pool.Exec(context.Background(),
		`INSERT INTO "user" (username) VALUES ($1)`, username)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to insert user: %v\n", err)
		os.Exit(1)
	}

	// 查詢所有使用者
	rows, err := pool.Query(context.Background(), `SELECT id, username FROM "user"`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to query users: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()
	fmt.Println("Users in the table:")
	for rows.Next() {
		var id int
		var username string
		err := rows.Scan(&id, &username)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to scan row: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("ID: %d, Username: %s\n", id, username)
	}

	// ---------------

	router.NewRouter(e)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

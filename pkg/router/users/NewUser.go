package users

import (
	"context"
	"fmt"
	"net/http"
	"os"

	sqlc "github.com/NTUT-Database-System-Course/ACW-Backend/output/sqlcgen"
	"github.com/labstack/echo/v4"
)

func NewUser(q *sqlc.Queries) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, err := q.CreateUser(context.Background(), sqlc.CreateUserParams{
			Username: "test02",
			Email:    "test02@test02.com",
		})

		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}
		return c.JSON(http.StatusOK, user)
	}
	// return c.String(http.StatusOK, "test")
}

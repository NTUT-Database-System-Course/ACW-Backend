package users

import (
	"context"
	"fmt"
	"net/http"
	"os"

	sqlc "github.com/NTUT-Database-System-Course/ACW-Backend/output/sqlcgen"
	"github.com/labstack/echo/v4"
)

func GetUser(q *sqlc.Queries) echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := q.GetUsers(context.Background())

		if err != nil {
			fmt.Fprintf(os.Stderr, "Getuser %v\n", err)
			os.Exit(1)
		}

		return c.JSON(http.StatusOK, users)
	}
}

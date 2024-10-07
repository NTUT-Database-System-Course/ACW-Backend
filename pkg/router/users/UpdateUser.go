package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func UpdateUser(c echo.Context) error {
	return c.String(http.StatusOK, "test")
}

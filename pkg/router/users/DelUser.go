package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func DelUser(c echo.Context) error {
	return c.String(http.StatusOK, "test")
}

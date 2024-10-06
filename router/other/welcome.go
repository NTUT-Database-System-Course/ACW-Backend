package other

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Msg struct {
	Msg string `json:"msg"`
}

// @Summary		Get Welcom
// @Description	Print welcome
// @Tags			User
// @Accept			json
// @Produce		json
// @success		200	{object}	Msg
// @Router			/api/hello [get]
func WelcomeMsg(c echo.Context) error {
	return c.JSON(http.StatusOK, Msg{
		Msg: "welcome",
	})
}

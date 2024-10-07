package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/NTUT-Database-System-Course/ACW-Backend/pkg/router/other"
	"github.com/NTUT-Database-System-Course/ACW-Backend/tools"
)

var (
	helloJSON = `{"hello":{"msg":"Welcome to our API!"}}`
)

func TestHelloHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/hello")

	// Assertions
	if assert.NoError(t, other.WelcomeMsg(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		assert.Equal(t, helloJSON, test.WithoutNewLine(rec.Body.String()))
	}
}

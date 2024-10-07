package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/NTUT-Database-System-Course/ACW-Backend/router/other"
)

func NewRouter(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// e.GET("/api", other.WelcomeMsg)

	// 存放 api 所有路徑的索引
	g := e.Group("/api")
	{
		g.GET("/hello", other.WelcomeMsg)
	}

}

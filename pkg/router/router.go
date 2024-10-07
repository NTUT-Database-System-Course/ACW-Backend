package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/NTUT-Database-System-Course/ACW-Backend/pkg/router/other"
	"github.com/NTUT-Database-System-Course/ACW-Backend/pkg/router/users"
)

func NewRouter(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// e.GET("/api", other.WelcomeMsg)

	// 存放 api 所有路徑的索引
	g := e.Group("/api")
	{
		g.GET("/hello", other.WelcomeMsg)

		u := g.Group("/users")
		{
			u.GET("/new", users.NewUser)
			u.GET("/update", users.UpdateUser)
			u.GET("/get", users.GetUser)
			u.GET("/delete", users.DelUser)
		}
	}

}

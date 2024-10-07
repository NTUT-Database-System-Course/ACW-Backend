package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	sqlc "github.com/NTUT-Database-System-Course/ACW-Backend/output/sqlcgen"
	"github.com/NTUT-Database-System-Course/ACW-Backend/pkg/router/other"
	"github.com/NTUT-Database-System-Course/ACW-Backend/pkg/router/users"
)

func NewRouter(e *echo.Echo, q *sqlc.Queries) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// e.GET("/api", other.WelcomeMsg)

	// 存放 api 所有路徑的索引
	g := e.Group("/api")
	{
		g.GET("/hello", other.WelcomeMsg)

		u := g.Group("/users")
		{
			u.GET("/new", users.NewUser(q))
			u.GET("/update", users.UpdateUser)
			u.GET("/get", users.GetUser(q))
			u.GET("/delete", users.DelUser)
		}
	}

}

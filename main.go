package main

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/NTUT-Database-System-Course/ACW-Backend/db"
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

	pool, q := db.NewDB()

	router.NewRouter(e, q)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))

	pool.Close()
}

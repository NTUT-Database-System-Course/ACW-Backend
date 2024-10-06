package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
	_ "backend/docs"
)

// @Summary      獲取用戶資料
// @Description  根據用戶ID獲取用戶資料
// @Tags         用戶
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "用戶ID"
// @Router       /users/{id} [get]
func GetUser(c echo.Context) error {
    
    // 根據 ID 獲取用戶資料
    return c.String(http.StatusOK, "hello, world")
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	e.GET("/api", hello)

	t := e.Group("/api/")
	{
		t.GET("test", GetUser)
	}

	
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello, world")
}
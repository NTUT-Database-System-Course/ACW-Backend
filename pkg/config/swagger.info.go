package config

import (
	"github.com/NTUT-Database-System-Course/ACW-Backend/docs"
)

func NewSwaggerInfo() {
	docs.SwaggerInfo.Title = "ACW-Backend API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "0.0.1"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

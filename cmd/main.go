package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/sinisaos/gin-vue-starter/api/docs"
	"github.com/sinisaos/gin-vue-starter/pkg/config/db"
	"github.com/sinisaos/gin-vue-starter/pkg/routers"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title						Gin Vue starter API
//	@version					1.0
//	@description				Gin Vue starter project.
//	@securityDefinitions.apiKey	BearerAuth
//	@in							header
//	@name						Authorization
func main() {
	err := godotenv.Load("./pkg/config/envs/example.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	router := gin.Default()
	handler := db.InitDB(os.Getenv("DSN"))

	routers.RegisterRoutes(router, handler)

	docs.SwaggerInfo.BasePath = "/"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(os.Getenv("PORT"))
}

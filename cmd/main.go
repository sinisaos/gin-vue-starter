package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sinisaos/gin-vue-starter/pkg/config/db"
	"github.com/sinisaos/gin-vue-starter/pkg/routers"
)

func main() {
	err := godotenv.Load("./pkg/config/envs/example.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	router := gin.Default()
	handler := db.InitDB(os.Getenv("DSN"))

	routers.RegisterRoutes(router, handler)

	router.Run(os.Getenv("PORT"))
}

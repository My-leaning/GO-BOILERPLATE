package main

import (
	server "go_boilerplate/cmd/server"
	mongoDB "go_boilerplate/internal/database/mongodb"
	postgresDB "go_boilerplate/internal/database/postgres"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go BO-BOILERPLATE
// @version 1.0
// @description This is a sample server for a CRUD and Auth project using Gin and Swagger.
// @host localhost:8080
// @BasePath /
func main() {
	router := gin.Default()

	godotenv.Load(".env")

	mongoDB.MongoConnect()
	postgresDB.PostgresConnect()

	server.RegisterRoutes(router)

	// url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatal(router.Run("localhost:8080"))
}

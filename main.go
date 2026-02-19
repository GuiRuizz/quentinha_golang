package main

import (
	"context"
	"log"
	_ "quentinha_golang/docs"
	"quentinha_golang/src/configuration/database/mongodb"
	"quentinha_golang/src/configuration/logger"
	"quentinha_golang/src/modules/products_modules"
	"quentinha_golang/src/modules/users_modules"
	"quentinha_golang/src/routes/products_routes"
	"quentinha_golang/src/routes/users_routes"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title Meu Primeiro Projeto em Go | Quetinha App Go
// @version 1.0
// @description This is a sample server for user management.
// @host localhost:8080
// @BasePath /
// @schemes http
// @license MIT
const (
	version = "/v1"
)

func main() {
	logger.Info("About to start user applicaion")
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error connecting to the database: %s\n", err.Error())
		return
	}

	router := gin.Default()

	v1 := router.Group(version)

	usersModule := users_modules.NewModule(database)
	productsModule := products_modules.NewModule(database)

	// Routes
	users_routes.InitializeRoutes(v1, usersModule.Controller)
	products_routes.InitializeRoutes(v1, productsModule.Controller)
	// Swagger documentation route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Erro ao iniciar o servidor: ", err)
	}
}

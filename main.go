package main

import (
	"context"
	"log"
	"quentinha_golang/src/configuration/database/mongodb"
	"quentinha_golang/src/configuration/logger"
	"quentinha_golang/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user applicaion")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error connecting to the database: %s\n", err.Error())
		return
	}
	
	userController := initDependencies(database)

	router := gin.Default()
	routes.InitializeRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Erro ao iniciar o servidor: ", err)
	}
}

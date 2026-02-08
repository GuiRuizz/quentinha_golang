package main

import (
	"quentinha_golang/src/controller"
	"quentinha_golang/src/model/repository"
	"quentinha_golang/src/model/service"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func initDependencies(
	database *mongo.Database,
	) (controller.UserControllerInterface) {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
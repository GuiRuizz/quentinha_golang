package users_modules

import (
	"quentinha_golang/src/controller/users_controller"
	"quentinha_golang/src/model/repository/users_repository"
	"quentinha_golang/src/model/service/users_service"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Module struct {
	Controller users_controller.UserControllerInterface
}

func NewModule(database *mongo.Database) Module {
	repo := users_repository.NewUserRepository(database)
	service := users_service.NewUserDomainService(repo)
	controller := users_controller.NewUserControllerInterface(service)

	return Module{
		Controller: controller,
	}
}

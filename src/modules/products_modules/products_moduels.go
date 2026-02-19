package products_modules

import (
	"quentinha_golang/src/controller/products_controller"
	"quentinha_golang/src/model/repository/products_repository"
	"quentinha_golang/src/model/service/products_service"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Module struct {
	Controller products_controller.ProductControllerInterface
}

func NewModule(database *mongo.Database) Module {
	repo := products_repository.NewProductsRepository(database)
	service := products_service.NewProductsDomainService(repo)
	controller := products_controller.NewProductControllerInterface(service)

	return Module{
		Controller: controller,
	}
}

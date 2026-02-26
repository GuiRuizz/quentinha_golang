package address_modules

import (
	"quentinha_golang/src/controller/model/address_controller"
	"quentinha_golang/src/model/repository/address_repository"
	"quentinha_golang/src/model/repository/external"
	"quentinha_golang/src/model/service/address_service"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Module struct {
	Controller address_controller.AddressControllerInterface
}

func NewModule(database *mongo.Database) Module {
	repo := address_repository.NewAddressRepository(database)
	ext := external.ViaCEPClient{}
	service := address_service.NewAddressDomainService(repo, ext)
	controller := address_controller.NewProductControllerInterface(service)

	return Module{
		Controller: controller,
	}
}

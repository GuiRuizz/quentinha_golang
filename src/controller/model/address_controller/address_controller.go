package address_controller

import (
	"quentinha_golang/src/model/service/address_service"

	"github.com/gin-gonic/gin"
)

func NewProductControllerInterface(serviceInterface address_service.AddressDomainService) AddressControllerInterface {
	return &addressControllerInterface{
		service: serviceInterface,
	}
}

type AddressControllerInterface interface {
	CreateAddress(c *gin.Context)
}

type addressControllerInterface struct {
	service address_service.AddressDomainService
}

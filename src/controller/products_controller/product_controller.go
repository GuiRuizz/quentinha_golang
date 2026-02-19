package products_controller

import (
	"quentinha_golang/src/model/service/products_service"
)

func NewProductControllerInterface(serviceInterface products_service.ProductsDomainService) ProductControllerInterface {
	return &productControllerInterface{
		service: serviceInterface,
	}
}

type ProductControllerInterface interface {
	//TODO: Criar os controllers
	// FindProductByID(c *gin.Context)
	// FindAllProduct(c *gin.Context)
	// DeleteProduct(c *gin.Context)
	// UpdateProduct(c *gin.Context)
	// CreateProduct(c *gin.Context)

}

type productControllerInterface struct {
	service products_service.ProductsDomainService
}

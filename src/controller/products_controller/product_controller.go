package products_controller

import (
	"quentinha_golang/src/model/service/products_service"

	"github.com/gin-gonic/gin"
)

func NewProductControllerInterface(serviceInterface products_service.ProductsDomainService) ProductControllerInterface {
	return &productControllerInterface{
		service: serviceInterface,
	}
}

type ProductControllerInterface interface {
	//TODO: Criar os controllers
	FindProductsByID(c *gin.Context)
	FindAllProducts(c *gin.Context)
	// DeleteProduct(c *gin.Context)
	// UpdateProduct(c *gin.Context)
	CreateProduct(c *gin.Context)

}

type productControllerInterface struct {
	service products_service.ProductsDomainService
}

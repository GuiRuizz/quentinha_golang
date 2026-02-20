package products_routes

import (
	"quentinha_golang/src/controller/products_controller"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.RouterGroup, productController products_controller.ProductControllerInterface) {

	// Products routes
	//TODO: Criar as rotas no controller
	r.POST("/createProduct", productController.CreateProduct)
	r.GET("/getAllProducts", productController.FindAllProducts)
	// r.GET("/getProductById/:productId", productController)
	// r.PUT("/updateProduct/:productId", productController.UpdateUser)
	// r.DELETE("/deleteUser/:productId", productController.DeleteUser)

}

package products_routes

import (
	"quentinha_golang/src/controller/products_controller"
	"quentinha_golang/src/model/domain/users_domain"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.RouterGroup, productController products_controller.ProductControllerInterface) {

	// Products routes
	r.POST("/createProduct", productController.CreateProduct)
	r.GET("/getAllProducts", users_domain.VerifyTokenMiddleware, productController.FindAllProducts)
	r.GET("/getProductById/:productId", users_domain.VerifyTokenMiddleware, productController.FindProductsByID)
	r.PUT("/updateProduct/:productId", users_domain.VerifyTokenMiddleware, productController.UpdateProduct)
	r.DELETE("/deleteProduct/:productId", users_domain.VerifyTokenMiddleware, productController.DeleteProduct)

}

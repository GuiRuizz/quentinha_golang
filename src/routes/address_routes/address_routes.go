package address_routes

import (
	"quentinha_golang/src/controller/model/address_controller"
	"quentinha_golang/src/model/domain/users_domain"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.RouterGroup, addressController address_controller.AddressControllerInterface) {

	// Address routes
	r.POST("/createAddress", users_domain.VerifyTokenMiddleware, addressController.CreateAddress)


}

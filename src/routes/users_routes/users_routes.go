package users_routes

import (
	"quentinha_golang/src/controller/users_controller"
	"quentinha_golang/src/model/domain/users_domain"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.RouterGroup, userController users_controller.UserControllerInterface) {

	// User routes
	r.GET("/getUserById/:userId", users_domain.VerifyTokenMiddleware, userController.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", users_domain.VerifyTokenMiddleware, userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)

	// Login routes
	r.POST("/login", userController.LoginUser)

}

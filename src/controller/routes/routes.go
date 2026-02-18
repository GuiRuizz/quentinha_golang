package routes

import (
	"quentinha_golang/src/controller"
	"quentinha_golang/src/model"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	// User routes
	r.GET("/getUserById/:userId", model.VerifyTokenMiddleware, userController.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", model.VerifyTokenMiddleware, userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)

	// Login routes
	r.POST("/login", userController.LoginUser)

}

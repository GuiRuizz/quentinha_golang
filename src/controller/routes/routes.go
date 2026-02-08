package routes

import (
	"quentinha_golang/src/controller"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("/getUserById/:userId", userController.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)
	
}
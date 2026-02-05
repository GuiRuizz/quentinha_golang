package routes

import (
	"quentinha_golang/src/controller"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId", controller.FindUserById)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUserById)
	r.DELETE("/deleteUser/:userId", controller.DeleteUserById)
	
}
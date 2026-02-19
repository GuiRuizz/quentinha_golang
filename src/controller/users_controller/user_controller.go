package users_controller

import (
	"quentinha_golang/src/model/service/users_service"

	"github.com/gin-gonic/gin"
)

func NewUserControllerInterface(serviceInterface users_service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	FindUserByID(c *gin.Context)
	FindUserByEmail(c *gin.Context)

	DeleteUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	CreateUser(c *gin.Context)
	LoginUser(c *gin.Context)
}

type userControllerInterface struct {
	service users_service.UserDomainService
}

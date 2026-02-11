package controller

import (
	"net/http"
	"quentinha_golang/src/configuration/logger"
	"quentinha_golang/src/configuration/validation"
	"quentinha_golang/src/controller/model/request"
	"quentinha_golang/src/model"
	"quentinha_golang/src/view"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInerface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to marshal object", err,
			zap.String("controllers", "createUser"))

		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"),
	)

	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)

	domainResult, err := uc.service.CreateUserServices(domain)
	if err != nil {
		logger.Error("Error trying to call CreateUser service",
			err,
			zap.String("journey", "createUser"),
		)

		c.JSON(err.Code, err)
		return
	}

	logger.Info("CreateUser controller executed successfully",
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "createUser"),
	)


	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domainResult,
	))
}

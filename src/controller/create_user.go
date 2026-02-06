package controller

import (
	"net/http"
	"quentinha_golang/src/configuration/logger"
	"quentinha_golang/src/configuration/validation"
	"quentinha_golang/src/controller/model/request"
	"quentinha_golang/src/controller/model/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
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
	response := response.UserResponse{
		ID:    "teste",
		Name:  userRequest.Name,
		Email: userRequest.Email,
		Age:   userRequest.Age,
	}
	c.JSON(http.StatusOK, response)
}

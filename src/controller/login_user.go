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

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init loginUser controller",
		zap.String("journey", "loginUser"),
	)

	var userRequest request.UserLogin

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to marshal object", err,
			zap.String("controllers", "loginUser"))

		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserLoginDomain(userRequest.Email, userRequest.Password)

	domainResult, token, err := uc.service.LoginUserServices(domain)
	if err != nil {
		logger.Error("Error trying to call loginUser service",
			err,
			zap.String("journey", "loginUser"),
		)

		c.JSON(err.Code, err)
		return
	}

	logger.Info("loginUser controller executed successfully",
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "loginUser"),
	)

	c.Header("Authorization", token)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domainResult,
	))
}

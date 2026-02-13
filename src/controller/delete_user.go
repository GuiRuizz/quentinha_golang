package controller

import (
	"net/http"
	"quentinha_golang/src/configuration/logger"
	"quentinha_golang/src/configuration/rest_err"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init deleteUser controller",
		zap.String("journey", "deleteUser"),
	)

	userId := c.Param("userId")

	if _, err := bson.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to convert userId to hex",
			err,
			zap.String("journey", "deleteUser"),
		)

		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError("UserId is not valid hex"))
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "deleteUser"),
	)

	err := uc.service.DeleteUser(userId)
	if err != nil {
		logger.Error("Error trying to call deleteUser service",
			err,
			zap.String("journey", "deleteUser"),
		)

		c.JSON(err.Code, err)
		return
	}

	logger.Info("deleteUser controller executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"),
	)

	c.Status(http.StatusOK)
}

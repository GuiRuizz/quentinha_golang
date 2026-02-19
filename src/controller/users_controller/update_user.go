package users_controller

import (
	"net/http"
	"quentinha_golang/src/configuration/logger"
	"quentinha_golang/src/configuration/rest_err"
	"quentinha_golang/src/configuration/validation"
	"quentinha_golang/src/controller/model/request"
	"quentinha_golang/src/model/domain/users_domain"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

// UpdateUser updates user information with the specified ID.
// @Summary Update User
// @Description Updates user details based on the ID provided as a parameter.
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "ID of the user to be updated"
// @Param userRequest body request.UserUpdateRequest true "User information for update"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /updateUser/{userId} [put]
func (uc *userControllerInterface) UpdateUser(c *gin.Context) {

	logger.Info("Init updateUser controller",
		zap.String("journey", "updateUser"),
	)

	var userRequest request.UserUpdateRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to marshal object", err,
			zap.String("controllers", "updateUser"))

		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	userId := c.Param("userId")

	if _, err := bson.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to convert userId to hex",
			err,
			zap.String("journey", "updateUser"),
		)

		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError("UserId is not valid hex"))
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "updateUser"),
	)

	domain := users_domain.NewUserUpdateDomain(userRequest.Name, userRequest.Age)

	err := uc.service.UpdateUser(userId, domain)
	if err != nil {
		logger.Error("Error trying to call updateUser service",
			err,
			zap.String("journey", "updateUser"),
		)

		c.JSON(err.Code, err)
		return
	}

	logger.Info("updateUser controller executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"),
	)

	c.Status(http.StatusOK)

}

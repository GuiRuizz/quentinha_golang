package users_service

import (
	"quentinha_golang/src/configuration/logger"
	"quentinha_golang/src/configuration/rest_err"
	"quentinha_golang/src/model/domain/users_domain"

	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUser(userId string, userDomain users_domain.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init updateUser model", zap.String("journey", "updateUserModel"))

	err := ud.userRepository.UpdateUser(userId, userDomain)
	if err != nil {

		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "updateUser"),
		)

		return err
	}

	logger.Info("updateUser service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"),
	)

	return nil
}

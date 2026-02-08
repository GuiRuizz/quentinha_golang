package service

import (
	"quentinha_golang/src/configuration/logger"
	"quentinha_golang/src/configuration/rest_err"
	"quentinha_golang/src/model"

	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser model", zap.String("journey", "createUserModel"))

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {

		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "createUser"),
		)

		return nil, err
	}

	logger.Info("CreateUser service executed successfully",
		zap.String("userId", userDomainRepository.GetID()),
		zap.String("journey", "createUser"),
	)

	return userDomainRepository, nil
}

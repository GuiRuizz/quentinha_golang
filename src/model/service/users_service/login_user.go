package users_service

import (
	"quentinha_golang/src/configuration/logger"
	"quentinha_golang/src/configuration/rest_err"
	"quentinha_golang/src/model/domain/users_domain"

	"go.uber.org/zap"
)

func (ud *userDomainService) LoginUserServices(
	userDomain users_domain.UserDomainInterface,
) (users_domain.UserDomainInterface, string, *rest_err.RestErr) {
	logger.Info("Init loginUser model", zap.String("journey", "loginUserModel"))

	user, err := ud.findUserByEmailAndPasswordServices(userDomain.GetEmail(), userDomain.GetPassword())
	if err != nil {
		return nil, "", err
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, "", err
	}

	logger.Info("LoginUser service executed successfully",
		zap.String("userId", user.GetID()),
		zap.String("journey", "loginUser"),
	)

	return user, token, nil
}

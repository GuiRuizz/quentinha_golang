package service

import (
	"quentinha_golang/src/configuration/logger"
	"quentinha_golang/src/configuration/rest_err"
	"quentinha_golang/src/model"

	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByEmailServices(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmailServices", zap.String("journey", "FindUserByEmailServices"))
	return ud.userRepository.FindUserByEmail(email)
}

func (ud *userDomainService) FindUserByIDServices(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByIDServices 	", zap.String("journey", "FindUserByIDServices"))
	return ud.userRepository.FindUserByID(id)
}

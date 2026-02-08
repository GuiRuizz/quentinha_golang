package service

import (
	"fmt"
	"quentinha_golang/src/configuration/logger"
	"quentinha_golang/src/configuration/rest_err"
	"quentinha_golang/src/model"

	"go.uber.org/zap"
)


func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUserModel"))

	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())

	return nil
}
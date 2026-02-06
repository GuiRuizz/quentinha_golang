package model

import (
	"fmt"
	"quentinha_golang/src/configuration/logger"
	"quentinha_golang/src/configuration/rest_err"

	"go.uber.org/zap"
)


func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUserModel"))

	ud.EncryptPassword()

	fmt.Println(ud)

	return nil
}
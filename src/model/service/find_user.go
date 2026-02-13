package service

import (
	"quentinha_golang/src/configuration/logger"
	"quentinha_golang/src/configuration/rest_err"
	"quentinha_golang/src/model"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func (ud *userDomainService) FindUserByEmailServices(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmailServices", zap.String("journey", "FindUserByEmailServices"))
	return ud.userRepository.FindUserByEmail(email)
}

func (ud *userDomainService) FindUserByIDServices(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByIDServices 	", zap.String("journey", "FindUserByIDServices"))
	return ud.userRepository.FindUserByID(id)
}

func (ud *userDomainService) findUserByEmailAndPasswordServices(
	email string,
	password string,
) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init findUserByEmailAndPassword services.",
		zap.String("journey", "findUserByEmailAndPasswordServices"),
	)

	// 1️⃣ Buscar usuário apenas por email
	user, err := ud.userRepository.FindUserByEmail(email)
	if err != nil {
		return nil, rest_err.NewUnauthorizedError(
			"User or password is invalid",
			nil,
		)
	}

	// 2️⃣ Comparar senha usando bcrypt
	errCompare := bcrypt.CompareHashAndPassword(
		[]byte(user.GetPassword()), // hash salvo no banco
		[]byte(password),           // senha digitada
	)

	if errCompare != nil {
		return nil, rest_err.NewUnauthorizedError(
			"User or password is invalid",
			nil,
		)
	}

	return user, nil
}


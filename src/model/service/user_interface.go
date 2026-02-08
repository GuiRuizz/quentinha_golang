package service

import (
	"quentinha_golang/src/configuration/rest_err"
	"quentinha_golang/src/model"
	"quentinha_golang/src/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{
		userRepository: userRepository,
	}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}

package users_service

import (
	"quentinha_golang/src/configuration/rest_err"
	"quentinha_golang/src/model/domain/users_domain"
	"quentinha_golang/src/model/repository/users_repository"
)

func NewUserDomainService(
	userRepository users_repository.UserRepository,
) UserDomainService {
	return &userDomainService{
		userRepository: userRepository,
	}
}

type userDomainService struct {
	userRepository users_repository.UserRepository
}

type UserDomainService interface {
	CreateUserServices(users_domain.UserDomainInterface) (users_domain.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailServices(email string) (users_domain.UserDomainInterface, *rest_err.RestErr)
	FindUserByIDServices(id string) (users_domain.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(string, users_domain.UserDomainInterface) *rest_err.RestErr
	DeleteUser(string) *rest_err.RestErr
	LoginUserServices(userDomain users_domain.UserDomainInterface) (users_domain.UserDomainInterface, string, *rest_err.RestErr)
}

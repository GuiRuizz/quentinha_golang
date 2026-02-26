package users_repository

import (
	"quentinha_golang/src/configuration/rest_err"
	"quentinha_golang/src/model/domain/users_domain"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewUserRepository(
	databaseConnection *mongo.Database,
) UserRepository {
	return &userRepository{
		databaseConnection: databaseConnection,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(
		userDomain users_domain.UserDomainInterface,
	) (users_domain.UserDomainInterface, *rest_err.RestErr)

	FindUserByEmail(email string) (users_domain.UserDomainInterface, *rest_err.RestErr)

	FindUserByID(id string) (users_domain.UserDomainInterface, *rest_err.RestErr)
	
	FindUserByEmailAndPassword(email string, password string) (users_domain.UserDomainInterface, *rest_err.RestErr)

	UpdateUser(
		userId string,
		userDomain users_domain.UserDomainInterface,
	) *rest_err.RestErr

	DeleteUser(
		userId string,
	) *rest_err.RestErr
}

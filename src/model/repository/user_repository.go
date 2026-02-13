package repository

import (
	"quentinha_golang/src/configuration/rest_err"
	"quentinha_golang/src/model"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

const (
	MONGODB_USER_DB = "MONGODB_USER_DB"
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
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr)

	UpdateUser(
		userId string,
		userDomain model.UserDomainInterface,
	) *rest_err.RestErr

}

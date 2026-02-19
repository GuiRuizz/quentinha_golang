package products_repository

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewProductsRepository(
	databaseConnection *mongo.Database,
) ProductsRepository {
	return &productsRepository{
		databaseConnection: databaseConnection,
	}
}

type productsRepository struct {
	databaseConnection *mongo.Database
}

type ProductsRepository interface {
	// CreateUser(
	// 	userDomain users_domain.UserDomainInterface,
	// ) (users_domain.UserDomainInterface, *rest_err.RestErr)

	// FindUserByEmail(email string) (users_domain.UserDomainInterface, *rest_err.RestErr)

	// FindUserByID(id string) (users_domain.UserDomainInterface, *rest_err.RestErr)

	// UpdateUser(
	// 	userId string,
	// 	userDomain users_domain.UserDomainInterface,
	// ) *rest_err.RestErr

	// DeleteUser(
	// 	userId string,
	// ) *rest_err.RestErr
}

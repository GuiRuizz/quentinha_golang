package products_repository

import (
	"quentinha_golang/src/configuration/rest_err"
	"quentinha_golang/src/model/domain/products_domain"

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
	CreateProduct(
		userDomain products_domain.ProductDomainInterface,
	) (products_domain.ProductDomainInterface, *rest_err.RestErr)

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

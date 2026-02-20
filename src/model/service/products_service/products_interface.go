package products_service

import (
	"quentinha_golang/src/configuration/rest_err"
	"quentinha_golang/src/model/domain/products_domain"
	"quentinha_golang/src/model/repository/products_repository"
)

func NewProductsDomainService(
	productsRepository products_repository.ProductsRepository,
) ProductsDomainService {
	return &productsDomainService{
		productsRepository: productsRepository,
	}
}

type productsDomainService struct {
	productsRepository products_repository.ProductsRepository
}

type ProductsDomainService interface {
	CreateProductServices(products_domain.ProductDomainInterface) (products_domain.ProductDomainInterface, *rest_err.RestErr)
	FindAllProductsServices(offset, limit int64) ([]products_domain.ProductDomainInterface, int64, *rest_err.RestErr)
	// FindUserByIDServices(id string) (users_domain.UserDomainInterface, *rest_err.RestErr)
	// UpdateUser(string, users_domain.UserDomainInterface) *rest_err.RestErr
	// DeleteUser(string) *rest_err.RestErr
	// LoginUserServices(userDomain users_domain.UserDomainInterface) (users_domain.UserDomainInterface, string, *rest_err.RestErr)
}

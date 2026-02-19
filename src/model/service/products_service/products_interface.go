package products_service

import "quentinha_golang/src/model/repository/products_repository"

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
	// CreateUserServices(users_domain.UserDomainInterface) (users_domain.UserDomainInterface, *rest_err.RestErr)
	// FindUserByEmailServices(email string) (users_domain.UserDomainInterface, *rest_err.RestErr)
	// FindUserByIDServices(id string) (users_domain.UserDomainInterface, *rest_err.RestErr)
	// UpdateUser(string, users_domain.UserDomainInterface) *rest_err.RestErr
	// DeleteUser(string) *rest_err.RestErr
	// LoginUserServices(userDomain users_domain.UserDomainInterface) (users_domain.UserDomainInterface, string, *rest_err.RestErr)
}

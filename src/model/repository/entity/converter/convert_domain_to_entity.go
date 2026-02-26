package converter

import (
	"quentinha_golang/src/model/domain/address_domain"
	"quentinha_golang/src/model/domain/products_domain"
	"quentinha_golang/src/model/domain/users_domain"
	"quentinha_golang/src/model/repository/entity/address_entity"
	"quentinha_golang/src/model/repository/entity/product_entity"
	"quentinha_golang/src/model/repository/entity/user_entity"
)

func ConvertDomainToEntity(domain users_domain.UserDomainInterface) *user_entity.UserEntity {
	return &user_entity.UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}
}

func ConvertAddressDomainToEntity(domain address_domain.AddressDomainInterface) *address_entity.AddressEntity {
	return &address_entity.AddressEntity{
		City:    domain.GetCity(),
		Street:  domain.GetStreet(),
		Number:  domain.GetNumber(),
		CEP:     domain.GetCEP(),
		IsValid: domain.GetIsValid(),
	}
}

func ConvertProductDomainToEntity(
	domain products_domain.ProductDomainInterface,
) *product_entity.ProductEntity {

	entity := &product_entity.ProductEntity{
		Name:          domain.GetName(),
		Description:   domain.GetDescription(),
		Value:         domain.GetValue(),
		Images:        domain.GetImages(),
		Stock:         domain.GetStock(),
		RatingAverage: domain.GetRatingAverage(),
		RatingCount:   domain.GetRatingCount(),
		CreatedAt:     domain.GetCreatedAt(),
		UpdatedAt:     domain.GetUpdatedAt(),
		DeletedAt:     domain.GetDeletedAt(),
		IsActive:      *domain.GetIsActive(),
	}
	return entity
}

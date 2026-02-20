package view

import (
	"quentinha_golang/src/controller/model/response"
	"quentinha_golang/src/model/domain/products_domain"
	"quentinha_golang/src/model/domain/users_domain"
)

func ConvertUserDomainToResponse(
	userDomain users_domain.UserDomainInterface,

) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetID(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}

func ConvertProductDomainToResponse(
	productDomain products_domain.ProductDomainInterface,
) response.ProductResponse {

	return response.ProductResponse{
		ID:            productDomain.GetID(),
		Name:          productDomain.GetName(),
		Description:   productDomain.GetDescription(),
		Value:         productDomain.GetValue(),
		Images:        productDomain.GetImages(),
		Stock:         productDomain.GetStock(),
		RatingAverage: productDomain.GetRatingAverage(),
		RatingCount:   productDomain.GetRatingCount(),
		CreatedAt:     productDomain.GetCreatedAt(),
		UpdatedAt:     productDomain.GetUpdatedAt(),
		IsActive:      *productDomain.GetIsActive(),
	}
}

func ConvertProductDomainListToResponse(
	domains []products_domain.ProductDomainInterface,
) []response.ProductResponse {

	var responses []response.ProductResponse

	for _, domain := range domains {
		responses = append(responses, ConvertProductDomainToResponse(domain))
	}

	return responses
}

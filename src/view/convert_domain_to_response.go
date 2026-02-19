package view

import (
	"quentinha_golang/src/controller/model/response"
	"quentinha_golang/src/model/domain/users_domain"
)

func ConvertDomainToResponse(
	userDomain users_domain.UserDomainInterface,

) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetID(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}

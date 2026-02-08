package view

import (
	"quentinha_golang/src/controller/model/response"
	"quentinha_golang/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,

) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}

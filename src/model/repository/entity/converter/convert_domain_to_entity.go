package converter

import (
	"quentinha_golang/src/model/domain/users_domain"
	"quentinha_golang/src/model/repository/entity"
)

func ConvertDomainToEntity(domain users_domain.UserDomainInterface) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}
}

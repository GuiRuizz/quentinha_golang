package converter

import (
	"quentinha_golang/src/model/domain/users_domain"
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

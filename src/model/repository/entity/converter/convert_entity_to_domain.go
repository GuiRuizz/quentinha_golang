package converter

import (
	"quentinha_golang/src/model/domain/users_domain"
	"quentinha_golang/src/model/repository/entity"
)

func ConvertEntityToDomain(entity entity.UserEntity) users_domain.UserDomainInterface {
	domain := users_domain.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)
	domain.SetID(entity.ID.Hex())
	return domain
}


package converter

import (
	"quentinha_golang/src/model/domain/products_domain"
	"quentinha_golang/src/model/domain/users_domain"
	"quentinha_golang/src/model/repository/entity/product_entity"
	"quentinha_golang/src/model/repository/entity/user_entity"
)

func ConvertEntityToDomain(entity user_entity.UserEntity) users_domain.UserDomainInterface {
	domain := users_domain.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)
	domain.SetID(entity.ID.Hex())
	return domain
}

func ConvertProductEntityToDomain(
	entity product_entity.ProductEntity,
) products_domain.ProductDomainInterface {

	domain := products_domain.NewProductDomain(
		entity.Name,
		entity.Description,
		entity.Value,
		entity.Images,
		entity.Stock,
	)

	domain.SetID(entity.ID.Hex())
	domain.SetRatingAverage(entity.RatingAverage)
	domain.SetRatingCount(entity.RatingCount)
	domain.SetCreatedAt(entity.CreatedAt)
	domain.SetUpdatedAt()
	domain.SetDeletedAt(entity.DeletedAt)

	return domain
}

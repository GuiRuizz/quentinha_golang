package products_domain

import "time"

type ProductDomainInterface interface {
	GetID() string
	GetName() string
	GetDescription() string
	GetValue() int64
	GetImages() []string
	GetRatingAverage() *float64
	GetRatingCount() int32
	GetStock() int32
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetDeletedAt() *time.Time

	SetID(string)
	SetUpdatedAt()
	SetDeleted()
	SetCreatedAt(time.Time)
	SetUpdatedAtFromDB(time.Time)
	SetDeletedAt(*time.Time)
	SetRatingAverage(*float64)
	SetRatingCount(int32)
}

func NewProductDomain(
	name string,
	description string,
	value int64,
	images []string,
	stock int32,
) ProductDomainInterface {

	now := time.Now()

	return &productDomain{
		name:          name,
		description:   description,
		value:         value,
		images:        images,
		stock:         stock,
		ratingAverage: nil, 
		ratingCount:   0,
		createdAt:     now,
		updatedAt:     now,
	}
}


//TODO: Fazer o update domain
// func NewUserUpdateDomain(
// 	name string,
// 	age int8,
// ) ProductDomainInterface {
// 	return &userDomain{
// 		name: name,
// 		age:  age,
// 	}
// }

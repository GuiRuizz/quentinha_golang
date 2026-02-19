package products_domain

type ProductDomainInterface interface {
	SetID(string)
	GetID() string
	GetDescription() string
	GetValue() int
	GetImages() []string
	GetRating() int8
	GetStock() int8
}

func NewProductsDomain(
	description, name string,
	rating, stock int8,
	value int,
	image []string,
) ProductDomainInterface {
	return &productsDomain{
		name:        name,
		description: description,
		value:       value,
		image:       image,
		rating:      rating,
		stock:       stock,
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

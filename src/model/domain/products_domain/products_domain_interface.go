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
	GetIsActive() *bool

	SetID(string)
	SetUpdatedAt()
	SetDeleted()
	SetCreatedAt(time.Time)
	SetUpdatedAtFromDB(time.Time)
	SetDeletedAt(*time.Time)
	SetRatingAverage(*float64)
	SetRatingCount(int32)
	SetIsActive(bool)
}

func NewProductDomain(
	name string,
	description string,
	value int64,
	images []string,
	stock int32,
) ProductDomainInterface {

	now := time.Now()
	active := true

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
		isActive:      &active,
	}
}

func NewProductUpdateDomain(
	name string,
	description string,
	value int64,
	images []string,
	stock int32,
	isActive *bool,
) ProductDomainInterface {

	now := time.Now()

	return &productDomain{
		name:        name,
		description: description,
		value:       value,
		images:      images,
		stock:       stock,
		updatedAt:   now,
		isActive:    isActive,
	}
}

package address_repository

import (
	"quentinha_golang/src/configuration/rest_err"
	"quentinha_golang/src/model/domain/address_domain"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewAddressRepository(
	databaseConnection *mongo.Database,
) AddressRepository {
	return &addressRepository{
		databaseConnection: databaseConnection,
	}
}

type addressRepository struct {
	databaseConnection *mongo.Database
}

type AddressRepository interface {
	CreateAddress(
		addressDomain address_domain.AddressDomainInterface,
	) (address_domain.AddressDomainInterface, *rest_err.RestErr)

	CityExists(
		city string,
	) (bool, *rest_err.RestErr)
}
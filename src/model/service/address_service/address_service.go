package address_service

import (
	"quentinha_golang/src/configuration/rest_err"
	"quentinha_golang/src/model/domain/address_domain"
	"quentinha_golang/src/model/repository/address_repository"
	"quentinha_golang/src/model/repository/external"
)

func NewAddressDomainService(
	addressRepository address_repository.AddressRepository,
	viaCepClient external.ViaCEPClient,
) AddressDomainService {
	return &addressDomainService{
		addressRepository: addressRepository,

		viaCepClient: viaCepClient,
	}
}

type addressDomainService struct {
	addressRepository address_repository.AddressRepository

	viaCepClient external.ViaCEPClient
}

type AddressDomainService interface {
	CreateAddressServices(address_domain.AddressDomainInterface) (address_domain.AddressDomainInterface, *rest_err.RestErr)
	// FindAddressByIDServices(id string) (address_domain.AddressDomainInterface, *rest_err.RestErr)
	// FindAllAddressByUserServices(userID string) ([]address_domain.AddressDomainInterface, *rest_err.RestErr)
	// DeleteAddressServices(id string) *rest_err.RestErr
}

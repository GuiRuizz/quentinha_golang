package view

import (
	"quentinha_golang/src/controller/model/response"
	"quentinha_golang/src/model/domain/address_domain"
)

func ConvertAddressDomainToResponse(
	addressDomain address_domain.AddressDomainInterface,

) response.AddressResponse {
	return response.AddressResponse{
		ID:      addressDomain.GetID(),
		Street:  addressDomain.GetStreet(),
		City:    addressDomain.GetCity(),
		Number:  addressDomain.GetNumber(),
		CEP:     addressDomain.GetCEP(),
		IsValid: addressDomain.GetIsValid(),
	}
}

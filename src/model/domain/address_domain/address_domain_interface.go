package address_domain

type AddressDomainInterface interface {
	GetID() string
	GetCity() string
	GetStreet() string
	GetNumber() string
	GetCEP() string
	GetIsValid() bool

	SetID(string)
	SetIsValid(bool)
}

func NewAddressDomain(
	city string,
	street string,
	number string,
	cep string,
) AddressDomainInterface {

	valid := true

	return &addressDomain{
		city:   city,
		street: street,
		number: number,
		cep:    cep,
		isValid: valid,
	}
}

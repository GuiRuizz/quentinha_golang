package address_domain

type AllowedCityDomainInterface interface {
	GetName() string
}

func NewAllowedCityDomain(
	name string,

) AllowedCityDomainInterface {

	return &allowedCityDomain{
		name: name,
	}
}

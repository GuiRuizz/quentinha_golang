package address_domain

type allowedCityDomain struct {
	id   string
	name string
}

func (acd *allowedCityDomain) GetName() string {
	return acd.name
}

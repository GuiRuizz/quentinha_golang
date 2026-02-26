package address_domain

type addressDomain struct {
	id      string
	city    string
	street  string
	number  string
	cep     string
	isValid bool
}


func (ad *addressDomain) GetID() string {
	return ad.id
}

func (ad *addressDomain) GetIsValid() bool {
	return ad.isValid
}

func (ad *addressDomain) GetCity() string {
	return ad.city
}

func (ad *addressDomain) GetStreet() string {
	return ad.street
}

func (ad *addressDomain) GetNumber() string {
	return ad.number
}

func (ad *addressDomain) GetCEP() string {
	return ad.cep
}


func (ad *addressDomain) SetID(id string) {
	ad.id = id
}

func (ad *addressDomain) SetIsValid(isValid bool) {
	ad.isValid = isValid
}
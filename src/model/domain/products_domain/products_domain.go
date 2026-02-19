package products_domain

type productsDomain struct {
	id          string
	name        string
	description string
	value       int
	image       []string
	rating      int8
	stock 		int8
}

func (pd *productsDomain) SetID(id string) {
	pd.id = id
}

func (pd *productsDomain) GetID() string {
	return pd.id
}

func (pd *productsDomain) GetDescription() string {
	return pd.description
}
func (pd *productsDomain) GetValue() int {
	return pd.value
}
func (pd *productsDomain) GetImages() []string {
	return pd.image
}
func (pd *productsDomain) GetRating() int8 {
	return pd.rating
}
func (pd *productsDomain) GetStock() int8 {
	return pd.stock
}
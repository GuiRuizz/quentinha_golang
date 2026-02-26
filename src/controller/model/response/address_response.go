package response

type AddressResponse struct {
	ID      string `json:"id"`
	City    string `json:"city"`
	Street  string `json:"street"`
	Number  string `json:"number"`
	CEP     string `json:"cep"`
	IsValid bool   `json:"isValid"`
}

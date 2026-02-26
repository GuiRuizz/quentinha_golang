package address_request

// AddressRequest represents the input data for validating and registering a delivery address.
// @Summary Address Input Data
// @Description Structure containing the required fields for validating a delivery address.
type AddressRequest struct {

	// City where the delivery will be made (must be allowed in the system).
	// Example: São José dos Campos
	// @json
	// @jsonTag city
	// @jsonExample São José dos Campos
	// @binding required,min=2,max=100
	City string `json:"city" binding:"required,min=2,max=100" example:"São José dos Campos"`

	// Street name corresponding to the provided CEP.
	// Example: Rua Sumaré
	// @json
	// @jsonTag street
	// @jsonExample Rua Sumaré
	// @binding required,min=2,max=150
	Street string `json:"street" binding:"required,min=2,max=150" example:"Rua Sumaré"`

	// House/building number (must be greater than zero).
	// Example: "155" or "H1C"
	// @json
	// @jsonTag number
	// @jsonExample "155"
	// @binding required,min=1
	Number string `json:"number" binding:"required,min=1" example:"155"`

	// Brazilian ZIP Code (CEP) in format 00000-000 or only numbers.
	// Example: 12233-770
	// @json
	// @jsonTag cep
	// @jsonExample 12233-770
	// @binding required,len=9
	CEP string `json:"cep" binding:"required,len=9" example:"12233-770"`
}
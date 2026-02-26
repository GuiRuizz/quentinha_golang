package external

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ViaCEPResponse struct {
	Logradouro string `json:"logradouro"`
	Localidade string `json:"localidade"`
	Erro       bool   `json:"erro"`
}

type ViaCEPClient struct{}

func (v *ViaCEPClient) GetAddress(cep string) (*ViaCEPResponse, error) {

	resp, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ViaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if result.Erro {
		return nil, errors.New("cep inválido")
	}

	return &result, nil
}
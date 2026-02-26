package address_service

import (
	"strings"

	"quentinha_golang/src/configuration/logger"
	"quentinha_golang/src/configuration/rest_err"
	"quentinha_golang/src/model/domain/address_domain"

	"go.uber.org/zap"
)

func (ad *addressDomainService) CreateAddressServices(
	addressDomain address_domain.AddressDomainInterface,
) (address_domain.AddressDomainInterface, *rest_err.RestErr) {

	logger.Info("Init createAddress service",
		zap.String("journey", "createAddressService"),
	)

	// 1️⃣ verificar cidade permitida
	exists, errCity := ad.addressRepository.CityExists(addressDomain.GetCity())
	if errCity != nil {

		logger.Error("Error trying to validate allowed city",
			errCity,
			zap.String("journey", "createAddressService"),
		)

		return nil, errCity
	}

	if !exists {

		logger.Error("City not allowed",
			nil,
			zap.String("city", addressDomain.GetCity()),
			zap.String("journey", "createAddressService"),
		)

		return nil, rest_err.NewNotFoundError(
			"Sua cidade não está válida para entregas de quentinha",
		)
	}

	// 2️⃣ validar CEP via API externa
	viaCepResponse, errVia := ad.viaCepClient.GetAddress(addressDomain.GetCEP())
	if errVia != nil {

		logger.Error("Error trying to validate CEP",
			errVia,
			zap.String("journey", "createAddressService"),
		)

		return nil, rest_err.NewBadRequestError("CEP inválido")
	}

	// 3️⃣ validar cidade do CEP
	if !strings.EqualFold(viaCepResponse.Localidade, addressDomain.GetCity()) {

		logger.Error("City mismatch with CEP",
			nil,
			zap.String("journey", "createAddressService"),
		)

		return nil, rest_err.NewBadRequestError(
			"Cidade divergente do CEP informado",
		)
	}

	// 4️⃣ validar rua
	if !strings.EqualFold(viaCepResponse.Logradouro, addressDomain.GetStreet()) {

		logger.Error("Street mismatch with CEP",
			nil,
			zap.String("journey", "createAddressService"),
		)

		return nil, rest_err.NewBadRequestError(
			"Rua divergente do CEP informado",
		)
	}

	// 5️⃣ chamar repository
	addressDomainRepository, err := ad.addressRepository.CreateAddress(addressDomain)
	if err != nil {

		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "createAddressService"),
		)

		return nil, err
	}

	logger.Info("CreateAddress service executed successfully",
		zap.String("addressId", addressDomainRepository.GetID()),
		zap.String("journey", "createAddressService"),
	)

	return addressDomainRepository, nil
}

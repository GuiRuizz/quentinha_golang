package address_repository

import (
	"context"
	"os"
	"quentinha_golang/src/configuration/logger"
	"quentinha_golang/src/configuration/rest_err"
	"quentinha_golang/src/model/domain/address_domain"
	"quentinha_golang/src/model/repository/entity/converter"
	"quentinha_golang/src/model/repository/utils"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

func (ar *addressRepository) CreateAddress(
	addressDomain address_domain.AddressDomainInterface,
) (address_domain.AddressDomainInterface, *rest_err.RestErr) {

	logger.Info("Init creating address in database",
		zap.String("journey", "createAddressRepository"),
	)

	collectionName := os.Getenv(utils.MONGODB_ADDRESS_DB)
	collection := ar.databaseConnection.Collection(collectionName)

	value := converter.ConvertAddressDomainToEntity(addressDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {

		logger.Error("Error trying to create address",
			err,
			zap.String("journey", "createAddressRepository"),
		)

		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(bson.ObjectID)

	logger.Info("CreateAddress repository executed successfully",
		zap.String("addressId", value.ID.Hex()),
		zap.String("journey", "createAddressRepository"),
	)

	return converter.ConvertAddressEntityToDomain(*value), nil
}
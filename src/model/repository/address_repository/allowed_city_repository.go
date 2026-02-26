package address_repository

import (
	"context"
	"os"
	"quentinha_golang/src/configuration/logger"
	"quentinha_golang/src/configuration/rest_err"
	"quentinha_golang/src/model/repository/utils"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

func (ar *addressRepository) CityExists(
	city string,
) (bool, *rest_err.RestErr) {

	logger.Info("Init checking city in database",
		zap.String("journey", "cityExistsRepository"),
	)

	collectionName := os.Getenv(utils.MONGODB_ALLOWED_CITIES_DB)
	collection := ar.databaseConnection.Collection(collectionName)

	filter := bson.M{
		"name": city,
	}

	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {

		logger.Error("Error trying to check city",
			err,
			zap.String("journey", "cityExistsRepository"),
		)

		return false, rest_err.NewInternalServerError(err.Error())
	}

	return count > 0, nil
}
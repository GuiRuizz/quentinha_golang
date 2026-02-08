package repository

import (
	"context"
	"os"
	"quentinha_golang/src/configuration/logger"
	"quentinha_golang/src/configuration/rest_err"
	"quentinha_golang/src/model"
	"quentinha_golang/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

var (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init creating user in the database",

		zap.String("journey", "createUserRepository"),
	)
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error trying to create user",
			err,
			zap.String("journey", "createUser"))

		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(bson.ObjectID)

	logger.Info("CreateUser repository executed successfully",
		zap.String("userId", value.ID.Hex()),
		zap.String("journey", "createUser"),
	)


	return converter.ConvertEntityToDomain(*value), nil

}

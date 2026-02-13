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

func (ur *userRepository) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {

	logger.Info("Init updating user in the database",

		zap.String("journey", "updateUserRepository"),
	)
	collection_name := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)

	userIdHex, _ := bson.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: userIdHex}}

	update := bson.D{{Key: "$set", Value: value}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error("Error trying to update user",
			err,
			zap.String("journey", "updateUser"))

		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("updateUser repository executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"),
	)

	return nil

}

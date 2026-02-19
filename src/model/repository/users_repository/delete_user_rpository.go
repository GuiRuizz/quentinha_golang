package users_repository

import (
	"context"
	"os"
	"quentinha_golang/src/configuration/logger"
	"quentinha_golang/src/configuration/rest_err"
	"quentinha_golang/src/model/repository/utils"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUser(
	userId string,
) *rest_err.RestErr {

	logger.Info("Init deleting user in the database",

		zap.String("journey", "deleteUserRepository"),
	)
	collection_name := os.Getenv(utils.MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collection_name)

	userIdHex, _ := bson.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: userIdHex}}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		logger.Error("Error trying to delete user",
			err,
			zap.String("journey", "deleteUser"))

		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("deleteUser repository executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "delete"),
	)

	return nil

}

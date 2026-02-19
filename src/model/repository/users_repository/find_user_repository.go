package users_repository

import (
	"context"
	"fmt"
	"os"
	"quentinha_golang/src/configuration/logger"
	"quentinha_golang/src/configuration/rest_err"

	"quentinha_golang/src/model/domain/users_domain"
	"quentinha_golang/src/model/repository/entity"
	"quentinha_golang/src/model/repository/entity/converter"
	"quentinha_golang/src/model/repository/utils"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(email string) (users_domain.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init finding user in the database",

		zap.String("journey", "findUserByEmail"),
	)
	collection_name := os.Getenv(utils.MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMssage := fmt.Sprintf(
				"User not found with this email: %s", email)
			logger.Error(errorMssage,
				err,
				zap.String("journey", "findUserByEmail"))
			return nil, rest_err.NewNotFoundError(errorMssage)
		}

		errorMssage := "Error trying to find user by email"
		logger.Error(errorMssage,
			err,
			zap.String("journey", "findUserByEmail"))
		return nil, rest_err.NewInternalServerError(errorMssage)

	}

	logger.Info("FindUserByEmail repository executed successfully",
		zap.String("email", email),
		zap.String("userId", userEntity.ID.Hex()),
		zap.String("journey", "findUserByEmail"))

	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByID(
	id string,
) (users_domain.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByID repository",
		zap.String("journey", "findUserByID"))

	collection_name := os.Getenv(utils.MONGODB_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	objectId, _ := bson.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User not found with this ID: %s", id)
			logger.Error(errorMessage,
				err,
				zap.String("journey", "findUserByID"))

			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by ID"
		logger.Error(errorMessage,
			err,
			zap.String("journey", "findUserByID"))

		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByID repository executed successfully",
		zap.String("journey", "findUserByID"),
		zap.String("userId", userEntity.ID.Hex()))
	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByEmailAndPassword(
	email string,
	password string,
) (users_domain.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmailAndPassword repository",
		zap.String("journey", "findUserByEmailAndPassword"))

	collection_name := os.Getenv(utils.MONGODB_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{
		{Key: "email", Value: email},
		{Key: "password", Value: password},
	}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := "User or password is invalid"
			logger.Error(errorMessage,
				err,
				zap.String("journey", "findUserByEmailAndPassword"))

			return nil, rest_err.NewUnauthorizedError(errorMessage, []rest_err.Causes{})
		}
		errorMessage := "Error trying to find user by email and password"
		logger.Error(errorMessage,
			err,
			zap.String("journey", "findUserByEmailAndPassword"))

		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByEmailAndPassword repository executed successfully",
		zap.String("journey", "findUserByEmailAndPassword"),
		zap.String("email", email),
		zap.String("userId", userEntity.ID.Hex()))
	return converter.ConvertEntityToDomain(*userEntity), nil
}

package products_repository

import (
	"context"
	"fmt"
	"os"
	"quentinha_golang/src/configuration/logger"
	"quentinha_golang/src/configuration/rest_err"
	"quentinha_golang/src/model/domain/products_domain"
	"quentinha_golang/src/model/repository/entity/converter"
	"quentinha_golang/src/model/repository/entity/product_entity"
	"quentinha_golang/src/model/repository/utils"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.uber.org/zap"
)

func (pr *productsRepository) FindAllProducts(
	offset int64,
	limit int64,
) ([]products_domain.ProductDomainInterface, int64, *rest_err.RestErr) {

	logger.Info("Init findAllProducts repository",
		zap.String("journey", "findAllProducts"))

	collectionName := os.Getenv(utils.MONGODB_PRODUCT_DB)
	collection := pr.databaseConnection.Collection(collectionName)

	ctx := context.Background()

	total, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		errorMessage := "Error trying to find products"
		logger.Error(errorMessage, err,
			zap.String("journey", "findAllProducts"))

		return nil, 0, rest_err.NewInternalServerError(errorMessage)
	}

	var productsEntity []product_entity.ProductEntity

	findOptions := options.Find()
	findOptions.SetSkip(offset)
	findOptions.SetLimit(limit)
	cursor, err := collection.Find(ctx, bson.M{}, findOptions)

	if err := cursor.All(context.Background(), &productsEntity); err != nil {
		errorMessage := "Error decoding products"
		logger.Error(errorMessage, err,
			zap.String("journey", "findAllProducts"))

		return nil, 0, rest_err.NewInternalServerError(errorMessage)
	}

	if len(productsEntity) == 0 {
		return nil, 0, rest_err.NewNotFoundError("There isn't any products saved")
	}

	var productsDomain []products_domain.ProductDomainInterface

	for _, product := range productsEntity {
		productsDomain = append(productsDomain,
			converter.ConvertProductEntityToDomain(product))
	}

	logger.Info("FindAllProducts repository executed successfully",
		zap.String("journey", "findAllProducts"))

	return productsDomain, total, nil
}

func (pr *productsRepository) FindProductsByID(
	id string,
) (products_domain.ProductDomainInterface, *rest_err.RestErr) {

	logger.Info("Init findAllProducts repository",
		zap.String("journey", "findAllProducts"))

	collectionName := os.Getenv(utils.MONGODB_PRODUCT_DB)
	collection := pr.databaseConnection.Collection(collectionName)

	ctx := context.Background()

	productsEntity := &product_entity.ProductEntity{}

	objectId, _ := bson.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: objectId}}

	err := collection.FindOne(
		ctx,
		filter,
	).Decode(productsEntity)
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
		zap.String("userId", productsEntity.ID.Hex()))

	return converter.ConvertProductEntityToDomain(*productsEntity), nil
}

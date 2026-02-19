package products_repository

import (
	"context"
	"os"
	"quentinha_golang/src/configuration/logger"
	"quentinha_golang/src/configuration/rest_err"
	"quentinha_golang/src/model/domain/products_domain"
	"quentinha_golang/src/model/repository/entity/converter"
	"quentinha_golang/src/model/repository/utils"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

func (pr *productsRepository) CreateProduct(
	productDomain products_domain.ProductDomainInterface,
) (products_domain.ProductDomainInterface, *rest_err.RestErr) {

	logger.Info("Init creating product in the database",
		zap.String("journey", "createProductRepository"),
	)

	collectionName := os.Getenv(utils.MONGODB_PRODUCT_DB)

	collection := pr.databaseConnection.Collection(collectionName)

	value := converter.ConvertProductDomainToEntity(productDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {

		logger.Error("Error trying to create product",
			err,
			zap.String("journey", "createProductRepository"),
		)

		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(bson.ObjectID)

	logger.Info("CreateProduct repository executed successfully",
		zap.String("productId", value.ID.Hex()),
		zap.String("journey", "createProductRepository"),
	)

	return converter.ConvertProductEntityToDomain(*value), nil
}

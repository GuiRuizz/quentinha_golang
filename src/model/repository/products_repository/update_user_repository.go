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

func (pr *productsRepository) UpdateProduct(
	productId string,
	productDomain products_domain.ProductDomainInterface,
) *rest_err.RestErr {

	logger.Info("Init updating product in the database",

		zap.String("journey", "updateProductRepository"),
	)
	collection_name := os.Getenv(utils.MONGODB_PRODUCT_DB)

	collection := pr.databaseConnection.Collection(collection_name)

	value := converter.ConvertProductDomainToEntity(productDomain)

	productIdHex, _ := bson.ObjectIDFromHex(productId)

	filter := bson.D{{Key: "_id", Value: productIdHex}}

	update := bson.D{{Key: "$set", Value: value}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error("Error trying to update product",
			err,
			zap.String("journey", "updateProduct"))

		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("updateProduct repository executed successfully",
		zap.String("productId", productId),
		zap.String("journey", "updateProduct"),
	)

	return nil

}

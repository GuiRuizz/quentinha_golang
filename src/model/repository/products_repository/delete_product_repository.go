package products_repository

import (
	"context"
	"os"
	"quentinha_golang/src/configuration/logger"
	"quentinha_golang/src/configuration/rest_err"
	"quentinha_golang/src/model/repository/utils"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

func (pr *productsRepository) DeleteProduct(
	productId string,
) *rest_err.RestErr {

	logger.Info("Init deleting product in the database",

		zap.String("journey", "deleteProductRepository"),
	)
	collection_name := os.Getenv(utils.MONGODB_PRODUCT_DB)

	collection := pr.databaseConnection.Collection(collection_name)

	productIdHex, _ := bson.ObjectIDFromHex(productId)

	filter := bson.D{{Key: "_id", Value: productIdHex}}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		logger.Error("Error trying to delete product",
			err,
			zap.String("journey", "deleteproduct"))

		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("deleteproduct repository executed successfully",
		zap.String("productId", productId),
		zap.String("journey", "delete"),
	)

	return nil

}

package products_service

import (
	"quentinha_golang/src/configuration/logger"
	"quentinha_golang/src/configuration/rest_err"

	"go.uber.org/zap"
)

func (pd *productsDomainService) DeleteProductServices(productId string) *rest_err.RestErr {

	logger.Info("Init deleteProduct model", zap.String("journey", "deleteProductModel"))

	err := pd.productsRepository.DeleteProduct(productId)
	if err != nil {

		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "deleteProduct"),
		)

		return err
	}

	logger.Info("deleteProduct service executed successfully",
		zap.String("productId", productId),
		zap.String("journey", "deleteProduct"),
	)

	return nil
}

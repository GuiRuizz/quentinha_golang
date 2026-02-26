package products_service

import (
	"quentinha_golang/src/configuration/logger"
	"quentinha_golang/src/configuration/rest_err"
	"quentinha_golang/src/model/domain/products_domain"

	"go.uber.org/zap"
)

func (pd *productsDomainService) UpdateProductServices(productId string, productDomain products_domain.ProductDomainInterface) *rest_err.RestErr {
	logger.Info("Init updateProduct model", zap.String("journey", "updateProductModel"))

	err := pd.productsRepository.UpdateProduct(productId, productDomain)
	if err != nil {

		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "updateProduct"),
		)

		return err
	}

	logger.Info("updateProduct service executed successfully",
		zap.String("productId", productId),
		zap.String("journey", "updateProduct"),
	)

	return nil
}

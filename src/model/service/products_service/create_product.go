package products_service

import (
	"quentinha_golang/src/configuration/logger"
	"quentinha_golang/src/configuration/rest_err"
	"quentinha_golang/src/model/domain/products_domain"

	"go.uber.org/zap"
)

func (pd *productsDomainService) CreateProductServices(
	productDomain products_domain.ProductDomainInterface,
) (products_domain.ProductDomainInterface, *rest_err.RestErr) {

	logger.Info("Init createProduct service",
		zap.String("journey", "createProductService"),
	)

	productDomainRepository, err := pd.productsRepository.CreateProduct(productDomain)
	if err != nil {

		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "createProductService"),
		)

		return nil, err
	}

	logger.Info("CreateProduct service executed successfully",
		zap.String("productId", productDomainRepository.GetID()),
		zap.String("journey", "createProductService"),
	)

	return productDomainRepository, nil
}

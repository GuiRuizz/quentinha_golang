package products_service

import (
	"quentinha_golang/src/configuration/logger"
	"quentinha_golang/src/configuration/rest_err"
	"quentinha_golang/src/model/domain/products_domain"

	"go.uber.org/zap"
)

func (pd *productsDomainService) FindAllProductsServices(offset int64,
	limit int64) ([]products_domain.ProductDomainInterface, int64, *rest_err.RestErr) {
	logger.Info("Init FindAllProductsServices 	", zap.String("journey", "FindAllProductsServices"))
	return pd.productsRepository.FindAllProducts(offset, limit)
}

func (pd *productsDomainService) FindProductsByIDServices(id string) (products_domain.ProductDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindProductsByIDServices 	", zap.String("journey", "FindProductsByIDServices"))
	return pd.productsRepository.FindProductsByID(id)
}

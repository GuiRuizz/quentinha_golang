package products_controller

import (
	"net/http"
	"quentinha_golang/src/configuration/logger"
	"quentinha_golang/src/configuration/validation"
	"quentinha_golang/src/controller/model/request/product_request"
	"quentinha_golang/src/model/domain/products_domain"
	"quentinha_golang/src/view"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateProduct Creates a new product
// @Summary Create a new product
// @Description Create a new product with the provided product information
// @Tags Products
// @Accept json
// @Produce json
// @Param productRequest body product_request.ProductRequest true "Product information"
// @Success 200 {object} response.ProductResponse
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /products [post]
func (pc *productControllerInterface) CreateProduct(c *gin.Context) {

	logger.Info("Init CreateProduct controller",
		zap.String("journey", "createProduct"),
	)

	var productRequest product_request.ProductRequest

	if err := c.ShouldBindJSON(&productRequest); err != nil {
		logger.Error("Error trying to bind productRequest", err,
			zap.String("controller", "createProduct"),
		)

		errRest := validation.ValidateError(err, "Product")

		c.JSON(errRest.Code, errRest)
		return
	}

	logger.Info("Product request validated successfully",
		zap.String("journey", "createProduct"),
	)

	domain := products_domain.NewProductDomain(
		productRequest.Name,
		productRequest.Description,
		productRequest.Value,
		productRequest.Images,
		productRequest.Stock,
	)

	domainResult, err := pc.service.CreateProductServices(domain)
	if err != nil {
		logger.Error("Error trying to call CreateUserProductService",
			err,
			zap.String("journey", "createProduct"),
		)

		c.JSON(err.Code, err)
		return
	}

	logger.Info("CreateProduct controller executed successfully",
		zap.String("productId", domainResult.GetID()),
		zap.String("journey", "createProduct"),
	)

	c.JSON(http.StatusOK, view.ConvertProductDomainToResponse(domainResult))
}

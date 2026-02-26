package products_controller

import (
	"net/http"
	"quentinha_golang/src/configuration/logger"
	"quentinha_golang/src/configuration/rest_err"
	"quentinha_golang/src/configuration/validation"
	"quentinha_golang/src/controller/model/request/product_request"
	"quentinha_golang/src/model/domain/products_domain"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

// UpdateProduct updates Product information with the specified ID.
// @Summary Update Product
// @Description Updates Product details based on the ID provided as a parameter.
// @Tags Products
// @Accept json
// @Produce json
// @Param ProductId path string true "ID of the Product to be updated"
// @Param ProductRequest body request.ProductUpdateRequest true "Product information for update"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /updateProduct/{ProductId} [put]
func (pc *productControllerInterface) UpdateProduct(c *gin.Context) {

	logger.Info("Init updateProduct controller",
		zap.String("journey", "updateProduct"),
	)

	var productRequest product_request.ProductUpdateRequest

	if err := c.ShouldBindJSON(&productRequest); err != nil {
		logger.Error("Error trying to marshal object", err,
			zap.String("controllers", "updateProduct"))

		errRest := validation.ValidateError(err, "User")

		c.JSON(errRest.Code, errRest)
		return
	}

	productId := c.Param("productId")

	if _, err := bson.ObjectIDFromHex(productId); err != nil {
		logger.Error("Error trying to convert productId to hex",
			err,
			zap.String("journey", "updateProduct"),
		)

		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError("productId is not valid hex"))
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "updateProduct"),
	)

	domain := products_domain.NewProductUpdateDomain(
		productRequest.Name,
		productRequest.Description,
		productRequest.Value,
		productRequest.Images,
		productRequest.Stock, 
		productRequest.IsActive)

	err := pc.service.UpdateProductServices(productId, domain)
	if err != nil {
		logger.Error("Error trying to call updateProduct service",
			err,
			zap.String("journey", "updateProduct"),
		)

		c.JSON(err.Code, err)
		return
	}

	logger.Info("updateProduct controller executed successfully",
		zap.String("productId", productId),
		zap.String("journey", "updateProduct"),
	)

	c.Status(http.StatusOK)

}

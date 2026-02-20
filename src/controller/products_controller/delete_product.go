package products_controller

import (
	"net/http"
	"quentinha_golang/src/configuration/logger"
	"quentinha_golang/src/configuration/rest_err"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

// DeleteProduct deletes a product with the specified ID.
// @Summary Delete product
// @Description Deletes a product based on the ID provided as a parameter.
// @Tags products
// @Accept json
// @Produce json
// @Param productId path string true "ID of the product to be deleted"
// @Success 200
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /deleteProduct/{productId} [delete]
func (pc *productControllerInterface) DeleteProduct(c *gin.Context) {

	logger.Info("Init deleteProduct controller",
		zap.String("journey", "deleteProduct"),
	)

	productId := c.Param("productId")

	if _, err := bson.ObjectIDFromHex(productId); err != nil {
		logger.Error("Error trying to convert productId to hex",
			err,
			zap.String("journey", "deleteProduct"),
		)

		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError("productId is not valid hex"))
		return
	}

	logger.Info("Product created successfully",
		zap.String("journey", "deleteProduct"),
	)

	err := pc.service.DeleteProductServices(productId)
	if err != nil {
		logger.Error("Error trying to call deleteProduct service",
			err,
			zap.String("journey", "deleteProduct"),
		)

		c.JSON(err.Code, err)
		return
	}

	logger.Info("deleteProduct controller executed successfully",
		zap.String("productId", productId),
		zap.String("journey", "deleteProduct"),
	)

	c.Status(http.StatusOK)
}

package address_controller

import (
	"quentinha_golang/src/configuration/logger"
	"quentinha_golang/src/configuration/validation"
	"quentinha_golang/src/controller/model/request/address_request"
	"quentinha_golang/src/model/domain/address_domain"
	"quentinha_golang/src/view"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateAddress Creates a new Address
// @Summary Create a new Address
// @Description Create a new Address with the provided Address information
// @Tags Addresss
// @Accept json
// @Produce json
// @Param AddressRequest body Address_request.AddressRequest true "Address information"
// @Success 200 {object} response.AddressResponse
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /Addresss [post]
func (ac *addressControllerInterface) CreateAddress(c *gin.Context) {

	logger.Info("Init CreateAddress controller",
		zap.String("journey", "createAddress"),
	)

	var addressRequest address_request.AddressRequest

	if err := c.ShouldBindJSON(&addressRequest); err != nil {
		logger.Error("Error trying to bind productRequest", err,
			zap.String("controller", "createAddress"),
		)

		errRest := validation.ValidateError(err, "Product")

		c.JSON(errRest.Code, errRest)
		return
	}

	logger.Info("Product request validated successfully",
		zap.String("journey", "createAddress"),
	)

	domain := address_domain.NewAddressDomain(
		addressRequest.City,
		addressRequest.Street,
		addressRequest.Number,
		addressRequest.CEP,
	)

	domainResult, err := ac.service.CreateAddressServices(domain)
	if err != nil {
		logger.Error("Error trying to call CreateAddressProductService",
			err,
			zap.String("journey", "createAddress"),
		)

		c.JSON(err.Code, err)
		return
	}

	logger.Info("CreateAddress controller executed successfully",
		zap.String("productId", domainResult.GetID()),
		zap.String("journey", "createAddress"),
	)

	c.JSON(201, view.ConvertAddressDomainToResponse(domainResult))
}

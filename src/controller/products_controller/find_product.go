package products_controller

import (
	"math"
	"net/http"
	"quentinha_golang/src/configuration/logger"
	"quentinha_golang/src/configuration/rest_err"
	"quentinha_golang/src/controller/model/response"
	"quentinha_golang/src/view"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

// FindAllProducts retrieves all products with pagination.
// @Summary Find all products
// @Description Retrieves a paginated list of all products saved in the database.
// @Tags Products
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page (max 100)" default(10)
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} response.PaginatedProductResponse "Paginated list of products"
// @Failure 400 {object} rest_err.RestErr "Invalid query parameters"
// @Failure 500 {object} rest_err.RestErr "Internal server error"
// @Router /products [get]
func (pc *productControllerInterface) FindAllProducts(c *gin.Context) {

	logger.Info("Init FindAllProducts controller",
		zap.String("journey", "findAllProducts"))

	pageQuery := c.DefaultQuery("page", "1")
	limitQuery := c.DefaultQuery("limit", "10")

	page, err := strconv.ParseInt(pageQuery, 10, 64)
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.ParseInt(limitQuery, 10, 64)
	if err != nil || limit <= 0 {
		limit = 10
	}

	// 🔥 Proteção profissional
	if limit > 100 {
		limit = 100
	}

	offset := (page - 1) * limit

	productsDomain, total, errRest :=
		pc.service.FindAllProductsServices(offset, limit)

	if errRest != nil {
		c.JSON(errRest.Code, errRest)
		return
	}

	totalPages := int64(math.Ceil(float64(total) / float64(limit)))

	response := response.PaginatedProductResponse{
		Data:       view.ConvertProductDomainListToResponse(productsDomain),
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
	}

	c.JSON(http.StatusOK, response)
}


// FindProductsByID retrieves user information based on the provided product ID.
// @Summary Find product by ID
// @Description Retrieves product details based on the product ID provided as a parameter.
// @Tags Products
// @Accept json
// @Produce json
// @Param productId path string true "ID of the product to be retrieved"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} response.productResponse "product information retrieved successfully"
// @Failure 400 {object} rest_err.RestErr "Error: Invalid product ID"
// @Failure 404 {object} rest_err.RestErr "product not found"
// @Router /getproductById/{productId} [get]
func (pc *productControllerInterface) FindProductsByID(c *gin.Context) {
	logger.Info("Init FindProductsByID controller", zap.String("journey", "FindProductsByID"))

	productId := c.Param("productId")

	if _, err := bson.ObjectIDFromHex(productId); err != nil {
		logger.Error("Error trying to validate productId", err, zap.String("journey", "FindProductsByID"))
		errorMessage := rest_err.NewBadRequestError("productId is not a valid user id")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	productDomain, err := pc.service.FindProductsByIDServices(productId)
	if err != nil {
		logger.Error("Error trying to call service FindProductsByIDServices", err, zap.String("journey", "FindProductsByID"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindProductsByID controller executed successfully", zap.String("journey", "FindProductsByID"))
	c.JSON(http.StatusOK, view.ConvertProductDomainToResponse(productDomain))
}



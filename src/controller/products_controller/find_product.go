package products_controller

import (
	"math"
	"net/http"
	"quentinha_golang/src/configuration/logger"
	"quentinha_golang/src/controller/model/response"
	"quentinha_golang/src/view"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

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

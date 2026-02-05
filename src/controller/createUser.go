package controller

import (
	"quentinha_golang/src/configuration/rest_err"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("CreateUser not implemented yet")
	c.JSON(err.Code, err)
}
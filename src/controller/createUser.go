package controller

import (
	"fmt"
	"log"
	"net/http"
	"quentinha_golang/src/configuration/validation"
	"quentinha_golang/src/controller/model/request"
	"quentinha_golang/src/controller/model/response"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	log.Println("Creating user...")

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object, error%s\n", err.Error())
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}
	
	fmt.Println(userRequest)
	response := response.UserResponse{
		ID:    "teste",
		Name:  userRequest.Name,
		Email: userRequest.Email,
		Age:   userRequest.Age,
	}
	c.JSON(http.StatusOK, response)
}

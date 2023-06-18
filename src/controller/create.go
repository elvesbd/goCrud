package controller

import (
	"fmt"

	"github.com/elvesbd/goCrud/src/controller/model/request"
	"github.com/elvesbd/goCrud/src/validation"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var userRequest request.UserRequest

	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}

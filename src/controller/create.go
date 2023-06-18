package controller

import (
	"fmt"

	"github.com/elvesbd/goCrud/src/configuration/restErr"
	"github.com/elvesbd/goCrud/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var userRequest request.UserRequest

	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		restErr := restErr.NewBadRequestError(fmt.Sprintf("There are some incorrect fields, error=%s", err.Error()))
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}

package controller

import (
	"github.com/elvesbd/goCrud/src/configuration/restErr"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	err := restErr.NewBadRequestError("Erro na rota")
	c.JSON(err.Code, err)
}

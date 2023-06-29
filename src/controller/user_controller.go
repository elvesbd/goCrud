package controller

import (
	"github.com/elvesbd/goCrud/src/model/service"
	"github.com/gin-gonic/gin"
)

type UserControllerInterface interface {
	FindById(c *gin.Context)
	FindByEmail(c *gin.Context)
	Delete(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}

func NewUserControllerInterface(
	serviceInterface service.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

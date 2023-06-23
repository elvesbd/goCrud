package controller

import (
	"net/http"

	"github.com/elvesbd/goCrud/src/configuration/logger"
	"github.com/elvesbd/goCrud/src/controller/model/request"
	"github.com/elvesbd/goCrud/src/model"
	"github.com/elvesbd/goCrud/src/model/service"
	"github.com/elvesbd/goCrud/src/validation"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func Create(c *gin.Context) {
	logger.Info("Init create user controller",
		zap.String("journey", "create user"),
	)
	var userRequest request.UserRequest

	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		logger.Error("Erro trying to validate user", err,
			zap.String("journey", "create user"),
		)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	user := model.NewUserDomain(
		userRequest.Name,
		userRequest.Email,
		userRequest.Password,
		userRequest.Age,
	)

	service := service.NewUserService()
	if err := service.Create(user); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "create user"),
	)

	c.String(http.StatusOK, "")
}

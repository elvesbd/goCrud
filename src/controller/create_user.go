package controller

import (
	"net/http"

	"github.com/elvesbd/goCrud/src/configuration/logger"
	"github.com/elvesbd/goCrud/src/controller/model/request"
	"github.com/elvesbd/goCrud/src/model"
	"github.com/elvesbd/goCrud/src/validation"
	"github.com/elvesbd/goCrud/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) Create(c *gin.Context) {
	logger.Info("Init create user controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
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

	result, err := uc.service.Create(user)
	if err != nil {
		logger.Error("Erro trying to call CreateUser service", err,
			zap.String("journey", "createUser"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("CreateUser controller executed successfully",
		zap.String("userId", user.GetId()),
		zap.String("journey", "create user"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(result))
}

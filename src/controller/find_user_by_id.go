package controller

import (
	"net/http"

	"github.com/elvesbd/goCrud/src/configuration/logger"
	"github.com/elvesbd/goCrud/src/configuration/restErr"
	"github.com/elvesbd/goCrud/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindById(c *gin.Context) {
	logger.Info("Init findUserById controller",
		zap.String("journey", "findUserById"),
	)

	id := c.Param("id")

	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		logger.Error(
			"Caiu aqui Init findUserById controller",
			err,
			zap.String("journey", "findUserById"),
		)
		errorMessage := restErr.NewBadRequestError("User ID is not a valid id")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByID(id)
	if err != nil {
		logger.Error(
			"Error trying to call findUserById services",
			err,
			zap.String("journey", "findUserById"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindById controller executed successfully",
		zap.String("journey", "findUserById"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

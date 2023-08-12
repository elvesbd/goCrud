package controller

import (
	"net/http"
	"net/mail"

	"github.com/elvesbd/goCrud/src/configuration/logger"
	"github.com/elvesbd/goCrud/src/configuration/restErr"
	"github.com/elvesbd/goCrud/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindByEmail(c *gin.Context) {
	logger.Info("Init findUserByEmail controller",
		zap.String("journey", "findUserByEmail"),
	)

	email := c.Param("email")

	if _, err := mail.ParseAddress(email); err != nil {
		logger.Error(
			"Init findUserByEmail controller",
			err,
			zap.String("journey", "findUserByEmail"),
		)
		errorMessage := restErr.NewBadRequestError("User Email is not a valid email")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmail(email)
	if err != nil {
		logger.Error(
			"Error trying to call findUserByEmail services",
			err,
			zap.String("journey", "findUserByEmail"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindByEmail controller executed successfully",
		zap.String("journey", "findUserByEmail"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

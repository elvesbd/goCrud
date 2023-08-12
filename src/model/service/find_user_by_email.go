package service

import (
	"github.com/elvesbd/goCrud/src/configuration/logger"
	"github.com/elvesbd/goCrud/src/configuration/restErr"
	"github.com/elvesbd/goCrud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByEmail(email string) (
	model.UserDomainInterface, *restErr.RestErr,
) {
	logger.Info("Init findUserByID services",
		zap.String("journey", "findUserByID"),
	)
	return ud.userRepository.FindUserByEmail(email)
}

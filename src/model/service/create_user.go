package service

import (
	"github.com/elvesbd/goCrud/src/configuration/logger"
	"github.com/elvesbd/goCrud/src/configuration/restErr"
	"github.com/elvesbd/goCrud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) Create(userDomain model.UserDomainInterface) (model.UserDomainInterface, *restErr.RestErr) {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.Create(userDomain)
	if err != nil {
		logger.Error("Init createUser model", err, zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info("CreateUser service executed successfully",
		zap.String("userId", userDomainRepository.GetId()),
		zap.String("journey", "createUser"),
	)
	return userDomainRepository, nil
}

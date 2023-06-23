package service

import (
	"fmt"

	"github.com/elvesbd/goCrud/src/configuration/logger"
	"github.com/elvesbd/goCrud/src/configuration/restErr"
	"github.com/elvesbd/goCrud/src/model"
	"go.uber.org/zap"
)

func (u *userDomainService) Create(userDomain model.UserDomainInterface) *restErr.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()
	fmt.Println(userDomain.GetPassword())
	return nil
}

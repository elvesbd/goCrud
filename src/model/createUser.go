package model

import (
	"fmt"

	"github.com/elvesbd/goCrud/src/configuration/logger"
	"github.com/elvesbd/goCrud/src/configuration/restErr"
	"go.uber.org/zap"
)

func (u *UserDomain) CreateUser(UserDomain) *restErr.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	u.EncryptPassword()
	fmt.Println(u)
	return nil
}

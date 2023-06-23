package service

import (
	"github.com/elvesbd/goCrud/src/configuration/restErr"
	"github.com/elvesbd/goCrud/src/model"
)

func (*userDomainService) Update(
	userId string, userDomain model.UserDomainInterface,
) *restErr.RestErr {
	return nil
}

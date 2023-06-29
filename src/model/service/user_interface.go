package service

import (
	"github.com/elvesbd/goCrud/src/configuration/restErr"
	"github.com/elvesbd/goCrud/src/model"
)

type userDomainService struct {
}

type UserDomainService interface {
	Create(model.UserDomainInterface) *restErr.RestErr
	Update(string, model.UserDomainInterface) *restErr.RestErr
	Find(string) (*model.UserDomainInterface, *restErr.RestErr)
	Delete(string) *restErr.RestErr
}

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

package service

import (
	"github.com/elvesbd/goCrud/src/configuration/restErr"
	"github.com/elvesbd/goCrud/src/model"
	"github.com/elvesbd/goCrud/src/model/repository"
)

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	Create(model.UserDomainInterface) (model.UserDomainInterface, *restErr.RestErr)
	Update(string, model.UserDomainInterface) *restErr.RestErr
	Find(string) (*model.UserDomainInterface, *restErr.RestErr)
	Delete(string) *restErr.RestErr
}

func NewUserDomainService(userRepository repository.UserRepository) UserDomainService {
	return &userDomainService{userRepository}
}

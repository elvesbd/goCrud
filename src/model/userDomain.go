package model

import "github.com/elvesbd/goCrud/src/configuration/restErr"

type UserDomain struct {
	Name     string
	Email    string
	Age      string
	Password string
}

type UserDomainInterface interface {
	CreateUser(UserDomain) *restErr.RestErr
	UpdateUser(string, UserDomain) *restErr.RestErr
	FindUser(string) (*UserDomain, *restErr.RestErr)
	DeleteUser(string) *restErr.RestErr
}

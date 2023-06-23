package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/elvesbd/goCrud/src/configuration/restErr"
)

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

func (u *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(u.Password))
	u.Password = hex.EncodeToString(hash.Sum(nil))
}

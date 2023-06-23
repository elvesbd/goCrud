package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/elvesbd/goCrud/src/configuration/restErr"
)

type UserDomain struct {
	Name     string
	Email    string
	Password string
	Age      int8
}

func NewUser(name, email, password string, age int8) UserDomainInterface {
	return &UserDomain{
		name, email, password, age,
	}
}

type UserDomainInterface interface {
	Create() *restErr.RestErr
	Update(string) *restErr.RestErr
	Find(string) (*UserDomain, *restErr.RestErr)
	Delete(string) *restErr.RestErr
}

func (u *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(u.Password))
	u.Password = hex.EncodeToString(hash.Sum(nil))
}

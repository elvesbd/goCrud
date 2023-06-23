package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/elvesbd/goCrud/src/configuration/restErr"
)

type userDomain struct {
	name     string
	email    string
	password string
	age      int8
}

func NewUser(name, email, password string, age int8) UserDomainInterface {
	return &userDomain{
		name, email, password, age,
	}
}

type UserDomainInterface interface {
	Create() *restErr.RestErr
	Update(string) *restErr.RestErr
	Find(string) (*userDomain, *restErr.RestErr)
	Delete(string) *restErr.RestErr
}

func (u *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(u.password))
	u.password = hex.EncodeToString(hash.Sum(nil))
}

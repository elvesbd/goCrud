package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/elvesbd/goCrud/src/configuration/restErr"
)

type UserDomain struct {
	Name     string
	Email    string
	Age      int8
	Password string
}

func NewUser(name, email, password string, age int8) *UserDomain {
	return &UserDomain{
		Name:     name,
		Email:    email,
		Age:      age,
		Password: password,
	}
}

type UserDomainInterface interface {
	CreateUser() *restErr.RestErr
	UpdateUser(string) *restErr.RestErr
	FindUser(string) (*UserDomain, *restErr.RestErr)
	DeleteUser(string) *restErr.RestErr
}

func (u *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(u.Password))
	u.Password = hex.EncodeToString(hash.Sum(nil))
}

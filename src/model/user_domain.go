package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

type userDomain struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int8   `json:"age"`
}

type UserDomainInterface interface {
	GetName() string
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetJSONValue() (string, error)

	SetId(string)

	EncryptPassword()
}

func NewUserDomain(name, email, password string, age int8) UserDomainInterface {
	return &userDomain{
		Name:     name,
		Email:    email,
		Password: password,
		Age:      age,
	}
}

func (u *userDomain) GetName() string {
	return u.Name
}

func (u *userDomain) GetEmail() string {
	return u.Email
}

func (u *userDomain) GetPassword() string {
	return u.Password
}

func (u *userDomain) GetAge() int8 {
	return u.Age
}

func (u *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(u.Password))
	u.Password = hex.EncodeToString(hash.Sum(nil))
}

func (ud *userDomain) GetJSONValue() (string, error) {
	b, err := json.Marshal(ud)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (ud *userDomain) SetId(id string) {
	ud.Id = id
}

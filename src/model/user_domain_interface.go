package model

type UserDomainInterface interface {
	GetName() string
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetId() string
	SetId(string)
	EncryptPassword()
}

func NewUserDomain(name, email, password string, age int8) UserDomainInterface {
	return &userDomain{
		name:     name,
		email:    email,
		password: password,
		age:      age,
	}
}

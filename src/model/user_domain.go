package model

type userDomain struct {
	id       string
	email    string
	password string
	name     string
	age      int8
}

func (ud *userDomain) GetId() string {
	return ud.id
}

func (ud *userDomain) SetId(id string) {
	ud.id = id
}

func (u *userDomain) GetName() string {
	return u.name
}

func (u *userDomain) GetEmail() string {
	return u.email
}

func (u *userDomain) GetPassword() string {
	return u.password
}

func (u *userDomain) GetAge() int8 {
	return u.age
}

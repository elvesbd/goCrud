package model

type userDomain struct {
	Id       string
	Name     string
	Email    string
	Password string
	Age      int8
}

func (ud *userDomain) GetId() string {
	return ud.Id
}

func (ud *userDomain) SetId(id string) {
	ud.Id = id
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

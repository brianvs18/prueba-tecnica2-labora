package entities

import "fmt"

type User struct {
	Id    string
	Name  string
	Email string
}

func UserBuilder(id string, name string, email string) *User {
	return &User{
		Id:    id,
		Name:  name,
		Email: email,
	}
}

func (u *User) String() string {
	return fmt.Sprintf("User{Id: %s, Name: %s, Email: %s}", u.Id, u.Name, u.Email)
}

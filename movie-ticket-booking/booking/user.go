package booking

import "github.com/google/uuid"

type User struct {
	Id    string
	Name  string
	Email string
}

func NewUser(name, email string) *User {
	return &User{
		Id:    uuid.NewString(),
		Name:  name,
		Email: email,
	}
}

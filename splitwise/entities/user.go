package entities

type User struct {
	id           string
	name         string
	email        string
	balanceSheet *BalanceSheet
}

func NewUser(id, name, email string) *User {
	user := User{
		id:    id,
		name:  name,
		email: email,
	}
	user.balanceSheet = NewBalanceSheet(&user)
	return &user
}

func (u *User) GetId() string {
	return u.id
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetBalanceSheet() *BalanceSheet {
	return u.balanceSheet
}

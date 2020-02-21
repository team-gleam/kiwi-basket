package user

import "github.com/the-gleam/kiwi-basket/domain/model/user/username"

type User struct {
	username username.Username
}

func NewUser(u username.Username) *User {
	return &User{u}
}

func (u *User) Username() username.Username {
	return u.username
}

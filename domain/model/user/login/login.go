package login

import "github.com/the-gleam/kiwi-basket/domain/model/user/username"

type Login struct {
	username       username.Username
	hashedPassword string
}

func NewLogin(u username.Username, p string) Login {
	return Login{u, p}
}

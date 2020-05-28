package login

import "github.com/team-gleam/kiwi-basket/server/src/domain/model/user/username"

type Login struct {
	username       username.Username
	hashedPassword string
}

func NewLogin(u username.Username, p string) Login {
	return Login{u, p}
}

func (l Login) Username() username.Username {
	return l.username
}

func (l Login) HashedPassword() string {
	return l.hashedPassword
}

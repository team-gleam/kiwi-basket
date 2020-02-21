package auth

import "github.com/the-gleam/kiwi-basket/domain/model/user"

type Auth struct {
	username user.Username
	token    string
}

func (a Auth) Token() string {
	return a.token
}

func (a Auth) Username() user.Username {
	return a.username
}

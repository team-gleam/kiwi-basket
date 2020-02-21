package auth

import "github.com/the-gleam/kiwi-basket/domain/model/user/username"

type Auth struct {
	username username.Username
	token    string
}

func (a Auth) Token() string {
	return a.token
}

func (a Auth) Username() username.Username {
	return a.username
}

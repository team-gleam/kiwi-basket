package auth

import (
	"github.com/the-gleam/kiwi-basket/domain/model/user/token"
	"github.com/the-gleam/kiwi-basket/domain/model/user/username"
)

type Auth struct {
	username username.Username
	token    token.Token
}

func (a Auth) Username() username.Username {
	return a.username
}

func (a Auth) Token() token.Token {
	return a.token
}

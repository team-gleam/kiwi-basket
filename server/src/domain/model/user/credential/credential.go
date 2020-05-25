package credential

import (
	"github.com/team-gleam/kiwi-basket/domain/model/user/token"
	"github.com/team-gleam/kiwi-basket/domain/model/user/username"
)

type Auth struct {
	username username.Username
	token    token.Token
}

func NewAuth(u username.Username, t token.Token) Auth {
	return Auth{u, t}
}

func (a Auth) Username() username.Username {
	return a.username
}

func (a Auth) Token() token.Token {
	return a.token
}

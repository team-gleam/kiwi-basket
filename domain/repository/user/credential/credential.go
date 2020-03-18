package credential

import (
	"github.com/the-gleam/kiwi-basket/domain/model/user/credential"
	"github.com/the-gleam/kiwi-basket/domain/model/user/token"
)

type ICredentialRepository interface {
	Append(credential.Auth) error
	Remove(credential.Auth) error
	Exists(token.Token) (bool, error)
	Get(token.Token) (credential.Auth, error)
}

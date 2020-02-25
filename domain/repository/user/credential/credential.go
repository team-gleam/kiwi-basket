package credential

import "github.com/the-gleam/kiwi-basket/domain/model/user/credential"

type ICredentialRepository interface {
	Append(credential.Auth) error
	Remove(credential.Auth) error
	Exists(credential.Auth) bool
}

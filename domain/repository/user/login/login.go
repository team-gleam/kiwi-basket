package login

import (
	"github.com/the-gleam/kiwi-basket/domain/model/user/login"
	"github.com/the-gleam/kiwi-basket/domain/model/user/username"
)

type ILoginRepository interface {
	Create(login.Login) error
	Delete(login.Login) error
	Exists(username.Username) (bool, error)
	Get(username.Username) (login.Login, error)
}

package login

import "github.com/the-gleam/kiwi-basket/domain/model/user/login"

type ILoginRepository interface {
	Create(login.Login) error
	Delete(login.Login) error
	Exists(login.Login) (bool, error)
	Match(login.Login) (bool, error)
}

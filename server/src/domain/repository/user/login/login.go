package login

import (
	"github.com/team-gleam/kiwi-basket/server/src/domain/model/user/login"
	"github.com/team-gleam/kiwi-basket/server/src/domain/model/user/username"
)

type ILoginRepository interface {
	Create(login.Login) error
	Delete(login.Login) error
	Exists(username.Username) (bool, error)
	Get(username.Username) (login.Login, error)
}

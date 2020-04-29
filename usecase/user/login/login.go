package login

import (
	"fmt"

	loginModel "github.com/team-gleam/kiwi-basket/domain/model/user/login"
	loginRepository "github.com/team-gleam/kiwi-basket/domain/repository/user/login"
)

type LoginUsecase struct {
	loginRepository loginRepository.ILoginRepository
}

func NewLoginUsecase(r loginRepository.ILoginRepository) LoginUsecase {
	return LoginUsecase{r}
}

const (
	UsernameAlreadyExists = "username already exists"
	UsernameNotFound      = "username not found"
)

func (u LoginUsecase) Add(l loginModel.Login) error {
	exist, err := u.loginRepository.Exists(l.Username())
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf(UsernameAlreadyExists)
	}

	return u.loginRepository.Create(l)
}

func (u LoginUsecase) Delete(l loginModel.Login) error {
	exist, err := u.loginRepository.Exists(l.Username())
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf(UsernameNotFound)
	}

	return u.loginRepository.Delete(l)
}

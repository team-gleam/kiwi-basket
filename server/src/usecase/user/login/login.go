package login

import (
	"fmt"

	loginModel "github.com/team-gleam/kiwi-basket/server/src/domain/model/user/login"
	loginRepository "github.com/team-gleam/kiwi-basket/server/src/domain/repository/user/login"
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

func (u LoginUsecase) Verify(login loginModel.Login) (bool, error) {
	exist, err := u.loginRepository.Exists(login.Username())
	if err != nil {
		return false, err
	}
	if !exist {
		return false, fmt.Errorf(UsernameNotFound)
	}

	l, err := u.loginRepository.Get(login.Username())
	if err != nil {
		return false, err
	}
	if !(login.HashedPassword() == l.HashedPassword()) {
		return false, nil
	}

	return true, nil
}
